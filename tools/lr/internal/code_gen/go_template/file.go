// Auto-generated from source: file.template

package go_template

import (
	_fmt "fmt"
	_io "io"

	"fmt"
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/codegen"
	lr "github.com/pattyshack/gt/tools/lr/internal"
	"github.com/pattyshack/gt/tools/lr/internal/code_gen/debug_template"
	parser "github.com/pattyshack/gt/tools/lr/internal/parser"
)

type File struct {
	Package                   string
	Imports                   io.WriterTo
	ActionType                string
	ActionIdType              string
	ShiftAction               string
	ReduceAction              string
	AcceptAction              string
	ShiftAndReduceAction      string
	StateIdType               string
	ReduceType                string
	SymbolType                string
	GenericSymbolType         string
	StackItemType             string
	StackType                 string
	SymbolStackType           string
	SymbolIdType              string
	EndSymbolId               string
	WildcardSymbolId          string
	LocationType              string
	RealLocationType          interface{}
	TokenType                 string
	LexerType                 string
	ReducerType               string
	ErrHandlerType            string
	DefaultErrHandlerType     string
	ExpectedTerminalsFunc     string
	ParseFuncPrefix           string
	InternalParseFunc         string
	TableKeyType              string
	ActionTableType           string
	ActionTable               string
	Sprintf                   interface{}
	Errorf                    interface{}
	NewLocationError          interface{}
	EOF                       interface{}
	OrderedSymbolNames        []string
	Grammar                   *lr.Grammar
	States                    *lr.LRStates
	OrderedValueTypes         lr.ParamList
	OutputDebugNonKernelItems bool
	GenerateEndPos            bool
}

func (File) Name() string { return "File" }

func (template *File) writeValue(
	output _io.Writer,
	value interface{},
	loc string) (
	int,
	error) {

	var valueBytes []byte
	switch val := value.(type) {
	case _fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(_fmt.Sprintf("%v", val))
	default:
		return 0, _fmt.Errorf(
			"Unsupported output value type (%s): %v",
			loc,
			value)
	}

	return output.Write(valueBytes)
}

