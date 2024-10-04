package code_gen

import (
	"fmt"
	"io"
	"sort"

	"github.com/pattyshack/gt/codegen"
	lr "github.com/pattyshack/gt/tools/lr/internal"
	"github.com/pattyshack/gt/tools/lr/internal/code_gen/go_template"
	"github.com/pattyshack/gt/tools/lr/internal/parser"
)

var (
	escapedChar = map[string]byte{
		"'\\t'":  '\t',
		"'\\n'":  '\n',
		"'\\''":  '\'',
		"'\\\\'": '\\',
	}
)

type NameGenerator struct {
	prefix string

	nameCount map[string]int
}

func NewNameGenerator(prefix string) *NameGenerator {
	return &NameGenerator{
		prefix:    prefix,
		nameCount: map[string]int{},
	}
}

func (ng *NameGenerator) Add(name string) string {
	ng.nameCount[name] += 1
	cnt := ng.nameCount[name]
	if cnt > 1 {
		name = fmt.Sprintf("%s_%d", name, cnt)
	}
	return name
}

func (ng *NameGenerator) Public(name string) string {
	name = ng.prefix + name
	return ng.Add(name)
}

func (ng *NameGenerator) Internal(name string) string {
	name = "_" + ng.prefix + name
	return ng.Add(name)
}

func populateGoCodeGenVariables(
	prefix string,
	Terms map[string]*lr.Term,
	OrderedStates []*lr.ItemSet,
	valueTypes map[string]*lr.Param,
	nameGen *NameGenerator) error {

	for _, term := range Terms {
		valueType := valueTypes[term.ValueType]
		if valueType == nil {
			return fmt.Errorf(
				"Undefined value type for <%s> %s",
				term.ValueType,
				term.LRLocation)
		}
		term.CodeGenType = valueType.ParamType

		if term.SymbolId == parser.LRCharacterToken {
			term.CodeGenSymbolConst = term.Name
		} else {
			suffix := "Type"
			if term.IsTerminal {
				suffix = "Token"
			}
			term.CodeGenSymbolConst = nameGen.Public(
				codegen.SnakeToCamel(term.Name) + suffix)
		}

		if !term.IsTerminal {
			term.CodeGenReducerInterface = nameGen.Public(
				codegen.SnakeToCamel(term.Name) + "Reducer")
		}

		for _, clause := range term.Clauses {
			reducerName := nameGen.Add(
				codegen.SnakeToCamel(clause.Label) +
					"To" +
					codegen.SnakeToCamel(term.Name))
			clause.CodeGenReducerName = reducerName
			clause.CodeGenReducerNameConst = nameGen.Internal(
				"Reduce" + reducerName)
			clause.CodeGenReduceAction = nameGen.Internal(
				"Reduce" + reducerName + "Action")
		}
	}

	for _, state := range OrderedStates {
		state.CodeGenConst = nameGen.Internal(
			fmt.Sprintf("State%d", state.StateNum))
		state.CodeGenAction = nameGen.Internal(
			fmt.Sprintf("GotoState%dAction", state.StateNum))
	}

	return nil
}

