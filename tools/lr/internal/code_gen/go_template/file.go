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
	SortSlice                 interface{}
	Sprintf                   interface{}
	Errorf                    interface{}
	EOF                       interface{}
	OrderedSymbolNames        []string
	Grammar                   *lr.Grammar
	States                    *lr.LRStates
	OrderedValueTypes         lr.ParamList
	OutputDebugNonKernelItems bool
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
	SortSlice := _template.SortSlice
	Sprintf := _template.Sprintf
	Errorf := _template.Errorf
	EOF := _template.EOF
	OrderedSymbolNames := _template.OrderedSymbolNames
	Grammar := _template.Grammar
	States := _template.States
	OrderedValueTypes := _template.OrderedValueTypes
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems

	// file.template:75:0
	{
		_n, _err := _output.Write([]byte(`// Auto-generated from source: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:75:31
	{
		_n, _err := _template.writeValue(
			_output,
			(Grammar.Source),
			"file.template:75:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:75:48
	{
		_n, _err := _output.Write([]byte(`

package `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:77:8
	{
		_n, _err := _template.writeValue(
			_output,
			(Package),
			"file.template:77:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:77:16
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:79:0
	{
		_n, _err := (Imports).WriteTo(_output)
		_numWritten += _n
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:79:20
	{
		_n, _err := _output.Write([]byte(`
type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:81:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:81:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:81:18
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:84:0

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

	// file.template:110:0
	for _, term := range Grammar.Terminals {
		// file.template:111:4
		if term.SymbolId != parser.LRIdentifierToken {
			// file.template:111:53
			{
				_n, _err := _output.Write([]byte(`
    // char token `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:18
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:112:18")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:112:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:49
			{
				_n, _err := _template.writeValue(
					_output,
					(charTermToIds[term.Name]),
					"file.template:112:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:112:76
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:115:0
	for _, term := range Grammar.Terminals {
		// file.template:116:4
		if term.SymbolId == parser.LRIdentifierToken {
			// file.template:116:53
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:4
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:117:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:30
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:33
			{
				_n, _err := _template.writeValue(
					_output,
					(SymbolIdType),
					"file.template:117:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:48
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:49
			{
				_n, _err := _template.writeValue(
					_output,
					(nextId()),
					"file.template:117:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:117:60
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:119:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:122:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:122:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:122:18
	{
		_n, _err := _output.Write([]byte(` struct {
    FileName string
    Line int
    Column int
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:128:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:129:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:132:8
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:132:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:132:21
	{
		_n, _err := _output.Write([]byte(`) ShortString() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:133:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:19
	{
		_n, _err := _output.Write([]byte(`("%v:%v", l.Line, l.Column)
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:136:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:15
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:137:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:137:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:137:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:138:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:138:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:138:23
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:141:5
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:141:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:141:23
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:142:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:142:17
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:4
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:143:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:143:17
	{
		_n, _err := _output.Write([]byte(`
}

func (t `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:8
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:146:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:26
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:33
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:146:33")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:46
	{
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:58
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:146:58")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:146:71
	{
		_n, _err := _output.Write([]byte(` }

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
		_n, _err := _output.Write([]byte(`) Loc() `))
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
		_n, _err := _output.Write([]byte(` { return t.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:148:59
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:148:59")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:148:72
	{
		_n, _err := _output.Write([]byte(` }

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

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:158:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:158:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:158:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:159:0
	firstRule := true
	// file.template:160:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:161:4
		if len(rule.Clauses) == 0 {
			// file.template:162:8
			continue
		}
		// file.template:165:4
		if !firstRule {
			// file.template:165:22
			{
				_n, _err := _output.Write([]byte(`
`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:168:4
		firstRule = false
		// file.template:170:4
		for clauseIdx, clause := range rule.Clauses {
			// file.template:171:8
			if clauseIdx > 0 {
				// file.template:171:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:175:8
			if clause.Label == "" {
				// file.template:175:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:176:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:176:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:176:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:176:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:176:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:176:55
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:177:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.LRLocation.ShortString()),
						"file.template:178:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:41
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:43
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:178:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:55
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:59
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:178:59")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:178:74
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:181:8
			paramNameCount := map[string]int{}
			// file.template:181:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:182:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:182:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:182:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:183:8
			for termIdx, term := range clause.Bindings {
				// file.template:185:12

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

				// file.template:208:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:208:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:208:10
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:208:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"file.template:208:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:208:30
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"file.template:208:30")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:209:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:210:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"file.template:210:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:210:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:212:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:215:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:215:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:215:20
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:216:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:216:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:216:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:216:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:216:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:216:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:219:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:27
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:221:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:28
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:46
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:221:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:56
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:64
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:221:64")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:74
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:222:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:222:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:222:18
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
	// file.template:225:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:225:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:225:30
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
	// file.template:229:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:229:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:27
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:31
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:229:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:43
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:47
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:229:47")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:229:60
	{
		_n, _err := _output.Write([]byte(` {
    result := []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:230:16
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:230:16")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:230:29
	{
		_n, _err := _output.Write([]byte(`{}
    for key, _ := range `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:231:24
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:231:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:231:36
	{
		_n, _err := _output.Write([]byte(` {
        if key.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:232:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:232:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:232:27
	{
		_n, _err := _output.Write([]byte(` != id {
            continue
        }
        result = append(result, key.`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:235:36
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:235:36")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:235:49
	{
		_n, _err := _output.Write([]byte(`)
    }

    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:238:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SortSlice),
			"file.template:238:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:238:14
	{
		_n, _err := _output.Write([]byte(`(result, func(i int, j int) bool {return result[i] < result[j]})
    return result
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:242:0
	for idx, start := range Grammar.Starts {
		// file.template:243:4

		parseSuffix := ""
		if len(Grammar.Starts) > 1 {
			parseSuffix = codegen.SnakeToCamel(start.Name)
		}

		// file.template:250:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:251:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:251:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:37
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:44
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:251:44")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:54
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:64
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:251:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:251:76
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:253:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:253:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:253:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:255:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:255:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:255:29
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:255:29")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:255:43
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
		// file.template:258:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandlerType),
				"file.template:258:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:258:30
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:261:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:261:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:261:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:261:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:261:37
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:262:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:263:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:263:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:263:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:264:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandlerType),
				"file.template:264:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:264:30
		{
			_n, _err := _output.Write([]byte(`) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:265:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:265:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:265:24
		{
			_n, _err := _output.Write([]byte(`,
    error) {

    item, err := `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:268:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParseFunc),
				"file.template:268:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:268:35
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:268:64
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:268:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:269:47
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:271:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:271:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:271:42
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
		// file.template:274:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"file.template:274:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:274:34
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
	// file.template:276:7
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
	// file.template:283:5
	{
		_n, _err := _template.writeValue(
			_output,
			(InternalParseFunc),
			"file.template:283:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:283:23
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:284:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:284:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:284:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:285:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:285:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:285:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:286:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:286:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:286:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:287:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:287:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:287:27
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:288:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:288:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:288:19
	{
		_n, _err := _output.Write([]byte(`,
    error) {

    stateStack := `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:291:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:291:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:291:28
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
	// file.template:294:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:294:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:294:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:297:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:297:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:297:36
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
	// file.template:305:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:305:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:305:34
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
	// file.template:312:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:312:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:312:44
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
	// file.template:319:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:319:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:319:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:320:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:320:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:320:41
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
	// file.template:329:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:329:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:329:52
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
	// file.template:340:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:340:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:340:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:342:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:344:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:344:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:344:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:346:0
	for _, symbolName := range OrderedSymbolNames[3:] {
		// file.template:347:4
		term := Grammar.Terms[symbolName]
		// file.template:347:44
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:348:9
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:348:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:348:35
		{
			_n, _err := _output.Write([]byte(`:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:349:4
		if term.SymbolId == parser.LRCharacterToken {
			// file.template:350:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// file.template:359:10
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:360:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"file.template:360:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:360:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:361:13
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:362:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:362:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:362:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:364:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:366:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:366:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:366:23
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
	// file.template:371:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:371:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:371:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:371:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:371:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:371:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:372:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:372:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:374:0
	for idx, term := range Grammar.NonTerminals {
		// file.template:374:48
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:375:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:375:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + len(Grammar.Terminals) + idx),
				"file.template:375:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:375:80
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:376:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:379:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:379:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:379:18
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
	// file.template:383:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:383:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:383:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:383:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:383:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:383:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:384:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:384:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:384:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:385:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:385:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:385:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:385:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:385:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:385:33
	{
		_n, _err := _output.Write([]byte(`(2)
)

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
			(ActionIdType),
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
			(ShiftAction),
			"file.template:390:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:390:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
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
			(ReduceAction),
			"file.template:392:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:392:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:394:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:394:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:394:22
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
	// file.template:397:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:397:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:397:23
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
	// file.template:401:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:401:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:401:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:404:0
	clauseIdx := 1
	// file.template:405:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:406:4
		for _, clause := range rule.Clauses {
			// file.template:406:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:407:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"file.template:407:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"file.template:407:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:407:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:408:8
			clauseIdx += 1
		}
	}
	// file.template:410:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:413:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:413:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:415:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:416:4
		for _, clause := range rule.Clauses {
			// file.template:416:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:417:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:417:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:417:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:418:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:418:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:418:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:420:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:422:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:422:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:422:23
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
	// file.template:426:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:426:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:426:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:428:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:428:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:429:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:429:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:429:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:433:0
	for _, state := range States.OrderedStates {
		// file.template:433:47
		{
			_n, _err := _output.Write([]byte(`
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
				(state.CodeGenConst),
				"file.template:434:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:434:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:434:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"file.template:434:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:434:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:434:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:434:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:434:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:435:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:438:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:438:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:438:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:439:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:439:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:439:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:441:13
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:441:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:441:31
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:443:0
	for _, valueType := range OrderedValueTypes {
		// file.template:444:4
		if valueType.Name == lr.Generic {
			// file.template:445:8
			continue
		}
		// file.template:446:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:447:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:447:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:447:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:447:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:447:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:448:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:451:0

	valueTerms := map[string][]*lr.Term{}
	for _, symbolName := range OrderedSymbolNames[2:] {
		term := Grammar.Terms[symbolName]
		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// file.template:459:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:460:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:460:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:460:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:460:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:460:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:460:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:461:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:461:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:461:37
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
	// file.template:466:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:466:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:466:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:468:0
	for _, valueType := range OrderedValueTypes {
		// file.template:469:4

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

		// file.template:484:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:486:4
		for idx, kconst := range consts {
			// file.template:487:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"file.template:487:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:487:8
			if idx != len(consts)-1 {
				// file.template:487:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:488:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:490:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:490:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:490:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:492:24
		{
			_n, _err := _template.writeValue(
				_output,
				(Errorf),
				"file.template:492:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:492:31
		{
			_n, _err := _output.Write([]byte(`(
                "Invalid value type for token %s.  "+
                    "Expecting `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:494:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:494:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:494:53
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
		// file.template:498:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:498:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:498:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:499:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:501:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:501:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:501:27
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
	// file.template:506:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:506:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:506:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:506:40
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
	// file.template:510:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:510:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:510:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:510:28
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:510:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:510:41
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:35
	{
		_n, _err := _template.writeValue(
			_output,
			(LocationType),
			"file.template:511:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:511:48
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:513:0
	for _, field := range OrderedValueTypes {
		// file.template:514:4
		if field.Name == lr.Generic {
			// file.template:515:8
			continue
		}
		// file.template:517:4
		terms := valueTerms[field.Name]
		// file.template:517:42
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:519:4
		for idx, term := range terms {
			// file.template:520:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:520:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:521:8
			if idx != len(terms)-1 {
				// file.template:521:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:522:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:524:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"file.template:524:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:524:46
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
	// file.template:528:8
	{
		_n, _err := _output.Write([]byte(`
    }
    return s.Generic_.Loc()
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:533:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:533:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:533:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:534:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:534:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:534:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:535:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:535:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:535:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:538:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:538:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:538:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:538:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:538:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:538:50
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
	// file.template:542:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"file.template:542:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:542:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:543:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:543:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:543:35
	{
		_n, _err := _output.Write([]byte(`("Unexpected lex error: %s", err)
            }
            token = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:20
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbolType),
			"file.template:545:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:38
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:39
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:545:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:545:51
	{
		_n, _err := _output.Write([]byte(`, stack.lexer.CurrentLocation()}
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
	// file.template:556:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:556:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:29
	{
		_n, _err := _output.Write([]byte(`) Push(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:44
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:556:44")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:556:55
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
	// file.template:560:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:560:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:29
	{
		_n, _err := _output.Write([]byte(`) Pop() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:560:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:560:50
	{
		_n, _err := _output.Write([]byte(`, error) {
    if len(stack.top) == 0 {
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:562:20
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:562:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:562:27
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
	// file.template:569:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:569:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:569:19
	{
		_n, _err := _output.Write([]byte(` struct {
    StateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:570:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:570:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:570:24
	{
		_n, _err := _output.Write([]byte(`

    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:572:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:572:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:572:16
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:575:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:575:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:575:15
	{
		_n, _err := _output.Write([]byte(` []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:575:19
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:575:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:575:33
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:577:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:577:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:577:16
	{
		_n, _err := _output.Write([]byte(` struct {
    ActionType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:578:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:578:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:578:28
	{
		_n, _err := _output.Write([]byte(`

    ShiftStateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:580:17
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:580:17")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:580:29
	{
		_n, _err := _output.Write([]byte(`
    ReduceType `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:581:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:581:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:581:26
	{
		_n, _err := _output.Write([]byte(`
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:584:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:22
	{
		_n, _err := _output.Write([]byte(`) ShiftItem(symbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:42
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:584:42")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:53
	{
		_n, _err := _output.Write([]byte(`) *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:56
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:584:56")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:584:70
	{
		_n, _err := _output.Write([]byte(` {
    return &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:585:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:585:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:585:26
	{
		_n, _err := _output.Write([]byte(`{StateId: act.ShiftStateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:585:54
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:585:54")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:585:65
	{
		_n, _err := _output.Write([]byte(`: symbol}
}

func (act *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:588:11
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:588:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:588:22
	{
		_n, _err := _output.Write([]byte(`) ReduceSymbol(
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:589:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:589:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:589:24
	{
		_n, _err := _output.Write([]byte(`,
    stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:590:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:590:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:590:20
	{
		_n, _err := _output.Write([]byte(`) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:591:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:591:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:591:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:592:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:592:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:592:16
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
	// file.template:596:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:596:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:596:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:598:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:599:4
		for _, clause := range rule.Clauses {
			// file.template:599:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:600:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:600:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:600:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:601:8
			if len(clause.Bindings) > 0 {
				// file.template:601:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:602:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:602:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:602:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:603:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:603:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:603:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:604:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:605:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"file.template:605:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:605:53
			{
				_n, _err := _output.Write([]byte(`
        symbol.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:606:15
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.ValueType),
					"file.template:606:15")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:606:32
			{
				_n, _err := _output.Write([]byte(`, err = reducer.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:606:48
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:606:48")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:606:76
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:607:8
			for idx, term := range clause.Bindings {
				// file.template:607:52
				{
					_n, _err := _output.Write([]byte(`args[`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:608:5
				{
					_n, _err := _template.writeValue(
						_output,
						(idx),
						"file.template:608:5")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:608:9
				{
					_n, _err := _output.Write([]byte(`].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:608:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.ValueType),
						"file.template:608:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:609:12
				if idx != len(clause.Bindings)-1 {
					// file.template:609:49
					{
						_n, _err := _output.Write([]byte(`,`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:610:17
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:613:8
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
	// file.template:619:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:619:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:619:21
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
	// file.template:625:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:625:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:625:18
	{
		_n, _err := _output.Write([]byte(` struct {
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:626:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:626:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:626:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:627:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:627:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:627:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:630:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:21
	{
		_n, _err := _output.Write([]byte(` map[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:26
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:630:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:39
	{
		_n, _err := _output.Write([]byte(`]*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:41
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:630:41")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:630:52
	{
		_n, _err := _output.Write([]byte(`

func (table `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:632:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:632:28
	{
		_n, _err := _output.Write([]byte(`) Get(
    stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:633:12
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:633:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:633:24
	{
		_n, _err := _output.Write([]byte(`,
    symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:634:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:634:26
	{
		_n, _err := _output.Write([]byte(`) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:635:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:635:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:635:16
	{
		_n, _err := _output.Write([]byte(`,
    bool) {

    action, ok := table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:638:24
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:638:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:638:37
	{
		_n, _err := _output.Write([]byte(`{stateId, symbolId}]
    if ok {
        return action, ok
    }

    action, ok = table[`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:643:23
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:643:23")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:643:36
	{
		_n, _err := _output.Write([]byte(`{stateId, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:643:46
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:643:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:643:63
	{
		_n, _err := _output.Write([]byte(`}]
    return action, ok
}

var (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:648:0
	for _, state := range States.OrderedStates {
		// file.template:648:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenAction),
				"file.template:649:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:26
		{
			_n, _err := _output.Write([]byte(` = &`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:30
		{
			_n, _err := _template.writeValue(
				_output,
				(ActionType),
				"file.template:649:30")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:41
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:42
		{
			_n, _err := _template.writeValue(
				_output,
				(ShiftAction),
				"file.template:649:42")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:54
		{
			_n, _err := _output.Write([]byte(`, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:56
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:649:56")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:649:77
		{
			_n, _err := _output.Write([]byte(`, 0}`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:652:0
	for _, term := range Grammar.NonTerminals {
		// file.template:653:4
		for _, clause := range term.Clauses {
			// file.template:653:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReduceAction),
					"file.template:654:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:33
			{
				_n, _err := _output.Write([]byte(` = &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:37
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"file.template:654:37")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:48
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:49
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceAction),
					"file.template:654:49")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:62
			{
				_n, _err := _output.Write([]byte(`, 0, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:67
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:654:67")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:654:100
			{
				_n, _err := _output.Write([]byte(`}`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:656:8
	{
		_n, _err := _output.Write([]byte(`
)

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:659:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:659:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:659:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:659:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:659:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:659:35
	{
		_n, _err := _output.Write([]byte(`{`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:660:0
	for _, state := range States.OrderedStates {
		// file.template:661:4
		for _, item := range state.Items {
			// file.template:662:8

			if !item.IsReduce {
				continue
			}

			if item.Name != lr.AcceptRule || item.LookAhead != lr.EndMarker {
				continue
			}

			// file.template:672:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:673:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:28
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"file.template:673:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:40
			{
				_n, _err := _output.Write([]byte(`}: &`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:44
			{
				_n, _err := _template.writeValue(
					_output,
					(ActionType),
					"file.template:673:44")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:55
			{
				_n, _err := _output.Write([]byte(`{`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:56
			{
				_n, _err := _template.writeValue(
					_output,
					(AcceptAction),
					"file.template:673:56")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:673:69
			{
				_n, _err := _output.Write([]byte(`, 0, 0},`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:677:0
	for _, state := range States.OrderedStates {
		// file.template:678:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:679:8

			symbol := Grammar.Terms[symbolName]
			child, ok := state.Goto[symbolName]
			if !ok {
				continue
			}

			// file.template:687:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:688:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:28
			{
				_n, _err := _template.writeValue(
					_output,
					(symbol.CodeGenSymbolConst),
					"file.template:688:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:56
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:59
			{
				_n, _err := _template.writeValue(
					_output,
					(child.CodeGenAction),
					"file.template:688:59")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:688:81
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:692:0
	for _, state := range States.OrderedStates {
		// file.template:693:4
		for _, item := range state.Items {
			// file.template:694:8

			if !item.IsReduce {
				continue
			}

			if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
				continue
			}

			idConst := Grammar.Terms[item.LookAhead].CodeGenSymbolConst
			reduceAction := item.Clause.CodeGenReduceAction

			// file.template:707:10
			{
				_n, _err := _output.Write([]byte(`
    {`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:5
			{
				_n, _err := _template.writeValue(
					_output,
					(state.CodeGenConst),
					"file.template:708:5")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:26
			{
				_n, _err := _output.Write([]byte(`, `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:28
			{
				_n, _err := _template.writeValue(
					_output,
					(idConst),
					"file.template:708:28")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:36
			{
				_n, _err := _output.Write([]byte(`}: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:39
			{
				_n, _err := _template.writeValue(
					_output,
					(reduceAction),
					"file.template:708:39")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:708:52
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:710:8
	{
		_n, _err := _output.Write([]byte(`
}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:714:0
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
	// file.template:720:3
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