func (_template *File) WriteTo(
	_output _io.Writer) (
	int64,
	error) {

	_numWritten := int64(0)

	Package := _template.Package
	Imports := _template.Imports
	ActionType := _template.ActionType
	ActionIdType := _template.ActionIdType
	ShiftAction := _template.ShiftAction
	ReduceAction := _template.ReduceAction
	AcceptAction := _template.AcceptAction
	ShiftAndReduceAction := _template.ShiftAndReduceAction
	StateIdType := _template.StateIdType
	ReduceType := _template.ReduceType
	SymbolType := _template.SymbolType
	GenericSymbolType := _template.GenericSymbolType
	StackItemType := _template.StackItemType
	StackType := _template.StackType
	SymbolStackType := _template.SymbolStackType
	SymbolIdType := _template.SymbolIdType
	EndSymbolId := _template.EndSymbolId
	WildcardSymbolId := _template.WildcardSymbolId
	LocationType := _template.LocationType
	RealLocationType := _template.RealLocationType
	TokenType := _template.TokenType
	LexerType := _template.LexerType
	ReducerType := _template.ReducerType
	ErrHandlerType := _template.ErrHandlerType
	DefaultErrHandlerType := _template.DefaultErrHandlerType
	ExpectedTerminalsFunc := _template.ExpectedTerminalsFunc
	ParseFuncPrefix := _template.ParseFuncPrefix
	InternalParseFunc := _template.InternalParseFunc
	TableKeyType := _template.TableKeyType
	ActionTableType := _template.ActionTableType
	ActionTable := _template.ActionTable
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	NewLocationError := _template.NewLocationError
	EOF := _template.EOF
	OrderedSymbolNames := _template.OrderedSymbolNames
	Grammar := _template.Grammar
	States := _template.States
	OrderedValueTypes := _template.OrderedValueTypes
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems
	GenerateEndPos := _template.GenerateEndPos

	// file.template:78:0
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:31
	{
		_n, _err := _template.writeValue(
			_output,
			(Grammar.Source),
			"file.template:78:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:78:48
	{
		_n, _err := _output.Write([]byte(`

package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:80:8
	{
		_n, _err := _template.writeValue(
			_output,
			(Package),
			"file.template:80:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:80:16
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:82:0
	{
		_n, _err := (Imports).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:82:20
	{
		_n, _err := _output.Write([]byte(`
type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:84:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:84:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:84:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:87:0

	charIds := map[int]struct{}{}
	charTermToIds := map[string]int{}
	for _, term := range Grammar.Terminals {
		if term.SymbolId == parser.LRCharacterToken {
			rune, _ := utf8.DecodeRuneInString(term.Name[1 : len(term.Name)-1])
			charIds[int(rune)] = struct{}{}
			charTermToIds[term.Name] = int(rune)
		}
	}

	_nextId := 256
	nextId := func() int {
		for {
			id := _nextId
			_, ok := charIds[id]

			_nextId += 1

			if !ok {
				return id
			}
		}
	}

	// file.template:113:0
	for _, term := range Grammar.Terminals {
		// file.template:114:4
		if term.SymbolId != parser.LRIdentifierToken {
			// file.template:114:53
			{
				_n, _err := _output.Write([]byte(`
    // char token `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:18
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:115:18")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:115:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:49
			{
				_n, _err := _template.writeValue(
					_output,
					(charTermToIds[term.Name]),
					"file.template:115:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:115:76
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:118:0
	for _, term := range Grammar.Terminals {
		// file.template:119:4
		if term.SymbolId == parser.LRIdentifierToken {
			// file.template:119:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:120:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:120:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId()),
					"file.template:120:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:120:60
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:122:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:125:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:18
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:21
	{
		_n, _err := _template.writeValue(
			_output,
			(RealLocationType),
			"file.template:125:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:38
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:127:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:15
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:128:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:129:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:130:0
	if GenerateEndPos {
		// file.template:130:22
		{
			_n, _err := _output.Write([]byte(`
    End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:131:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:131:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:132:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:135:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:136:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:17
	{
		_n, _err := _output.Write([]byte(`
    StartPos `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:137:13
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:137:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:138:0
	if GenerateEndPos {
		// file.template:138:22
		{
			_n, _err := _output.Write([]byte(`
    EndPos `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:139:11
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:139:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:140:8
	{
		_n, _err := _output.Write([]byte(`
}

func (t `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:8
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:143:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:26
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:33
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:143:33")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:46
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:58
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:143:58")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:71
	{
		_n, _err := _output.Write([]byte(` }

func (t `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:145:8
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:145:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:145:26
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:145:34
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:145:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:145:47
	{
		_n, _err := _output.Write([]byte(` { return t.StartPos }
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:147:0
	if GenerateEndPos {
		// file.template:147:22
		{
			_n, _err := _output.Write([]byte(`
func (t `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:148:8
		{
			_n, _err := _template.writeValue(
				_output,
				(GenericSymbolType),
				"file.template:148:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:148:26
		{
			_n, _err := _output.Write([]byte(`) End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:148:34
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:148:34")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:148:47
		{
			_n, _err := _output.Write([]byte(` { return t.EndPos }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:149:8
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:151:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:151:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:151:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:153:55
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:153:55")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:153:73
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:154:12
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:154:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:154:22
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:156:22
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:156:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:156:35
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:159:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:160:4
		if rule.NumReducerClauses == 0 {
			// file.template:161:8
			continue
		}
		// file.template:162:12
		{
			_n, _err := _output.Write([]byte(`
type `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:163:5
		{
			_n, _err := _template.writeValue(
				_output,
				(rule.CodeGenReducerInterface),
				"file.template:163:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:163:36
		{
			_n, _err := _output.Write([]byte(` interface {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:164:4
		for clauseIdx, clause := range rule.Clauses {
			// file.template:165:8
			if clause.Passthrough {
				// file.template:166:10
				continue
			}
			// file.template:168:8
			if clauseIdx > 0 {
				// file.template:168:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:172:8
			if clause.Label == "" {
				// file.template:172:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:173:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:173:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:173:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:173:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:173:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:173:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:174:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:175:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:175:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:175:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:175:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:178:8
			paramNameCount := map[string]int{}
			// file.template:178:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:179:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:179:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:179:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:180:8
			for termIdx, term := range clause.Bindings {
				// file.template:182:12

				paramName := ""
				if term.SymbolId == parser.LRCharacterToken {
					paramName = "char"
				} else {
					// hack: append "_" to the end of the name ensures the
					// name is never a go keyword
					paramName = codegen.SnakeToCamel(term.Name) + "_"
				}

				paramNameCount[paramName] += 1
				cnt := paramNameCount[paramName]
				if cnt > 1 {
					paramName = fmt.Sprintf("%s%d", paramName, cnt)
				}

				suffix := ""
				if termIdx != len(clause.Bindings) {
					suffix = ", "
				}

				// file.template:205:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:205:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:205:10
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:205:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"file.template:205:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:205:30
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"file.template:205:30")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:206:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:207:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"file.template:207:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:207:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:208:13
		{
			_n, _err := _output.Write([]byte(`}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:210:7
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:212:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:212:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:212:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:213:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:214:4
		if len(rule.Clauses) == 0 {
			// file.template:215:8
			continue
		}
		// file.template:217:2
		if rule.NumReducerClauses > 0 {
			// file.template:217:36
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:218:2
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenReducerInterface),
					"file.template:218:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:220:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:223:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:20
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:224:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:224:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:224:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:224:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:224:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:224:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:227:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:227:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:227:27
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:229:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:28
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:46
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:229:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:56
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:64
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:229:64")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:74
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:230:11
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:230:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:230:28
	{
		_n, _err := _output.Write([]byte(`(
        nextToken.Loc(),
        "syntax error: unexpected symbol %s. expecting %v",
        nextToken.Id(),
        `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:234:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:234:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:234:30
	{
		_n, _err := _output.Write([]byte(`(stack[len(stack)-1].StateId))
}

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:237:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:27
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:31
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:237:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:43
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:47
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:237:47")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:237:60
	{
		_n, _err := _output.Write([]byte(` {
  switch id {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:239:0
	for _, state := range States.OrderedStates {
		// file.template:240:2

		_, ok := state.Reduce[lr.Wildcard]
		if ok {
			continue
		}

		// file.template:245:4
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:246:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:246:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:246:28
		{
			_n, _err := _output.Write([]byte(`:
    return []`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:247:13
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:247:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:247:26
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:248:4
		for _, term := range Grammar.Terminals {
			// file.template:249:6

			_, foundGoto := state.Goto[term.Name]
			_, foundReduce := state.Reduce[term.Name]
			_, foundShiftAndReduce := state.ShiftAndReduce[term.Name]

			if !foundGoto && !foundReduce && !foundShiftAndReduce {
				continue
			}

			// file.template:258:9
			{
				_n, _err := _output.Write([]byte(`      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:259:6
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:259:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:259:32
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:261:4
		if _, ok := state.Reduce[lr.EndMarker]; ok {
			// file.template:261:52
			{
				_n, _err := _output.Write([]byte(`      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:262:6
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"file.template:262:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:262:18
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:263:13
		{
			_n, _err := _output.Write([]byte(`    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:265:8
	{
		_n, _err := _output.Write([]byte(`
  }

  return nil
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:271:0
	for idx, start := range Grammar.Starts {
		// file.template:272:4

		parseSuffix := ""
		if len(Grammar.Starts) > 1 {
			parseSuffix = codegen.SnakeToCamel(start.Name)
		}

		// file.template:279:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:280:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:280:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:37
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:44
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:280:44")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:54
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:64
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:280:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:76
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:282:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:282:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:282:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:284:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:284:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:284:29
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:284:29")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:284:43
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
        lexer,
        reducer,
        `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:287:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandlerType),
				"file.template:287:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:287:30
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:290:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:290:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:290:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:290:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:290:37
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:291:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:291:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:291:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:292:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:292:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:292:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:293:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandlerType),
				"file.template:293:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:293:30
		{
			_n, _err := _output.Write([]byte(`,
) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:295:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:295:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:295:24
		{
			_n, _err := _output.Write([]byte(`,
    error,
) {
    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:298:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParseFunc),
				"file.template:298:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:298:35
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:298:64
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:298:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:299:47
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:301:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:301:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:301:42
		{
			_n, _err := _output.Write([]byte(`
        return errRetVal, err
    }
    return item.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:304:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"file.template:304:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:304:34
		{
			_n, _err := _output.Write([]byte(`, nil
}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:306:7
	{
		_n, _err := _output.Write([]byte(`

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:313:5
	{
		_n, _err := _template.writeValue(
			_output,
			(InternalParseFunc),
			"file.template:313:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:313:23
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:314:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:314:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:314:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:315:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:315:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:315:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:316:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:316:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:316:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:317:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:317:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:317:27
	{
		_n, _err := _output.Write([]byte(`,
) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:319:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:319:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:319:19
	{
		_n, _err := _output.Write([]byte(`,
    error,
) {
    stateStack := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:322:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:322:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:322:28
	{
		_n, _err := _output.Write([]byte(`{
        // Note: we don't have to populate the start symbol since its value
        // is never accessed.
        &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:325:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:325:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:325:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:328:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:328:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:328:36
	{
		_n, _err := _output.Write([]byte(`{lexer: lexer}

    for {
        nextSymbol, err := symbolStack.Top()
        if err != nil {
            return nil, err
        }

        action, ok := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:336:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:336:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:336:34
	{
		_n, _err := _output.Write([]byte(`.Get(
            stateStack[len(stateStack)-1].StateId,
            nextSymbol.Id())
        if !ok {
            return nil, errHandler.Error(nextSymbol, stateStack)
        }

        if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:343:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:343:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:343:44
	{
		_n, _err := _output.Write([]byte(` {
            stateStack = append(stateStack, action.ShiftItem(nextSymbol))

            _, err = symbolStack.Pop()
            if err != nil {
                return nil, err
            }
        } else if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:350:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:350:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:350:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:351:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:351:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:351:41
	{
		_n, _err := _output.Write([]byte(`
            stateStack, reduceSymbol, err = action.ReduceSymbol(
                reducer,
                stateStack)
            if err != nil {
                return nil, err
            }

            symbolStack.Push(reduceSymbol)
        } else if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:360:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:360:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:360:60
	{
		_n, _err := _output.Write([]byte(` {
            stateStack = append(stateStack, action.ShiftItem(nextSymbol))

            _, err = symbolStack.Pop()
            if err != nil {
                return nil, err
            }

            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:368:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:368:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:368:41
	{
		_n, _err := _output.Write([]byte(`
            stateStack, reduceSymbol, err = action.ReduceSymbol(
                reducer,
                stateStack)
            if err != nil {
                return nil, err
            }

            symbolStack.Push(reduceSymbol)
        } else if action.ActionType == `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:377:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:377:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:377:52
	{
		_n, _err := _output.Write([]byte(` {
            if len(stateStack) != 2 {
                panic("This should never happen")
            }
            return stateStack[1], nil
        } else {
            panic("Unknown action type: " + action.ActionType.String())
        }
    }
}

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:388:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:388:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:388:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:390:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:390:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:390:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:392:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:392:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:392:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:394:0
	for _, symbolName := range OrderedSymbolNames[3:] {
		// file.template:395:4
		term := Grammar.Terms[symbolName]
		// file.template:395:44
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:396:9
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:396:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:396:35
		{
			_n, _err := _output.Write([]byte(`:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:397:4
		if term.SymbolId == parser.LRCharacterToken {
			// file.template:398:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// file.template:407:10
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:408:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"file.template:408:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:408:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:409:13
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:410:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:410:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:410:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:412:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:414:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:23
	{
		_n, _err := _output.Write([]byte(`("?unknown symbol %d?", int(i))
    }
}

const (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:419:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:419:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:420:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:420:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:420:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:420:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:420:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:420:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:422:0
	for idx, term := range Grammar.NonTerminals {
		// file.template:422:48
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:423:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:423:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + len(Grammar.Terminals) + idx),
				"file.template:423:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:423:80
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:424:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:427:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:18
	{
		_n, _err := _output.Write([]byte(` int

const (
    // NOTE: error action is implicit
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:431:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:431:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:432:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:432:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:432:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:432:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:432:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:432:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:433:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:25
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:28
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:433:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:41
	{
		_n, _err := _output.Write([]byte(`(2)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:434:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:434:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:434:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:434:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:434:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:434:33
	{
		_n, _err := _output.Write([]byte(`(3)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:437:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:437:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:437:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:439:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:439:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:439:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:441:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:441:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:441:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:443:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:443:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:443:30
	{
		_n, _err := _output.Write([]byte(`:
        return "shift-and-reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:445:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:445:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:445:22
	{
		_n, _err := _output.Write([]byte(`:
        return "accept"
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:448:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:448:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:448:23
	{
		_n, _err := _output.Write([]byte(`("?Unknown action %d?", int(i))
    }
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:452:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:452:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:452:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:455:0
	clauseIdx := 1
	// file.template:456:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:457:4
		for _, clause := range rule.Clauses {
			// file.template:457:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:458:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"file.template:458:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"file.template:458:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:458:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:459:8
			clauseIdx += 1
		}
	}
	// file.template:461:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:464:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:464:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:464:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:466:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:467:4
		for _, clause := range rule.Clauses {
			// file.template:467:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:468:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:468:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:468:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:469:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:469:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:471:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:473:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:473:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:473:23
	{
		_n, _err := _output.Write([]byte(`("?unknown reduce type %d?", int(i))
    }
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:477:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:477:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:477:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:479:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:479:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:479:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:480:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:480:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:480:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:484:0
	for _, state := range States.OrderedStates {
		// file.template:484:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:485:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"file.template:485:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:485:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:485:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:486:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:489:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:489:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:489:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:490:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:490:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:490:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:492:13
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:492:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:492:31
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:494:0
	for _, valueType := range OrderedValueTypes {
		// file.template:495:4
		if valueType.Name == lr.Generic {
			// file.template:496:8
			continue
		}
		// file.template:497:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:498:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:498:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:498:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:498:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:498:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:499:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:502:0

	valueTerms := map[string][]*lr.Term{}
	for _, symbolName := range OrderedSymbolNames[2:] {
		term := Grammar.Terms[symbolName]
		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// file.template:510:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:511:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:511:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:512:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:512:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:512:37
	{
		_n, _err := _output.Write([]byte(`)
    if ok {
        return symbol, nil
    }

    symbol = &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:517:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:517:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:517:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:519:0
	for _, valueType := range OrderedValueTypes {
		// file.template:520:4

		consts := []string{}
		for _, term := range valueTerms[valueType.Name] {
			if !term.IsTerminal {
				break
			}

			consts = append(consts, term.CodeGenSymbolConst)
		}

		if len(consts) == 0 {
			continue
		}

		// file.template:535:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:537:4
		for idx, kconst := range consts {
			// file.template:538:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"file.template:538:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:538:8
			if idx != len(consts)-1 {
				// file.template:538:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:539:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:541:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:541:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:541:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:543:24
		{
			_n, _err := _template.writeValue(
				_output,
				(NewLocationError),
				"file.template:543:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:543:41
		{
			_n, _err := _output.Write([]byte(`(
                token.Loc(),
                "invalid value type for token %s. "+
                    "expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:546:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:546:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:546:53
		{
			_n, _err := _output.Write([]byte(`",
                token.Id())
        }
        symbol.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:549:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:549:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:549:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:550:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:552:20
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:552:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:552:37
	{
		_n, _err := _output.Write([]byte(`(
          token.Loc(),
          "unexpected token type: %s",
          token.Id())
    }
    return symbol, nil
}

func (s *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:560:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:560:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:40
	{
		_n, _err := _output.Write([]byte(` {
    return s.SymbolId_
}

func (s *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:564:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:564:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:564:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:565:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:565:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:565:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:567:0
	for _, field := range OrderedValueTypes {
		// file.template:568:4
		if field.Name == lr.Generic {
			// file.template:569:8
			continue
		}
		// file.template:571:4
		terms := valueTerms[field.Name]
		// file.template:572:4
		if len(terms) == 0 {
			// file.template:573:6
			continue
		}
		// file.template:574:12
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:576:4
		for idx, term := range terms {
			// file.template:577:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:577:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:578:8
			if idx != len(terms)-1 {
				// file.template:578:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:579:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:581:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"file.template:581:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:581:46
		{
			_n, _err := _output.Write([]byte(`).(locator)
        if ok {
            return loc.Loc()
        }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:585:8
	{
		_n, _err := _output.Write([]byte(`
    }
    return s.Generic_.Loc()
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:590:0
	if GenerateEndPos {
		// file.template:590:22
		{
			_n, _err := _output.Write([]byte(`
func (s *`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:591:9
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolType),
				"file.template:591:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:591:20
		{
			_n, _err := _output.Write([]byte(`) End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:591:28
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:591:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:591:41
		{
			_n, _err := _output.Write([]byte(` {
    type locator interface { End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:592:35
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:592:35")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:592:48
		{
			_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:594:0
		for _, field := range OrderedValueTypes {
			// file.template:595:4
			if field.Name == lr.Generic {
				// file.template:596:8
				continue
			}
			// file.template:598:4
			terms := valueTerms[field.Name]
			// file.template:599:4
			if len(terms) == 0 {
				// file.template:600:6
				continue
			}
			// file.template:601:12
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:603:4
			for idx, term := range terms {
				// file.template:604:0
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenSymbolConst),
						"file.template:604:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:605:8
				if idx != len(terms)-1 {
					// file.template:605:37
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:606:13
			{
				_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:608:33
			{
				_n, _err := _template.writeValue(
					_output,
					(field.Name),
					"file.template:608:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:608:46
			{
				_n, _err := _output.Write([]byte(`).(locator)
        if ok {
            return loc.End()
        }`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:612:8
		{
			_n, _err := _output.Write([]byte(`
    }
    return s.Generic_.End()
}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:616:8
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:618:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:618:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:618:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:619:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:620:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:620:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:620:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:623:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:623:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:623:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:623:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:623:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:623:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        token, err := stack.lexer.Next()
        if err != nil {
            if err != `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:627:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"file.template:627:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:627:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:628:28
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:628:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:628:45
	{
		_n, _err := _output.Write([]byte(`(
                  stack.lexer.CurrentLocation(),
                  "unexpected lex error: %w",
                  err)
            }
            token = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:633:20
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:633:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:633:38
	{
		_n, _err := _output.Write([]byte(`{
              `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:634:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:27
	{
		_n, _err := _output.Write([]byte(`: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:29
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:634:29")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:41
	{
		_n, _err := _output.Write([]byte(`,
              StartPos: stack.lexer.CurrentLocation(),
            }
        }
        item, err := NewSymbol(token)
        if err != nil {
            return nil, err
        }
        stack.top = append(stack.top, item)
    }
    return stack.top[len(stack.top)-1], nil
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:647:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:647:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:647:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:647:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:647:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:647:55
	{
		_n, _err := _output.Write([]byte(`) {
    stack.top = append(stack.top, symbol)
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:651:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:651:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:653:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:653:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:653:27
	{
		_n, _err := _output.Write([]byte(`("internal error: cannot pop an empty top")
    }
    ret := stack.top[len(stack.top)-1]
    stack.top = stack.top[:len(stack.top)-1]
    return ret, nil
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:660:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:661:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:663:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:663:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:663:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:666:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:666:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:668:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:669:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:669:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:669:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:671:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:671:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:671:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:672:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:672:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:672:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:675:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:675:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:675:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:676:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:676:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:676:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:676:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:676:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:676:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:679:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:679:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:679:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:680:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:680:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:680:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:681:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:681:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:681:20
	{
		_n, _err := _output.Write([]byte(`,
) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:683:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:683:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:683:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:684:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:684:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:684:16
	{
		_n, _err := _output.Write([]byte(`,
    error,
) {
    var err error
    symbol := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:688:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:688:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:688:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:690:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:691:4
		for _, clause := range rule.Clauses {
			// file.template:691:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:692:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:692:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:692:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:693:8
			if len(clause.Bindings) > 0 {
				// file.template:693:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:694:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:694:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:694:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:695:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:695:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:695:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:696:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:697:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"file.template:697:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:698:8
			if clause.Passthrough {
				// file.template:698:34
				{
					_n, _err := _output.Write([]byte(`
        //line `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:699:15
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation),
						"file.template:699:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:699:35
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:700:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:700:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:700:32
				{
					_n, _err := _output.Write([]byte(` = args[0].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:700:43
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Bindings[0].ValueType),
						"file.template:700:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:700:74
				{
					_n, _err := _output.Write([]byte(`
        err = nil`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:702:17
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:703:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:703:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:703:32
				{
					_n, _err := _output.Write([]byte(`, err = reducer.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:703:48
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.CodeGenReducerName),
						"file.template:703:48")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:703:76
				{
					_n, _err := _output.Write([]byte(`(`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:704:10
				for idx, term := range clause.Bindings {
					// file.template:704:54
					{
						_n, _err := _output.Write([]byte(`  args[`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:705:7
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:705:7")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:705:11
					{
						_n, _err := _output.Write([]byte(`].`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:705:13
					{
						_n, _err := _template.writeValue(
							_output,
							(term.ValueType),
							"file.template:705:13")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:706:14
					if idx != len(clause.Bindings)-1 {
						// file.template:706:51
						{
							_n, _err := _output.Write([]byte(`,`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:707:19
				{
					_n, _err := _output.Write([]byte(`)`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
	}
	// file.template:711:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        panic("Unknown reduce type: " + act.ReduceType.String())
    }

    if err != nil {
        err = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:717:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:717:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:717:21
	{
		_n, _err := _output.Write([]byte(`("unexpected %s reduce error: %w", act.ReduceType, err)
    }

    return stack, symbol, err
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:723:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:723:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:723:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:724:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:724:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:724:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:725:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:725:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:725:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:728:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:728:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:728:21
	{
		_n, _err := _output.Write([]byte(` struct{}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:730:6
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:730:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:730:22
	{
		_n, _err := _output.Write([]byte(`) Get(
  stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:731:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:731:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:731:22
	{
		_n, _err := _output.Write([]byte(`,
  symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:732:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:732:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:732:24
	{
		_n, _err := _output.Write([]byte(`,
) (
  `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:734:2
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:734:2")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:734:13
	{
		_n, _err := _output.Write([]byte(`,
  bool,
) {
  switch stateId {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:738:0
	for _, state := range States.OrderedStates {
		// file.template:738:47
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:739:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:739:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:739:28
		{
			_n, _err := _output.Write([]byte(`:
    switch symbolId {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:741:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:742:4
			if symbolName == lr.Wildcard {
				// file.template:743:6
				continue
			}
			// file.template:746:4

			symbol := Grammar.Terms[symbolName]
			nextState, ok := state.Goto[symbolName]

			// file.template:750:4
			if ok {
				// file.template:750:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:751:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:751:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:751:37
				{
					_n, _err := _output.Write([]byte(`:
      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:752:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ShiftAction),
						"file.template:752:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:39
				{
					_n, _err := _template.writeValue(
						_output,
						(nextState.CodeGenConst),
						"file.template:752:39")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:64
				{
					_n, _err := _output.Write([]byte(`, 0}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:6
				continue
			}
		}
		// file.template:757:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:758:4
			if symbolName == lr.Wildcard {
				// file.template:759:6
				continue
			}
			// file.template:762:4

			symbol := Grammar.Terms[symbolName]
			reduceItem, ok := state.ShiftAndReduce[symbolName]

			// file.template:766:4
			if ok {
				// file.template:766:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:767:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:767:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:767:37
				{
					_n, _err := _output.Write([]byte(`:
      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:768:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ShiftAndReduceAction),
						"file.template:768:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:46
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:51
				{
					_n, _err := _template.writeValue(
						_output,
						(reduceItem.Clause.CodeGenReducerNameConst),
						"file.template:768:51")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:95
				{
					_n, _err := _output.Write([]byte(`}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:769:6
				continue
			}
		}
		// file.template:773:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:774:4
			if symbolName == lr.Wildcard {
				// file.template:775:6
				continue
			}
			// file.template:778:4

			symbol := Grammar.Terms[symbolName]
			reduceItems, ok := state.Reduce[symbolName]

			// file.template:782:4
			if ok {
				// file.template:782:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:783:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:783:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:783:37
				{
					_n, _err := _output.Write([]byte(`:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:784:6
				for _, item := range reduceItems {
					// file.template:785:8
					if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
						// file.template:785:76
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:786:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:786:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:786:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:786:25
						{
							_n, _err := _template.writeValue(
								_output,
								(AcceptAction),
								"file.template:786:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:786:38
						{
							_n, _err := _output.Write([]byte(`, 0, 0}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					} else {
						// file.template:787:17
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:788:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:25
						{
							_n, _err := _template.writeValue(
								_output,
								(ReduceAction),
								"file.template:788:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:38
						{
							_n, _err := _output.Write([]byte(`, 0, `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:43
						{
							_n, _err := _template.writeValue(
								_output,
								(item.Clause.CodeGenReducerNameConst),
								"file.template:788:43")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:788:81
						{
							_n, _err := _output.Write([]byte(`}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:791:6
				continue
			}
		}
		// file.template:793:10
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:795:4

		reduceItems, ok := state.Reduce[lr.Wildcard]

		// file.template:798:4
		if ok {
			// file.template:798:14
			{
				_n, _err := _output.Write([]byte(`
    default:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:800:6
			for _, item := range reduceItems {
				// file.template:800:44
				{
					_n, _err := _output.Write([]byte(`      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:801:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ReduceAction),
						"file.template:801:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:38
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:43
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Clause.CodeGenReducerNameConst),
						"file.template:801:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:801:81
				{
					_n, _err := _output.Write([]byte(`}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:803:13
		{
			_n, _err := _output.Write([]byte(`    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:805:8
	{
		_n, _err := _output.Write([]byte(`
  }

  return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:808:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:808:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:808:20
	{
		_n, _err := _output.Write([]byte(`{}, false
}

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:811:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:811:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:811:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:811:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:811:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:811:35
	{
		_n, _err := _output.Write([]byte(`{}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:814:0
	{
		_n, _err := (&debug_template.File{
			OrderedSymbolNames:        OrderedSymbolNames,
			States:                    States,
			OutputDebugNonKernelItems: OutputDebugNonKernelItems,
		}).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:820:3
	{
		_n, _err := _output.Write([]byte(`*/
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
