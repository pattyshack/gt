// Auto-generated from source: file.template

package debug_template

import (
	_fmt "fmt"
	_io "io"

	lr "github.com/pattyshack/gt/tools/lr/internal"
)

type File struct {
	OrderedSymbolNames        []string
	States                    *lr.LRStates
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

	OrderedSymbolNames := _template.OrderedSymbolNames
	States := _template.States
	OutputDebugNonKernelItems := _template.OutputDebugNonKernelItems

	// file.template:14:0

	gotoCount := 0
	reduceCount := 0
	shiftAndReduceCount := 0
	shiftReduceCount := 0
	reduceReduceCount := 0

	// file.template:22:3
	{
		_n, _err := _output.Write([]byte(`Parser Debug States:`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:24:0
	for _, state := range States.OrderedStates {
		// file.template:24:47
		{
			_n, _err := _output.Write([]byte(`
  State `))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:25:8
		{
			_n, _err := _template.writeValue(
				_output,
				(state.StateNum),
				"file.template:25:8")
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:25:25
		{
			_n, _err := _output.Write([]byte(`:
    Kernel Items:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:27:4
		firstNonKernel := true
		// file.template:28:4
		for _, item := range state.Items {
			// file.template:29:8
			if !item.IsKernel && firstNonKernel {
				// file.template:30:12

				if !OutputDebugNonKernelItems &&
					len(state.ShiftReduceConflictSymbols) == 0 &&
					len(state.ReduceReduceConflictSymbols) == 0 {

					break
				}

				firstNonKernel = false

				// file.template:41:14
				{
					_n, _err := _output.Write([]byte(`
    Non-kernel Items:`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
			// file.template:43:17
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:45:6
			{
				_n, _err := _template.writeValue(
					_output,
					(item),
					"file.template:45:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:46:13
		{
			_n, _err := _output.Write([]byte(`
    Reduce:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:49:4
		if len(state.Reduce) == 0 {
			// file.template:49:34
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:53:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:54:8

			items := state.Reduce[symbolName]
			reduceCount += len(items)

			if len(items) == 0 {
				continue
			}

			// file.template:63:11
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:65:6
			{
				_n, _err := _template.writeValue(
					_output,
					(symbolName),
					"file.template:65:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:65:17
			{
				_n, _err := _output.Write([]byte(` -> [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:66:8
			for idx, item := range items {
				// file.template:67:0
				{
					_n, _err := _template.writeValue(
						_output,
						(item.Name),
						"file.template:67:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:68:12
				if idx != len(items)-1 {
					// file.template:68:41
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:69:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:71:13
		{
			_n, _err := _output.Write([]byte(`
    ShiftAndReduce:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:74:4
		shiftAndReduceCount += len(state.ShiftAndReduce)
		// file.template:75:4
		if len(state.ShiftAndReduce) == 0 {
			// file.template:75:43
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:79:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:80:6

			item, ok := state.ShiftAndReduce[symbolName]
			if !ok {
				continue
			}

			// file.template:86:8
			{
				_n, _err := _output.Write([]byte(`
      `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:87:6
			{
				_n, _err := _template.writeValue(
					_output,
					(symbolName),
					"file.template:87:6")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:87:17
			{
				_n, _err := _output.Write([]byte(` -> [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:87:22
			{
				_n, _err := _template.writeValue(
					_output,
					(item.Name),
					"file.template:87:22")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:87:34
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:88:13
		{
			_n, _err := _output.Write([]byte(`
    Goto:`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
		// file.template:91:4
		gotoCount += len(state.Goto)
		// file.template:92:4
		if len(state.Goto) == 0 {
			// file.template:92:33
			{
				_n, _err := _output.Write([]byte(`
      (nil)`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:96:4
		for _, symbolName := range OrderedSymbolNames {
			// file.template:97:8
			child, ok := state.Goto[symbolName]
			// file.template:98:8
			if ok {
				// file.template:98:18
				{
					_n, _err := _output.Write([]byte(`
      `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:99:6
				{
					_n, _err := _template.writeValue(
						_output,
						(symbolName),
						"file.template:99:6")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:99:17
				{
					_n, _err := _output.Write([]byte(` -> State `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:99:27
				{
					_n, _err := _template.writeValue(
						_output,
						(child.StateNum),
						"file.template:99:27")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}
		}
		// file.template:103:4
		if len(state.ShiftReduceConflictSymbols) > 0 {
			// file.template:104:8
			shiftReduceCount += len(state.ShiftReduceConflictSymbols)
			// file.template:104:73
			{
				_n, _err := _output.Write([]byte(`
    Shift/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:107:8
			for idx, symbol := range state.ShiftReduceConflictSymbols {
				// file.template:108:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"file.template:108:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:109:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// file.template:109:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:110:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:114:4
		if len(state.ReduceReduceConflictSymbols) > 0 {
			// file.template:115:8
			reduceReduceCount += len(state.ReduceReduceConflictSymbols)
			// file.template:115:75
			{
				_n, _err := _output.Write([]byte(`
    Reduce/reduce conflict symbols:
      [`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
			// file.template:118:8
			for idx, symbol := range state.ReduceReduceConflictSymbols {
				// file.template:119:0
				{
					_n, _err := _template.writeValue(
						_output,
						(symbol),
						"file.template:119:0")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
				// file.template:120:12
				if idx != len(state.ShiftReduceConflictSymbols)-1 {
					// file.template:120:68
					{
						_n, _err := _output.Write([]byte(` `))
						_numWritten += int64(_n)
						if _err != nil {
							return _numWritten, _err
						}
					}
				}
			}
			// file.template:121:17
			{
				_n, _err := _output.Write([]byte(`]`))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}
		// file.template:123:13
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}
	// file.template:125:7
	{
		_n, _err := _output.Write([]byte(`
Number of states: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:18
	{
		_n, _err := _template.writeValue(
			_output,
			(len(States.OrderedStates)),
			"file.template:126:18")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:126:46
	{
		_n, _err := _output.Write([]byte(`
Number of shift actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:25
	{
		_n, _err := _template.writeValue(
			_output,
			(gotoCount),
			"file.template:127:25")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:127:35
	{
		_n, _err := _output.Write([]byte(`
Number of reduce actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:26
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceCount),
			"file.template:128:26")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:128:38
	{
		_n, _err := _output.Write([]byte(`
Number of shift-and-reduce actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:36
	{
		_n, _err := _template.writeValue(
			_output,
			(shiftAndReduceCount),
			"file.template:129:36")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:129:56
	{
		_n, _err := _output.Write([]byte(`
Number of shift/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:130:34
	{
		_n, _err := _template.writeValue(
			_output,
			(shiftReduceCount),
			"file.template:130:34")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:130:51
	{
		_n, _err := _output.Write([]byte(`
Number of reduce/reduce conflicts: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:131:35
	{
		_n, _err := _template.writeValue(
			_output,
			(reduceReduceCount),
			"file.template:131:35")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:131:53
	{
		_n, _err := _output.Write([]byte(`
Number of unoptimized states: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:132:30
	{
		_n, _err := _template.writeValue(
			_output,
			(States.NumPreMergedStates),
			"file.template:132:30")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:132:58
	{
		_n, _err := _output.Write([]byte(`
Number of unoptimized shift actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:37
	{
		_n, _err := _template.writeValue(
			_output,
			(States.NumPreMergedShift),
			"file.template:133:37")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:133:64
	{
		_n, _err := _output.Write([]byte(`
Number of unoptimized reduce actions: `))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:134:38
	{
		_n, _err := _template.writeValue(
			_output,
			(States.NumPreMergedReduce),
			"file.template:134:38")
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}
	// file.template:134:66
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
