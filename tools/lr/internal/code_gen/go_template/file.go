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
	EOF := _template.EOF
	OrderedSymbolNames := _template.OrderedSymbolNames
	Grammar := _template.Grammar
	States := _template.States
	OrderedValueTypes := _template.OrderedValueTypes
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems
	GenerateEndPos := _template.GenerateEndPos

	// file.template:77:0
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:77:31
	{
		_n, _err := _template.writeValue(
			_output,
			(Grammar.Source),
			"file.template:77:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:77:48
	{
		_n, _err := _output.Write([]byte(`

package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:79:8
	{
		_n, _err := _template.writeValue(
			_output,
			(Package),
			"file.template:79:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:79:16
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:81:0
	{
		_n, _err := (Imports).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:81:20
	{
		_n, _err := _output.Write([]byte(`
type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:83:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:83:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:83:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:86:0

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

	// file.template:112:0
	for _, term := range Grammar.Terminals {
		// file.template:113:4
		if term.SymbolId != parser.LRIdentifierToken {
			// file.template:113:53
			{
				_n, _err := _output.Write([]byte(`
    // char token `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:18
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:114:18")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:114:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:49
			{
				_n, _err := _template.writeValue(
					_output,
					(charTermToIds[term.Name]),
					"file.template:114:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:114:76
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:117:0
	for _, term := range Grammar.Terminals {
		// file.template:118:4
		if term.SymbolId == parser.LRIdentifierToken {
			// file.template:118:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:119:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:119:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId()),
					"file.template:119:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:119:60
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:121:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:124:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:124:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:124:18
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:124:21
	{
		_n, _err := _template.writeValue(
			_output,
			(RealLocationType),
			"file.template:124:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:124:38
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:126:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:15
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:127:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:128:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:0
	if GenerateEndPos {
		// file.template:129:22
		{
			_n, _err := _output.Write([]byte(`
    End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:130:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:130:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:131:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:134:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:134:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:134:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:135:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:17
	{
		_n, _err := _output.Write([]byte(`
    StartPos `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:13
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:136:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:137:0
	if GenerateEndPos {
		// file.template:137:22
		{
			_n, _err := _output.Write([]byte(`
    EndPos `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:138:11
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:138:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:139:8
	{
		_n, _err := _output.Write([]byte(`
}

func (t `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:8
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:142:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:26
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:33
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:142:33")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:46
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:58
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:142:58")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:71
	{
		_n, _err := _output.Write([]byte(` }

func (t `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:144:8
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:144:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:144:26
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:144:34
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:144:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:144:47
	{
		_n, _err := _output.Write([]byte(` { return t.StartPos }
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:0
	if GenerateEndPos {
		// file.template:146:22
		{
			_n, _err := _output.Write([]byte(`
func (t `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:147:8
		{
			_n, _err := _template.writeValue(
				_output,
				(GenericSymbolType),
				"file.template:147:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:147:26
		{
			_n, _err := _output.Write([]byte(`) End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:147:34
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:147:34")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:147:47
		{
			_n, _err := _output.Write([]byte(` { return t.EndPos }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:148:8
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:150:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:150:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:150:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:152:55
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:152:55")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:152:73
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:153:12
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:153:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:153:22
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:155:22
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:155:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:155:35
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:158:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:159:4
		if rule.NumReducerClauses == 0 {
			// file.template:160:8
			continue
		}
		// file.template:161:12
		{
			_n, _err := _output.Write([]byte(`
type `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:162:5
		{
			_n, _err := _template.writeValue(
				_output,
				(rule.CodeGenReducerInterface),
				"file.template:162:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:162:36
		{
			_n, _err := _output.Write([]byte(` interface {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:163:4
		for clauseIdx, clause := range rule.Clauses {
			// file.template:164:8
			if clause.Passthrough {
				// file.template:165:10
				continue
			}
			// file.template:167:8
			if clauseIdx > 0 {
				// file.template:167:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:171:8
			if clause.Label == "" {
				// file.template:171:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:172:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:172:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:172:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:172:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:172:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:172:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:173:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:174:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:174:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:174:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:174:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:177:8
			paramNameCount := map[string]int{}
			// file.template:177:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:178:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:178:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:178:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:179:8
			for termIdx, term := range clause.Bindings {
				// file.template:181:12

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

				// file.template:204:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:204:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:204:10
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:204:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"file.template:204:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:204:30
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"file.template:204:30")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:205:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:206:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"file.template:206:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:206:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:207:13
		{
			_n, _err := _output.Write([]byte(`}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:209:7
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:211:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:212:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:213:4
		if len(rule.Clauses) == 0 {
			// file.template:214:8
			continue
		}
		// file.template:216:2
		if rule.NumReducerClauses > 0 {
			// file.template:216:36
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:217:2
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenReducerInterface),
					"file.template:217:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:219:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:222:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:222:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:222:20
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:223:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:223:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:223:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:226:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:226:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:226:27
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:228:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:28
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:46
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:228:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:56
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:64
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:228:64")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:228:74
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:229:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:18
	{
		_n, _err := _output.Write([]byte(`(
        "Syntax error: unexpected symbol %v. Expecting %v (%v)",
        nextToken.Id(),
        `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:232:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:232:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:232:30
	{
		_n, _err := _output.Write([]byte(`(stack[len(stack)-1].StateId),
        nextToken.Loc())
}

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:236:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:27
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:31
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:236:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:43
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:47
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:236:47")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:236:60
	{
		_n, _err := _output.Write([]byte(` {
  switch id {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:238:0
	for _, state := range States.OrderedStates {
		// file.template:239:2

		_, ok := state.Reduce[lr.Wildcard]
		if ok {
			continue
		}

		// file.template:244:4
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:245:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:245:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:245:28
		{
			_n, _err := _output.Write([]byte(`:
    return []`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:246:13
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:246:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:246:26
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:247:4
		for _, term := range Grammar.Terminals {
			// file.template:248:6

			_, foundGoto := state.Goto[term.Name]
			_, foundReduce := state.Reduce[term.Name]
			_, foundShiftAndReduce := state.ShiftAndReduce[term.Name]

			if !foundGoto && !foundReduce && !foundShiftAndReduce {
				continue
			}

			// file.template:257:9
			{
				_n, _err := _output.Write([]byte(`      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:258:6
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:258:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:258:32
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:259:12
		{
			_n, _err := _output.Write([]byte(`
    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:261:8
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
	// file.template:267:0
	for idx, start := range Grammar.Starts {
		// file.template:268:4

		parseSuffix := ""
		if len(Grammar.Starts) > 1 {
			parseSuffix = codegen.SnakeToCamel(start.Name)
		}

		// file.template:275:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:276:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:276:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:37
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:44
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:276:44")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:54
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:64
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:276:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:276:76
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:278:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:278:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:278:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:280:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:29
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:280:29")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:43
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
		// file.template:283:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandlerType),
				"file.template:283:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:283:30
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:286:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:286:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:286:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:286:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:286:37
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:287:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:287:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:287:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:288:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:288:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:288:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:289:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandlerType),
				"file.template:289:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:289:30
		{
			_n, _err := _output.Write([]byte(`) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:290:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:290:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:290:24
		{
			_n, _err := _output.Write([]byte(`,
    error) {

    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:293:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParseFunc),
				"file.template:293:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:293:35
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:293:64
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:293:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:294:47
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:296:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:296:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:296:42
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
		// file.template:299:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"file.template:299:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:299:34
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
	// file.template:301:7
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
	// file.template:308:5
	{
		_n, _err := _template.writeValue(
			_output,
			(InternalParseFunc),
			"file.template:308:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:308:23
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:309:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:309:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:309:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:310:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:311:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:311:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:311:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:312:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:312:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:312:27
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:313:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:313:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:313:19
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    stateStack := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:316:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:316:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:316:28
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
	// file.template:319:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:319:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:319:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:322:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:322:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:322:36
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
	// file.template:330:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:330:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:330:34
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
	// file.template:337:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:337:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:337:44
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
	// file.template:344:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:344:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:344:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:345:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:345:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:345:41
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
	// file.template:354:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:354:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:354:60
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
	// file.template:362:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:362:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:362:41
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
	// file.template:371:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:371:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:371:52
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
	// file.template:382:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:382:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:382:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:384:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:386:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:386:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:386:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:388:0
	for _, symbolName := range OrderedSymbolNames[3:] {
		// file.template:389:4
		term := Grammar.Terms[symbolName]
		// file.template:389:44
		{
			_n, _err := _output.Write([]byte(`
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
				(term.CodeGenSymbolConst),
				"file.template:390:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:390:35
		{
			_n, _err := _output.Write([]byte(`:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:391:4
		if term.SymbolId == parser.LRCharacterToken {
			// file.template:392:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// file.template:401:10
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:402:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"file.template:402:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:402:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:403:13
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:404:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:404:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:404:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:406:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:408:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:408:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:408:23
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
	// file.template:413:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:413:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:413:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:414:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:414:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:416:0
	for idx, term := range Grammar.NonTerminals {
		// file.template:416:48
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:417:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:417:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + len(Grammar.Terminals) + idx),
				"file.template:417:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:417:80
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:418:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:421:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:421:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:421:18
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
	// file.template:425:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:425:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:425:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:425:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:425:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:425:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:426:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:426:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:426:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:426:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:426:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:426:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:427:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:25
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:28
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:427:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:41
	{
		_n, _err := _output.Write([]byte(`(2)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:428:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:428:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:33
	{
		_n, _err := _output.Write([]byte(`(3)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:431:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:431:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:433:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:435:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:435:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:435:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:437:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:437:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:437:30
	{
		_n, _err := _output.Write([]byte(`:
        return "shift-and-reduce"
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
			(AcceptAction),
			"file.template:439:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:439:22
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
	// file.template:442:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:442:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:442:23
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
	// file.template:446:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:446:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:446:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:449:0
	clauseIdx := 1
	// file.template:450:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:451:4
		for _, clause := range rule.Clauses {
			// file.template:451:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:452:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"file.template:452:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"file.template:452:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:452:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:453:8
			clauseIdx += 1
		}
	}
	// file.template:455:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:458:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:458:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:458:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:460:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:461:4
		for _, clause := range rule.Clauses {
			// file.template:461:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:462:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:462:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:462:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:463:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:463:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:463:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:465:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:467:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:467:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:467:23
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
	// file.template:471:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:471:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:471:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:473:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:473:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:473:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:474:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:474:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:474:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:478:0
	for _, state := range States.OrderedStates {
		// file.template:478:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:479:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"file.template:479:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:479:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:479:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:480:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:483:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:483:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:483:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:484:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:484:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:484:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:486:13
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:486:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:486:31
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:488:0
	for _, valueType := range OrderedValueTypes {
		// file.template:489:4
		if valueType.Name == lr.Generic {
			// file.template:490:8
			continue
		}
		// file.template:491:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:492:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:492:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:492:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:492:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:492:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:493:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:496:0

	valueTerms := map[string][]*lr.Term{}
	for _, symbolName := range OrderedSymbolNames[2:] {
		term := Grammar.Terms[symbolName]
		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// file.template:504:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:505:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:505:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:505:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:505:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:505:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:505:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:506:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:37
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
	// file.template:511:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:511:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:513:0
	for _, valueType := range OrderedValueTypes {
		// file.template:514:4

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

		// file.template:529:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:531:4
		for idx, kconst := range consts {
			// file.template:532:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"file.template:532:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:532:8
			if idx != len(consts)-1 {
				// file.template:532:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:533:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:535:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:535:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:535:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:537:24
		{
			_n, _err := _template.writeValue(
				_output,
				(Errorf),
				"file.template:537:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:537:31
		{
			_n, _err := _output.Write([]byte(`(
                "Invalid value type for token %s.  "+
                    "Expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:539:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:539:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:539:53
		{
			_n, _err := _output.Write([]byte(` (%v)",
                token.Id(),
                token.Loc())
        }
        symbol.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:543:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:543:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:543:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:544:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:546:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:546:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:546:27
	{
		_n, _err := _output.Write([]byte(`("Unexpected token type: %s", symbol.Id())
    }
    return symbol, nil
}

func (s *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:551:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:551:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:551:40
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
	// file.template:555:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:555:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:555:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:555:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:555:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:555:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:556:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:558:0
	for _, field := range OrderedValueTypes {
		// file.template:559:4
		if field.Name == lr.Generic {
			// file.template:560:8
			continue
		}
		// file.template:562:4
		terms := valueTerms[field.Name]
		// file.template:563:4
		if len(terms) == 0 {
			// file.template:564:6
			continue
		}
		// file.template:565:12
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:567:4
		for idx, term := range terms {
			// file.template:568:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:568:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:569:8
			if idx != len(terms)-1 {
				// file.template:569:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:570:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:572:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"file.template:572:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:572:46
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
	// file.template:576:8
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
	// file.template:581:0
	if GenerateEndPos {
		// file.template:581:22
		{
			_n, _err := _output.Write([]byte(`
func (s *`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:582:9
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolType),
				"file.template:582:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:582:20
		{
			_n, _err := _output.Write([]byte(`) End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:582:28
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:582:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:582:41
		{
			_n, _err := _output.Write([]byte(` {
    type locator interface { End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:583:35
		{
			_n, _err := _template.writeValue(
				_output,
				(LocationType),
				"file.template:583:35")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:583:48
		{
			_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:585:0
		for _, field := range OrderedValueTypes {
			// file.template:586:4
			if field.Name == lr.Generic {
				// file.template:587:8
				continue
			}
			// file.template:589:4
			terms := valueTerms[field.Name]
			// file.template:589:42
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:591:4
			for idx, term := range terms {
				// file.template:592:0
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenSymbolConst),
						"file.template:592:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:593:8
				if idx != len(terms)-1 {
					// file.template:593:37
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:594:13
			{
				_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:596:33
			{
				_n, _err := _template.writeValue(
					_output,
					(field.Name),
					"file.template:596:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:596:46
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
		// file.template:600:8
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
	// file.template:604:8
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:606:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:606:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:606:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:607:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:607:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:607:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:608:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:608:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:608:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:611:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:611:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:611:50
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
	// file.template:615:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"file.template:615:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:615:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:616:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:35
	{
		_n, _err := _output.Write([]byte(`("Unexpected lex error: %s", err)
            }
            token = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:618:20
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:618:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:618:38
	{
		_n, _err := _output.Write([]byte(`{
              `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:619:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:27
	{
		_n, _err := _output.Write([]byte(`: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:29
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:619:29")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:41
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
	// file.template:632:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:632:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:632:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:55
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
	// file.template:636:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:636:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:636:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:636:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:636:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:636:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:638:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:638:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:638:27
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
	// file.template:645:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:645:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:645:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:646:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:646:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:646:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:648:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:648:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:648:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:651:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:651:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:651:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:653:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:653:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:653:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:654:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:654:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:654:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:656:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:656:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:656:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:657:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:657:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:657:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:660:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:660:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:660:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:661:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:661:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:661:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:664:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:664:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:664:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:665:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:665:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:665:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:666:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:666:20
	{
		_n, _err := _output.Write([]byte(`) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:667:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:667:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:667:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:668:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:16
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    var err error
    symbol := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:672:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:672:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:672:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:674:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:675:4
		for _, clause := range rule.Clauses {
			// file.template:675:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:676:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:676:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:676:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:677:8
			if len(clause.Bindings) > 0 {
				// file.template:677:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:678:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:678:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:678:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:679:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:679:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:679:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:680:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:681:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"file.template:681:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:682:8
			if clause.Passthrough {
				// file.template:682:34
				{
					_n, _err := _output.Write([]byte(`
        //line `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:683:15
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation),
						"file.template:683:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:683:35
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:684:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:32
				{
					_n, _err := _output.Write([]byte(` = args[0].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:43
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Bindings[0].ValueType),
						"file.template:684:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:74
				{
					_n, _err := _output.Write([]byte(`
        err = nil`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:686:17
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:687:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:687:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:687:32
				{
					_n, _err := _output.Write([]byte(`, err = reducer.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:687:48
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.CodeGenReducerName),
						"file.template:687:48")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:687:76
				{
					_n, _err := _output.Write([]byte(`(`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:688:10
				for idx, term := range clause.Bindings {
					// file.template:688:54
					{
						_n, _err := _output.Write([]byte(`  args[`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:689:7
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:689:7")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:689:11
					{
						_n, _err := _output.Write([]byte(`].`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:689:13
					{
						_n, _err := _template.writeValue(
							_output,
							(term.ValueType),
							"file.template:689:13")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:690:14
					if idx != len(clause.Bindings)-1 {
						// file.template:690:51
						{
							_n, _err := _output.Write([]byte(`,`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:691:19
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
	// file.template:695:8
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
	// file.template:701:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:701:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:701:21
	{
		_n, _err := _output.Write([]byte(`("Unexpected %s reduce error: %s", act.ReduceType, err)
    }

    return stack, symbol, err
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:707:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:707:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:707:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:708:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:708:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:708:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:709:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:709:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:709:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:712:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:712:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:712:21
	{
		_n, _err := _output.Write([]byte(` struct{}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:714:6
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:714:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:714:22
	{
		_n, _err := _output.Write([]byte(`) Get(
  stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:715:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:715:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:715:22
	{
		_n, _err := _output.Write([]byte(`,
  symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:716:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:716:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:716:24
	{
		_n, _err := _output.Write([]byte(`,
) (
  `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:718:2
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:718:2")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:718:13
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
	// file.template:722:0
	for _, state := range States.OrderedStates {
		// file.template:722:47
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:723:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:723:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:723:28
		{
			_n, _err := _output.Write([]byte(`:
    switch symbolId {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:725:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:726:4
			if symbolName == lr.Wildcard {
				// file.template:727:6
				continue
			}
			// file.template:730:4

			symbol := Grammar.Terms[symbolName]
			nextState, ok := state.Goto[symbolName]

			// file.template:734:4
			if ok {
				// file.template:734:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:735:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:735:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:735:37
				{
					_n, _err := _output.Write([]byte(`:
      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:736:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ShiftAction),
						"file.template:736:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:39
				{
					_n, _err := _template.writeValue(
						_output,
						(nextState.CodeGenConst),
						"file.template:736:39")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:64
				{
					_n, _err := _output.Write([]byte(`, 0}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:6
				continue
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
			reduceItem, ok := state.ShiftAndReduce[symbolName]

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
						(ShiftAndReduceAction),
						"file.template:752:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:46
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:51
				{
					_n, _err := _template.writeValue(
						_output,
						(reduceItem.Clause.CodeGenReducerNameConst),
						"file.template:752:51")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:95
				{
					_n, _err := _output.Write([]byte(`}, true`))
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
			reduceItems, ok := state.Reduce[symbolName]

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
					_n, _err := _output.Write([]byte(`:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:6
				for _, item := range reduceItems {
					// file.template:769:8
					if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
						// file.template:769:76
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:770:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:770:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:770:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:770:25
						{
							_n, _err := _template.writeValue(
								_output,
								(AcceptAction),
								"file.template:770:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:770:38
						{
							_n, _err := _output.Write([]byte(`, 0, 0}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					} else {
						// file.template:771:17
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:772:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:25
						{
							_n, _err := _template.writeValue(
								_output,
								(ReduceAction),
								"file.template:772:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:38
						{
							_n, _err := _output.Write([]byte(`, 0, `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:43
						{
							_n, _err := _template.writeValue(
								_output,
								(item.Clause.CodeGenReducerNameConst),
								"file.template:772:43")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:772:81
						{
							_n, _err := _output.Write([]byte(`}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:775:6
				continue
			}
		}
		// file.template:777:10
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:779:4

		reduceItems, ok := state.Reduce[lr.Wildcard]

		// file.template:782:4
		if ok {
			// file.template:782:14
			{
				_n, _err := _output.Write([]byte(`
    default:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:784:6
			for _, item := range reduceItems {
				// file.template:784:44
				{
					_n, _err := _output.Write([]byte(`      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:785:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ReduceAction),
						"file.template:785:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:38
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:43
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Clause.CodeGenReducerNameConst),
						"file.template:785:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:785:81
				{
					_n, _err := _output.Write([]byte(`}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:787:13
		{
			_n, _err := _output.Write([]byte(`    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:789:8
	{
		_n, _err := _output.Write([]byte(`
  }

  return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:792:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:792:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:792:20
	{
		_n, _err := _output.Write([]byte(`{}, false
}

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:795:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:795:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:795:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:795:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:795:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:795:35
	{
		_n, _err := _output.Write([]byte(`{}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:798:0
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
	// file.template:804:3
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
