package parseutil

import (
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
)

func IsWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

type Symbol struct {
	Value string // can't be empty string
	Id    int
}

type Symbols []Symbol

func (s Symbols) Len() int { return len(s) }

func (s Symbols) Swap(i int, j int) { s[i], s[j] = s[j], s[i] }

// longer symbol before shorter symbol
// symbol value as tie break
func (s Symbols) Less(i int, j int) bool {
	if len(s[i].Value) == len(s[j].Value) {
		return s[i].Value < s[j].Value
	}

	return len(s[i].Value) > len(s[j].Value)
}

func MaybeTokenizeSymbol(
	reader lexutil.BufferedByteLocationReader,
	sortedSymbols Symbols,
) (
	*Symbol,
	lexutil.Location,
	error,
) {

	var bytes []byte
	var err error
	for idx, symbol := range sortedSymbols {
		if idx == 0 {
			bytes, err = reader.Peek(len(symbol.Value))
			if err != nil {
				if err != io.EOF || len(bytes) == 0 {
					return nil, lexutil.Location{}, err
				}
			}
		}

		if len(bytes) < len(symbol.Value) {
			continue
		}

		if string(bytes[:len(symbol.Value)]) == symbol.Value {
			loc := reader.Location

			bytes = bytes[:len(symbol.Value)]
			n, err := reader.Read(bytes)
			if len(bytes) != n || err != nil {
				panic(err) // should never happen
			}

			return &symbol, loc, nil
		}
	}

	return nil, lexutil.Location{}, nil
}

// If the reader's leading bytes are ascii character of the form 'x' (or
// '\t' for escaped characters), pop those bytes off the reader and return
// the value.  Otherwise, return a nil slice.
func MaybeTokenizeCharacter(
	reader lexutil.BufferedByteLocationReader,
) (string, lexutil.Location, error) {
	bytes, err := reader.Peek(6)
	if len(bytes) < 3 {
		return "", lexutil.Location{}, nil
	}

	if bytes[0] != '\'' {
		return "", lexutil.Location{}, nil
	}

	numBytes := 4
	if bytes[1] == '\\' {
		if len(bytes) < 4 || bytes[3] != '\'' {
			return "", lexutil.Location{}, nil
		}

		switch bytes[2] {
		case 't', 'n', '\'', '\\':
		default:
			return "", lexutil.Location{}, nil
		}
	} else {
		r, size := utf8.DecodeRune(bytes[1:])
		if len(bytes) < size+2 || bytes[size+1] != '\'' {
			return "", lexutil.Location{}, nil
		}

		switch r {
		case '\t', '\n', '\'', '\\':
			return "", lexutil.Location{}, nil
		}

		numBytes = size + 2
	}

	loc := reader.Location

	bytes = bytes[:numBytes]
	n, err := reader.Read(bytes)
	if len(bytes) != n || err != nil {
		panic(err) // should never happen
	}

	return string(bytes), loc, nil
}
