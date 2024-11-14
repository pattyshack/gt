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
	GenericSymbol             interface{}
	StackItemType             string
	StackType                 string
	SymbolStackType           string
	SymbolIdType              string
	EndSymbolId               string
	WildcardSymbolId          string
	Location                  interface{}
	StartEndPos               interface{}
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
	GenericSymbol := _template.GenericSymbol
	StackItemType := _template.StackItemType
	StackType := _template.StackType
	SymbolStackType := _template.SymbolStackType
	SymbolIdType := _template.SymbolIdType
	EndSymbolId := _template.EndSymbolId
	WildcardSymbolId := _template.WildcardSymbolId
	Location := _template.Location
	StartEndPos := _template.StartEndPos
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
			(TokenType),
			"file.template:125:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:125:15
	{
		_n, _err := _output.Write([]byte(` interface {
    Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:126:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:22
	{
		_n, _err := _output.Write([]byte(`
    Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:10
	{
		_n, _err := _template.writeValue(
			_output,
			(Location),
			"file.template:127:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:0
	if GenerateEndPos {
		// file.template:128:22
		{
			_n, _err := _output.Write([]byte(`
    End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:129:10
		{
			_n, _err := _template.writeValue(
				_output,
				(Location),
				"file.template:129:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:130:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:5
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:133:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:15
	{
		_n, _err := _output.Write([]byte(` interface {
    // Note: Return io.EOF to indicate end of stream
    // Token with unspecified value type should return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:55
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbol),
			"file.template:135:55")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:135:69
	{
		_n, _err := _output.Write([]byte(`
    Next() (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:12
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:136:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:136:22
	{
		_n, _err := _output.Write([]byte(`, error)

    CurrentLocation() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:138:22
	{
		_n, _err := _template.writeValue(
			_output,
			(Location),
			"file.template:138:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:138:31
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:141:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:142:4
		if rule.NumReducerClauses == 0 {
			// file.template:143:8
			continue
		}
		// file.template:144:12
		{
			_n, _err := _output.Write([]byte(`
type `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:145:5
		{
			_n, _err := _template.writeValue(
				_output,
				(rule.CodeGenReducerInterface),
				"file.template:145:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:145:36
		{
			_n, _err := _output.Write([]byte(` interface {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:146:4
		for clauseIdx, clause := range rule.Clauses {
			// file.template:147:8
			if clause.Passthrough {
				// file.template:148:10
				continue
			}
			// file.template:150:8
			if clauseIdx > 0 {
				// file.template:150:29
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:154:8
			if clause.Label == "" {
				// file.template:154:34
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:155:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Location.ShortString()),
						"file.template:155:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:155:39
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:155:41
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:155:41")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:155:53
				{
					_n, _err := _output.Write([]byte(` -> ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:156:17
				{
					_n, _err := _output.Write([]byte(`
    // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:7
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Location.ShortString()),
						"file.template:157:7")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:39
				{
					_n, _err := _output.Write([]byte(`: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:41
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.Name),
						"file.template:157:41")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:53
				{
					_n, _err := _output.Write([]byte(` -> `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:57
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Label),
						"file.template:157:57")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:157:72
				{
					_n, _err := _output.Write([]byte(`: ...`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:160:8
			paramNameCount := map[string]int{}
			// file.template:160:49
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:161:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:161:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:161:32
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:162:8
			for termIdx, term := range clause.Bindings {
				// file.template:164:12

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

				// file.template:187:0
				{
					_n, _err := _template.writeValue(
						_output,
						(paramName),
						"file.template:187:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:187:10
				{
					_n, _err := _output.Write([]byte(` `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:187:11
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenType),
						"file.template:187:11")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:187:30
				{
					_n, _err := _template.writeValue(
						_output,
						(suffix),
						"file.template:187:30")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:188:17
			{
				_n, _err := _output.Write([]byte(`) (`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:189:3
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenType),
					"file.template:189:3")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:189:22
			{
				_n, _err := _output.Write([]byte(`, error)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:190:13
		{
			_n, _err := _output.Write([]byte(`}
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:192:7
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:194:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:194:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:194:17
	{
		_n, _err := _output.Write([]byte(` interface {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:195:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:196:4
		if len(rule.Clauses) == 0 {
			// file.template:197:8
			continue
		}
		// file.template:199:2
		if rule.NumReducerClauses > 0 {
			// file.template:199:36
			{
				_n, _err := _output.Write([]byte(`
  `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:200:2
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenReducerInterface),
					"file.template:200:2")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:202:8
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:205:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:205:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:205:20
	{
		_n, _err := _output.Write([]byte(` interface {
    Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:206:20
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:206:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:206:30
	{
		_n, _err := _output.Write([]byte(`, parseStack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:206:43
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:206:43")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:206:53
	{
		_n, _err := _output.Write([]byte(`) error
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:209:5
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:209:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:209:27
	{
		_n, _err := _output.Write([]byte(` struct {}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:6
	{
		_n, _err := _template.writeValue(
			_output,
			(DefaultErrHandlerType),
			"file.template:211:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:28
	{
		_n, _err := _output.Write([]byte(`) Error(nextToken `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:46
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:211:46")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:56
	{
		_n, _err := _output.Write([]byte(`, stack `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:64
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:211:64")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:211:74
	{
		_n, _err := _output.Write([]byte(`) error {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:212:11
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:212:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:212:28
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
	// file.template:216:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:216:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:216:30
	{
		_n, _err := _output.Write([]byte(`(stack[len(stack)-1].StateId))
}

func `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ExpectedTerminalsFunc),
			"file.template:219:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:27
	{
		_n, _err := _output.Write([]byte(`(id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:31
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:219:31")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:43
	{
		_n, _err := _output.Write([]byte(`) []`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:47
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:219:47")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:219:60
	{
		_n, _err := _output.Write([]byte(` {
  switch id {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:221:0
	for _, state := range States.OrderedStates {
		// file.template:222:2

		_, ok := state.Reduce[lr.Wildcard]
		if ok {
			continue
		}

		// file.template:227:4
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:228:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:228:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:228:28
		{
			_n, _err := _output.Write([]byte(`:
    return []`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:229:13
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:229:13")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:229:26
		{
			_n, _err := _output.Write([]byte(`{`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:230:4
		for _, term := range Grammar.Terminals {
			// file.template:231:6

			_, foundGoto := state.Goto[term.Name]
			_, foundReduce := state.Reduce[term.Name]
			_, foundShiftAndReduce := state.ShiftAndReduce[term.Name]

			if !foundGoto && !foundReduce && !foundShiftAndReduce {
				continue
			}

			// file.template:240:9
			{
				_n, _err := _output.Write([]byte(`      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:241:6
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:241:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:241:32
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:243:4
		if _, ok := state.Reduce[lr.EndMarker]; ok {
			// file.template:243:52
			{
				_n, _err := _output.Write([]byte(`      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:244:6
			{
				_n, _err := _template.writeValue(
					_output,
					(EndSymbolId),
					"file.template:244:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:244:18
			{
				_n, _err := _output.Write([]byte(`,`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:245:13
		{
			_n, _err := _output.Write([]byte(`    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:247:8
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
	// file.template:253:0
	for idx, start := range Grammar.Starts {
		// file.template:254:4

		parseSuffix := ""
		if len(Grammar.Starts) > 1 {
			parseSuffix = codegen.SnakeToCamel(start.Name)
		}

		// file.template:261:6
		{
			_n, _err := _output.Write([]byte(`
func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:262:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:262:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:37
		{
			_n, _err := _output.Write([]byte(`(lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:44
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:262:44")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:54
		{
			_n, _err := _output.Write([]byte(`, reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:64
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:262:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:262:76
		{
			_n, _err := _output.Write([]byte(`) (`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:264:0
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:264:0")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:264:20
		{
			_n, _err := _output.Write([]byte(`, error) {

    return `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:266:11
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:266:11")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:266:29
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:266:29")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:266:43
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
		// file.template:269:8
		{
			_n, _err := _template.writeValue(
				_output,
				(DefaultErrHandlerType),
				"file.template:269:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:269:30
		{
			_n, _err := _output.Write([]byte(`{})
}

func `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:272:5
		{
			_n, _err := _template.writeValue(
				_output,
				(ParseFuncPrefix),
				"file.template:272:5")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:272:23
		{
			_n, _err := _template.writeValue(
				_output,
				(parseSuffix),
				"file.template:272:23")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:272:37
		{
			_n, _err := _output.Write([]byte(`WithCustomErrorHandler(
    lexer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:273:10
		{
			_n, _err := _template.writeValue(
				_output,
				(LexerType),
				"file.template:273:10")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:273:20
		{
			_n, _err := _output.Write([]byte(`,
    reducer `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:274:12
		{
			_n, _err := _template.writeValue(
				_output,
				(ReducerType),
				"file.template:274:12")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:274:24
		{
			_n, _err := _output.Write([]byte(`,
    errHandler `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:275:15
		{
			_n, _err := _template.writeValue(
				_output,
				(ErrHandlerType),
				"file.template:275:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:275:30
		{
			_n, _err := _output.Write([]byte(`,
) (
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:277:4
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:277:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:277:24
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
		// file.template:280:17
		{
			_n, _err := _template.writeValue(
				_output,
				(InternalParseFunc),
				"file.template:280:17")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:35
		{
			_n, _err := _output.Write([]byte(`(lexer, reducer, errHandler, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:280:64
		{
			_n, _err := _template.writeValue(
				_output,
				(States.OrderedStates[idx].CodeGenConst),
				"file.template:280:64")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:281:47
		{
			_n, _err := _output.Write([]byte(`)
    if err != nil {
        var errRetVal `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:283:22
		{
			_n, _err := _template.writeValue(
				_output,
				(start.CodeGenType),
				"file.template:283:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:283:42
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
		// file.template:286:16
		{
			_n, _err := _template.writeValue(
				_output,
				(start.ValueType),
				"file.template:286:16")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:286:34
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
	// file.template:288:7
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
	// file.template:295:5
	{
		_n, _err := _template.writeValue(
			_output,
			(InternalParseFunc),
			"file.template:295:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:295:23
	{
		_n, _err := _output.Write([]byte(`(
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:296:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:296:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:296:20
	{
		_n, _err := _output.Write([]byte(`,
    reducer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:297:12
	{
		_n, _err := _template.writeValue(
			_output,
			(ReducerType),
			"file.template:297:12")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:297:24
	{
		_n, _err := _output.Write([]byte(`,
    errHandler `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:298:15
	{
		_n, _err := _template.writeValue(
			_output,
			(ErrHandlerType),
			"file.template:298:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:298:30
	{
		_n, _err := _output.Write([]byte(`,
    startState `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:299:15
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:299:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:299:27
	{
		_n, _err := _output.Write([]byte(`,
) (
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:301:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:301:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:301:19
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
	// file.template:304:18
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:304:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:304:28
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
	// file.template:307:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StackItemType),
			"file.template:307:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:307:23
	{
		_n, _err := _output.Write([]byte(`{startState, nil},
    }

    symbolStack := &`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:20
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:310:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:310:36
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
	// file.template:318:22
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:318:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:318:34
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
	// file.template:325:32
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:325:32")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:325:44
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
	// file.template:332:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:332:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:332:52
	{
		_n, _err := _output.Write([]byte(` {
            var reduceSymbol *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:333:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:333:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:333:41
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
	// file.template:342:39
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:342:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:342:60
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
	// file.template:350:30
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:350:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:350:41
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
	// file.template:359:39
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:359:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:359:52
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
	// file.template:370:8
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:370:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:370:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:9
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:372:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:372:21
	{
		_n, _err := _output.Write([]byte(`:
        return "$"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:374:9
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:374:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:374:26
	{
		_n, _err := _output.Write([]byte(`:
        return "*"`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:376:0
	for _, symbolName := range OrderedSymbolNames[3:] {
		// file.template:377:4
		term := Grammar.Terms[symbolName]
		// file.template:377:44
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:378:9
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:378:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:378:35
		{
			_n, _err := _output.Write([]byte(`:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:379:4
		if term.SymbolId == parser.LRCharacterToken {
			// file.template:380:8

			escaped := term.Name
			if term.Name == "'\"'" {
				escaped = "'\\\"'"
			} else if escaped[1] == '\\' {
				escaped = "'\\\\" + term.Name[2:]
			}

			// file.template:389:10
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:390:16
			{
				_n, _err := _template.writeValue(
					_output,
					(escaped),
					"file.template:390:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:390:24
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		} else {
			// file.template:391:13
			{
				_n, _err := _output.Write([]byte(`
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:392:16
			{
				_n, _err := _template.writeValue(
					_output,
					(term.Name),
					"file.template:392:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:392:28
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:394:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:396:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:396:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:396:23
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
	// file.template:401:4
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:401:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:401:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:401:19
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:401:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:401:32
	{
		_n, _err := _output.Write([]byte(`(0)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:402:4
	{
		_n, _err := _template.writeValue(
			_output,
			(WildcardSymbolId),
			"file.template:402:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:402:21
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:402:24
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:402:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:402:37
	{
		_n, _err := _output.Write([]byte(`(-1)
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:404:0
	for idx, term := range Grammar.NonTerminals {
		// file.template:404:48
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:4
		{
			_n, _err := _template.writeValue(
				_output,
				(term.CodeGenSymbolConst),
				"file.template:405:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:30
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:33
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolIdType),
				"file.template:405:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:46
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:47
		{
			_n, _err := _template.writeValue(
				_output,
				(256 + len(Grammar.Terminals) + idx),
				"file.template:405:47")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:405:80
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:406:8
	{
		_n, _err := _output.Write([]byte(`
)

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:409:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:409:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:409:18
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
	// file.template:413:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
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
			(ActionIdType),
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
			(ReduceAction),
			"file.template:414:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:414:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:414:33
	{
		_n, _err := _output.Write([]byte(`(1)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:415:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:415:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:415:25
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:415:28
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:415:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:415:41
	{
		_n, _err := _output.Write([]byte(`(2)
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:416:4
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:416:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:416:17
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:416:20
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:416:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:416:33
	{
		_n, _err := _output.Write([]byte(`(3)
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionIdType),
			"file.template:419:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:419:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:421:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAction),
			"file.template:421:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:421:21
	{
		_n, _err := _output.Write([]byte(`:
        return "shift"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:423:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceAction),
			"file.template:423:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:423:22
	{
		_n, _err := _output.Write([]byte(`:
        return "reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:425:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ShiftAndReduceAction),
			"file.template:425:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:425:30
	{
		_n, _err := _output.Write([]byte(`:
        return "shift-and-reduce"
    case `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:9
	{
		_n, _err := _template.writeValue(
			_output,
			(AcceptAction),
			"file.template:427:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:427:22
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
	// file.template:430:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:430:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:430:23
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
	// file.template:434:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:434:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:434:16
	{
		_n, _err := _output.Write([]byte(` int

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:437:0
	clauseIdx := 1
	// file.template:438:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:439:4
		for _, clause := range rule.Clauses {
			// file.template:439:44
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:4
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:440:4")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:37
			{
				_n, _err := _output.Write([]byte(` = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:40
			{
				_n, _err := _template.writeValue(
					_output,
					(ReduceType),
					"file.template:440:40")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:51
			{
				_n, _err := _output.Write([]byte(`(`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:52
			{
				_n, _err := _template.writeValue(
					_output,
					(clauseIdx),
					"file.template:440:52")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:440:62
			{
				_n, _err := _output.Write([]byte(`)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:441:8
			clauseIdx += 1
		}
	}
	// file.template:443:8
	{
		_n, _err := _output.Write([]byte(`
)

func (i `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:446:8
	{
		_n, _err := _template.writeValue(
			_output,
			(ReduceType),
			"file.template:446:8")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:446:19
	{
		_n, _err := _output.Write([]byte(`) String() string {
    switch i {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:448:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:449:4
		for _, clause := range rule.Clauses {
			// file.template:449:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:450:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:450:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:450:42
			{
				_n, _err := _output.Write([]byte(`:
        return "`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:451:16
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerName),
					"file.template:451:16")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:451:44
			{
				_n, _err := _output.Write([]byte(`"`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
	}
	// file.template:453:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:455:15
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:455:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:455:23
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
	// file.template:459:5
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:459:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:459:17
	{
		_n, _err := _output.Write([]byte(` int

func (id `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:461:9
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:461:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:461:21
	{
		_n, _err := _output.Write([]byte(`) String() string {
    return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:462:11
	{
		_n, _err := _template.writeValue(
			_output,
			(Sprintf),
			"file.template:462:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:462:19
	{
		_n, _err := _output.Write([]byte(`("State %d", int(id))
}

const (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:466:0
	for _, state := range States.OrderedStates {
		// file.template:466:47
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:4
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:467:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:25
		{
			_n, _err := _output.Write([]byte(` = `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:28
		{
			_n, _err := _template.writeValue(
				_output,
				(StateIdType),
				"file.template:467:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:40
		{
			_n, _err := _output.Write([]byte(`(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:41
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:467:41")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:467:58
		{
			_n, _err := _output.Write([]byte(`)`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:468:8
	{
		_n, _err := _output.Write([]byte(`
)

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
			(SymbolType),
			"file.template:471:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:471:16
	{
		_n, _err := _output.Write([]byte(` struct {
    SymbolId_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:472:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:472:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:472:27
	{
		_n, _err := _output.Write([]byte(`

    Generic_ `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:474:13
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbol),
			"file.template:474:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:474:27
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:476:0
	for _, valueType := range OrderedValueTypes {
		// file.template:477:4
		if valueType.Name == lr.Generic {
			// file.template:478:8
			continue
		}
		// file.template:479:12
		{
			_n, _err := _output.Write([]byte(`
    `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:480:4
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:480:4")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:480:21
		{
			_n, _err := _output.Write([]byte(` `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:480:22
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:480:22")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:481:8
	{
		_n, _err := _output.Write([]byte(`
}
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:484:0

	valueTerms := map[string][]*lr.Term{}
	for _, symbolName := range OrderedSymbolNames[2:] {
		term := Grammar.Terms[symbolName]
		valueTerms[term.ValueType] = append(valueTerms[term.ValueType], term)
	}

	// file.template:492:3
	{
		_n, _err := _output.Write([]byte(`func NewSymbol(token `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:493:21
	{
		_n, _err := _template.writeValue(
			_output,
			(TokenType),
			"file.template:493:21")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:493:31
	{
		_n, _err := _output.Write([]byte(`) (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:493:35
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:493:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:493:46
	{
		_n, _err := _output.Write([]byte(`, error) {
    symbol, ok := token.(*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:494:26
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:494:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:494:37
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
	// file.template:499:14
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:499:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:499:25
	{
		_n, _err := _output.Write([]byte(`{SymbolId_: token.Id()}
    switch token.Id() {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:501:0
	for _, valueType := range OrderedValueTypes {
		// file.template:502:4

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

		// file.template:517:6
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:519:4
		for idx, kconst := range consts {
			// file.template:520:0
			{
				_n, _err := _template.writeValue(
					_output,
					(kconst),
					"file.template:520:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:520:8
			if idx != len(consts)-1 {
				// file.template:520:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:521:13
		{
			_n, _err := _output.Write([]byte(`:
        val, ok := token.(`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:523:26
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:523:26")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:523:48
		{
			_n, _err := _output.Write([]byte(`)
        if !ok {
            return nil, `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:525:24
		{
			_n, _err := _template.writeValue(
				_output,
				(NewLocationError),
				"file.template:525:24")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:525:41
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
		// file.template:528:31
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.ParamType),
				"file.template:528:31")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:528:53
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
		// file.template:531:15
		{
			_n, _err := _template.writeValue(
				_output,
				(valueType.Name),
				"file.template:531:15")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:531:32
		{
			_n, _err := _output.Write([]byte(` = val`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:532:8
	{
		_n, _err := _output.Write([]byte(`
    default:
        return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:534:20
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:534:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:534:37
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
	// file.template:542:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:542:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:542:20
	{
		_n, _err := _output.Write([]byte(`) Id() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:542:27
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:542:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:542:40
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
	// file.template:546:9
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:546:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:546:20
	{
		_n, _err := _output.Write([]byte(`) Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:546:28
	{
		_n, _err := _template.writeValue(
			_output,
			(Location),
			"file.template:546:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:546:37
	{
		_n, _err := _output.Write([]byte(` {
    type locator interface { Loc() `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:547:35
	{
		_n, _err := _template.writeValue(
			_output,
			(Location),
			"file.template:547:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:547:44
	{
		_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:549:0
	for _, field := range OrderedValueTypes {
		// file.template:550:4
		if field.Name == lr.Generic {
			// file.template:551:8
			continue
		}
		// file.template:553:4
		terms := valueTerms[field.Name]
		// file.template:554:4
		if len(terms) == 0 {
			// file.template:555:6
			continue
		}
		// file.template:556:12
		{
			_n, _err := _output.Write([]byte(`
    case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:558:4
		for idx, term := range terms {
			// file.template:559:0
			{
				_n, _err := _template.writeValue(
					_output,
					(term.CodeGenSymbolConst),
					"file.template:559:0")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:560:8
			if idx != len(terms)-1 {
				// file.template:560:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:561:13
		{
			_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:563:33
		{
			_n, _err := _template.writeValue(
				_output,
				(field.Name),
				"file.template:563:33")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:563:46
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
	// file.template:567:8
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
	// file.template:572:0
	if GenerateEndPos {
		// file.template:572:22
		{
			_n, _err := _output.Write([]byte(`
func (s *`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:573:9
		{
			_n, _err := _template.writeValue(
				_output,
				(SymbolType),
				"file.template:573:9")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:573:20
		{
			_n, _err := _output.Write([]byte(`) End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:573:28
		{
			_n, _err := _template.writeValue(
				_output,
				(Location),
				"file.template:573:28")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:573:37
		{
			_n, _err := _output.Write([]byte(` {
    type locator interface { End() `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:574:35
		{
			_n, _err := _template.writeValue(
				_output,
				(Location),
				"file.template:574:35")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:574:44
		{
			_n, _err := _output.Write([]byte(` }
    switch s.SymbolId_ {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:576:0
		for _, field := range OrderedValueTypes {
			// file.template:577:4
			if field.Name == lr.Generic {
				// file.template:578:8
				continue
			}
			// file.template:580:4
			terms := valueTerms[field.Name]
			// file.template:581:4
			if len(terms) == 0 {
				// file.template:582:6
				continue
			}
			// file.template:583:12
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:585:4
			for idx, term := range terms {
				// file.template:586:0
				{
					_n, _err := _template.writeValue(
						_output,
						(term.CodeGenSymbolConst),
						"file.template:586:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:587:8
				if idx != len(terms)-1 {
					// file.template:587:37
					{
						_n, _err := _output.Write([]byte(`, `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:588:13
			{
				_n, _err := _output.Write([]byte(`:
        loc, ok := interface{}(s.`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:590:33
			{
				_n, _err := _template.writeValue(
					_output,
					(field.Name),
					"file.template:590:33")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:590:46
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
		// file.template:594:8
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
	// file.template:598:8
	{
		_n, _err := _output.Write([]byte(`

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:600:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:600:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:600:21
	{
		_n, _err := _output.Write([]byte(` struct {
    lexer `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:601:10
	{
		_n, _err := _template.writeValue(
			_output,
			(LexerType),
			"file.template:601:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:601:20
	{
		_n, _err := _output.Write([]byte(`
    top []*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:602:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:602:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:602:22
	{
		_n, _err := _output.Write([]byte(`
}

func (stack *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:13
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolStackType),
			"file.template:605:13")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:29
	{
		_n, _err := _output.Write([]byte(`) Top() (*`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:39
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:605:39")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:605:50
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
	// file.template:609:22
	{
		_n, _err := _template.writeValue(
			_output,
			(EOF),
			"file.template:609:22")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:609:26
	{
		_n, _err := _output.Write([]byte(` {
                return nil, `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:610:28
	{
		_n, _err := _template.writeValue(
			_output,
			(NewLocationError),
			"file.template:610:28")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:610:45
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
	// file.template:615:20
	{
		_n, _err := _template.writeValue(
			_output,
			(GenericSymbol),
			"file.template:615:20")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:615:34
	{
		_n, _err := _output.Write([]byte(`{
              SymbolId: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:24
	{
		_n, _err := _template.writeValue(
			_output,
			(EndSymbolId),
			"file.template:616:24")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:616:36
	{
		_n, _err := _output.Write([]byte(`,
              StartEndPos: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:617:27
	{
		_n, _err := _template.writeValue(
			_output,
			(StartEndPos),
			"file.template:617:27")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:617:39
	{
		_n, _err := _output.Write([]byte(`{
                StartPos: stack.lexer.CurrentLocation(),
                EndPos: stack.lexer.CurrentLocation(),
              },
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
		_n, _err := _output.Write([]byte(`,
) (
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:4
	{
		_n, _err := _template.writeValue(
			_output,
			(StackType),
			"file.template:668:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:668:14
	{
		_n, _err := _output.Write([]byte(`,
    *`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:669:5
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:669:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:669:16
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
	// file.template:673:15
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolType),
			"file.template:673:15")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:673:26
	{
		_n, _err := _output.Write([]byte(`{}
    switch act.ReduceType {`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:675:0
	for _, rule := range Grammar.NonTerminals {
		// file.template:676:4
		for _, clause := range rule.Clauses {
			// file.template:676:44
			{
				_n, _err := _output.Write([]byte(`
    case `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:677:9
			{
				_n, _err := _template.writeValue(
					_output,
					(clause.CodeGenReducerNameConst),
					"file.template:677:9")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:677:42
			{
				_n, _err := _output.Write([]byte(`:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:678:8
			if len(clause.Bindings) > 0 {
				// file.template:678:40
				{
					_n, _err := _output.Write([]byte(`
        args := stack[len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:679:33
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:679:33")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:679:56
				{
					_n, _err := _output.Write([]byte(`:]
        stack = stack[:len(stack)-`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:680:34
				{
					_n, _err := _template.writeValue(
						_output,
						(len(clause.Bindings)),
						"file.template:680:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:680:57
				{
					_n, _err := _output.Write([]byte(`]`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:681:16
			{
				_n, _err := _output.Write([]byte(`
        symbol.SymbolId_ = `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:682:27
			{
				_n, _err := _template.writeValue(
					_output,
					(rule.CodeGenSymbolConst),
					"file.template:682:27")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:683:8
			if clause.Passthrough {
				// file.template:683:34
				{
					_n, _err := _output.Write([]byte(`
        //line `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:15
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Location),
						"file.template:684:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:684:33
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:685:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:685:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:685:32
				{
					_n, _err := _output.Write([]byte(` = args[0].`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:685:43
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.Bindings[0].ValueType),
						"file.template:685:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:685:74
				{
					_n, _err := _output.Write([]byte(`
        err = nil`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			} else {
				// file.template:687:17
				{
					_n, _err := _output.Write([]byte(`
        symbol.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:688:15
				{
					_n, _err := _template.writeValue(
						_output,
						(rule.ValueType),
						"file.template:688:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:688:32
				{
					_n, _err := _output.Write([]byte(`, err = reducer.`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:688:48
				{
					_n, _err := _template.writeValue(
						_output,
						(clause.CodeGenReducerName),
						"file.template:688:48")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:688:76
				{
					_n, _err := _output.Write([]byte(`(`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:689:10
				for idx, term := range clause.Bindings {
					// file.template:689:54
					{
						_n, _err := _output.Write([]byte(`  args[`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:690:7
					{
						_n, _err := _template.writeValue(
							_output,
							(idx),
							"file.template:690:7")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:690:11
					{
						_n, _err := _output.Write([]byte(`].`))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:690:13
					{
						_n, _err := _template.writeValue(
							_output,
							(term.ValueType),
							"file.template:690:13")
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
					// file.template:691:14
					if idx != len(clause.Bindings)-1 {
						// file.template:691:51
						{
							_n, _err := _output.Write([]byte(`,`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:692:19
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
	// file.template:696:8
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
	// file.template:702:14
	{
		_n, _err := _template.writeValue(
			_output,
			(Errorf),
			"file.template:702:14")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:702:21
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
	// file.template:708:5
	{
		_n, _err := _template.writeValue(
			_output,
			(TableKeyType),
			"file.template:708:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:708:18
	{
		_n, _err := _output.Write([]byte(` struct {
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
			(StateIdType),
			"file.template:709:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:709:16
	{
		_n, _err := _output.Write([]byte(`
    `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:710:4
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:710:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:710:17
	{
		_n, _err := _output.Write([]byte(`
}

type `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:713:5
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:713:5")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:713:21
	{
		_n, _err := _output.Write([]byte(` struct{}

func (`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:715:6
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:715:6")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:715:22
	{
		_n, _err := _output.Write([]byte(`) Get(
  stateId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:716:10
	{
		_n, _err := _template.writeValue(
			_output,
			(StateIdType),
			"file.template:716:10")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:716:22
	{
		_n, _err := _output.Write([]byte(`,
  symbolId `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:717:11
	{
		_n, _err := _template.writeValue(
			_output,
			(SymbolIdType),
			"file.template:717:11")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:717:24
	{
		_n, _err := _output.Write([]byte(`,
) (
  `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:719:2
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:719:2")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:719:13
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
	// file.template:723:0
	for _, state := range States.OrderedStates {
		// file.template:723:47
		{
			_n, _err := _output.Write([]byte(`
  case `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:724:7
		{
			_n, _err := _template.writeValue(
				_output,
				(state.CodeGenConst),
				"file.template:724:7")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:724:28
		{
			_n, _err := _output.Write([]byte(`:
    switch symbolId {`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:726:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:727:4
			if symbolName == lr.Wildcard {
				// file.template:728:6
				continue
			}
			// file.template:731:4

			symbol := Grammar.Terms[symbolName]
			nextState, ok := state.Goto[symbolName]

			// file.template:735:4
			if ok {
				// file.template:735:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:736:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:736:37
				{
					_n, _err := _output.Write([]byte(`:
      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:737:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ShiftAction),
						"file.template:737:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:37
				{
					_n, _err := _output.Write([]byte(`, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:39
				{
					_n, _err := _template.writeValue(
						_output,
						(nextState.CodeGenConst),
						"file.template:737:39")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:737:64
				{
					_n, _err := _output.Write([]byte(`, 0}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:738:6
				continue
			}
		}
		// file.template:742:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:743:4
			if symbolName == lr.Wildcard {
				// file.template:744:6
				continue
			}
			// file.template:747:4

			symbol := Grammar.Terms[symbolName]
			reduceItem, ok := state.ShiftAndReduce[symbolName]

			// file.template:751:4
			if ok {
				// file.template:751:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:752:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:752:37
				{
					_n, _err := _output.Write([]byte(`:
      return `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:13
				{
					_n, _err := _template.writeValue(
						_output,
						(ActionType),
						"file.template:753:13")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:24
				{
					_n, _err := _output.Write([]byte(`{`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:25
				{
					_n, _err := _template.writeValue(
						_output,
						(ShiftAndReduceAction),
						"file.template:753:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:46
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:51
				{
					_n, _err := _template.writeValue(
						_output,
						(reduceItem.Clause.CodeGenReducerNameConst),
						"file.template:753:51")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:753:95
				{
					_n, _err := _output.Write([]byte(`}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:754:6
				continue
			}
		}
		// file.template:758:2
		for _, symbolName := range OrderedSymbolNames {
			// file.template:759:4
			if symbolName == lr.Wildcard {
				// file.template:760:6
				continue
			}
			// file.template:763:4

			symbol := Grammar.Terms[symbolName]
			reduceItems, ok := state.Reduce[symbolName]

			// file.template:767:4
			if ok {
				// file.template:767:14
				{
					_n, _err := _output.Write([]byte(`
    case `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:9
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol.CodeGenSymbolConst),
						"file.template:768:9")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:768:37
				{
					_n, _err := _output.Write([]byte(`:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:769:6
				for _, item := range reduceItems {
					// file.template:770:8
					if item.Name == lr.AcceptRule && item.LookAhead == lr.EndMarker {
						// file.template:770:76
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:771:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:771:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:771:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:771:25
						{
							_n, _err := _template.writeValue(
								_output,
								(AcceptAction),
								"file.template:771:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:771:38
						{
							_n, _err := _output.Write([]byte(`, 0, 0}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					} else {
						// file.template:772:17
						{
							_n, _err := _output.Write([]byte(`
      return `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:13
						{
							_n, _err := _template.writeValue(
								_output,
								(ActionType),
								"file.template:773:13")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:24
						{
							_n, _err := _output.Write([]byte(`{`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:25
						{
							_n, _err := _template.writeValue(
								_output,
								(ReduceAction),
								"file.template:773:25")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:38
						{
							_n, _err := _output.Write([]byte(`, 0, `))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:43
						{
							_n, _err := _template.writeValue(
								_output,
								(item.Clause.CodeGenReducerNameConst),
								"file.template:773:43")
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
						// file.template:773:81
						{
							_n, _err := _output.Write([]byte(`}, true`))
							_numWritten += int64(_n)
							if _err != nil {
								return _numWritten, _err
							}
						}
					}
				}
				// file.template:776:6
				continue
			}
		}
		// file.template:778:10
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:780:4

		reduceItems, ok := state.Reduce[lr.Wildcard]

		// file.template:783:4
		if ok {
			// file.template:783:14
			{
				_n, _err := _output.Write([]byte(`
    default:`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:785:6
			for _, item := range reduceItems {
				// file.template:785:44
				{
					_n, _err := _output.Write([]byte(`      return `))
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
						(ReduceAction),
						"file.template:786:25")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:786:38
				{
					_n, _err := _output.Write([]byte(`, 0, `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:786:43
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Clause.CodeGenReducerNameConst),
						"file.template:786:43")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:786:81
				{
					_n, _err := _output.Write([]byte(`}, true`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:788:13
		{
			_n, _err := _output.Write([]byte(`    }`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:790:8
	{
		_n, _err := _output.Write([]byte(`
  }

  return `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:793:9
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionType),
			"file.template:793:9")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:793:20
	{
		_n, _err := _output.Write([]byte(`{}, false
}

var `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:796:4
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTable),
			"file.template:796:4")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:796:16
	{
		_n, _err := _output.Write([]byte(` = `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:796:19
	{
		_n, _err := _template.writeValue(
			_output,
			(ActionTableType),
			"file.template:796:19")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:796:35
	{
		_n, _err := _output.Write([]byte(`{}

/*
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:799:0
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
	// file.template:805:3
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