func GenerateGoLRCode(
	grammar *lr.Grammar,
	states *lr.LRStates) (
	io.WriterTo,
	error) {

	cfg := grammar.LangSpecs.Go
	if cfg == nil {
		return nil, fmt.Errorf("go configuration not specified in lang_specs")
	}

	if cfg.Package == "" {
		return nil, fmt.Errorf("package name not specified")
	}

	imports := codegen.NewGoImports()
	nameGen := NewNameGenerator(cfg.Prefix)

	endSymbol := nameGen.Internal("EndMarker")
	wildcardSymbol := nameGen.Internal("WildcardMarker")
	genericSymbol := nameGen.Public("GenericSymbol")

	orderedValueTypes := lr.ParamList{
		&lr.Param{lr.Generic, imports.Obj(genericSymbol)},
	}
	for name, valueType := range cfg.ValueTypes {
		orderedValueTypes = append(
			orderedValueTypes,
			&lr.Param{name, imports.Obj(valueType)})
	}
	sort.Sort(orderedValueTypes)

	valueTypes := make(map[string]*lr.Param, len(orderedValueTypes))
	for _, vt := range orderedValueTypes {
		valueTypes[vt.Name] = vt
	}

	err := populateGoCodeGenVariables(
		cfg.Prefix,
		grammar.Terms,
		states.OrderedStates,
		valueTypes,
		nameGen)
	if err != nil {
		return nil, err
	}

	startMarker := &lr.Term{
		Name:               lr.StartMarker,
		IsTerminal:         true,
		ValueType:          lr.Generic,
		CodeGenSymbolConst: lr.StartMarker,
		CodeGenType:        genericSymbol,
	}
	wildcard := &lr.Term{
		Name:               lr.Wildcard,
		IsTerminal:         true,
		ValueType:          lr.Generic,
		CodeGenSymbolConst: wildcardSymbol,
		CodeGenType:        genericSymbol,
	}
	endMarker := &lr.Term{
		Name:               lr.EndMarker,
		IsTerminal:         true,
		ValueType:          lr.Generic,
		CodeGenSymbolConst: endSymbol,
		CodeGenType:        genericSymbol,
	}

	grammar.Terms[startMarker.Name] = startMarker
	grammar.Terms[wildcard.Name] = wildcard
	grammar.Terms[endMarker.Name] = endMarker

	orderedSymbolNames := []string{lr.StartMarker, lr.Wildcard, lr.EndMarker}
	for _, term := range grammar.Terminals {
		orderedSymbolNames = append(orderedSymbolNames, term.Name)
	}
	for _, term := range grammar.NonTerminals {
		orderedSymbolNames = append(orderedSymbolNames, term.Name)
	}

	file := &go_template.File{
		Package:              cfg.Package,
		Imports:              imports,
		ActionType:           nameGen.Internal("Action"),
		ActionIdType:         nameGen.Internal("ActionType"),
		ShiftAction:          nameGen.Internal("ShiftAction"),
		ReduceAction:         nameGen.Internal("ReduceAction"),
		AcceptAction:         nameGen.Internal("AcceptAction"),
		ShiftAndReduceAction: nameGen.Internal("ShiftAndReduceAction"),
		StateIdType:          nameGen.Internal("StateId"),
		ReduceType:           nameGen.Internal("ReduceType"),
		SymbolType:           nameGen.Public("Symbol"),
		GenericSymbolType:    genericSymbol,
		StackItemType:        nameGen.Internal("StackItem"),
		StackType:            nameGen.Internal("Stack"),
		SymbolStackType:      nameGen.Internal("PseudoSymbolStack"),
		SymbolIdType:         nameGen.Public("SymbolId"),
		EndSymbolId:          endSymbol,
		WildcardSymbolId:     wildcardSymbol,
		LocationType:         nameGen.Public("Location"),
		RealLocationType: imports.Obj(
			"github.com/pattyshack/gt/lexutil.Location"),
		TokenType:             nameGen.Public("Token"),
		LexerType:             nameGen.Public("Lexer"),
		ReducerType:           nameGen.Public("Reducer"),
		ErrHandlerType:        nameGen.Public("ParseErrorHandler"),
		DefaultErrHandlerType: nameGen.Public("DefaultParseErrorHandler"),
		ExpectedTerminalsFunc: nameGen.Public("ExpectedTerminals"),
		ParseFuncPrefix:       nameGen.Public("Parse"),
		InternalParseFunc:     nameGen.Internal("Parse"),
		TableKeyType:          nameGen.Internal("ActionTableKey"),
		ActionTableType:       nameGen.Internal("ActionTableType"),
		ActionTable:           nameGen.Internal("ActionTable"),
		Sprintf:               imports.Obj("fmt.Sprintf"),
		Errorf:                imports.Obj("fmt.Errorf"),
		NewLocationError: imports.Obj(
			"github.com/pattyshack/gt/lexutil.NewLocationError"),
		EOF:                       imports.Obj("io.EOF"),
		OrderedSymbolNames:        orderedSymbolNames,
		Grammar:                   grammar,
		States:                    states,
		OrderedValueTypes:         orderedValueTypes,
		OutputDebugNonKernelItems: cfg.OutputDebugNonKernelItems,
		GenerateEndPos:            cfg.GenerateEndPos,
	}

	return codegen.NewFormattedGoSource(file), nil
}
