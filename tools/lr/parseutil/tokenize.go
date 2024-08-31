package parseutil

import (
	"io"
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
)

func IsWhitespace(char byte) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

func stripLeadingWhitespaces(
	reader lexutil.BufferedByteLocationReader,
) (bool, error) {
	consumedBytes := false

	for {
		bytes, err := reader.Peek(1)
		if err != nil {
			return consumedBytes, err
		}

		char := bytes[0]
		if IsWhitespace(char) {
			_, err = reader.Discard(1)
			if err != nil {
				panic(err) // should never happen
			}

			consumedBytes = true
		} else {
			break
		}
	}

	return consumedBytes, nil
}

// Strip all leading whitespaces
func StripLeadingWhitespaces(reader lexutil.BufferedByteLocationReader) error {
	_, err := stripLeadingWhitespaces(reader)
	return err
}

// Strip all leading whitespaces as well as golang style comments
// (i.e., /* */  and //)
func StripLeadingWhitespacesAndComments(
	reader lexutil.BufferedByteLocationReader,
) error {
	modified := true

	singleByte := [1]byte{}

	for modified {
		modified = false

		var err error
		modified, err = stripLeadingWhitespaces(reader)
		if err != nil {
			return err
		}

		bytes, _ := reader.Peek(2)

		if string(bytes) == "//" {
			for {
				numRead, err := reader.Read(singleByte[:])
				if err != nil || numRead == 0 {
					return err
				}

				if singleByte[0] == '\n' {
					break
				}
			}

			modified = true
			continue
		}

		if string(bytes) == "/*" {
			n, err := reader.Read(bytes)
			if n != 2 || err != nil {
				panic(err) // should never happen
			}

			for {
				bytes, err = reader.Peek(2)
				if err != nil {
					return err
				}

				if string(bytes) == "*/" {
					n, err := reader.Read(bytes)
					if n != 2 || err != nil {
						panic(err) // should never happen
					}

					break
				}

				_, err = reader.Discard(1)
				if err != nil {
					panic(err)
				}
			}

			modified = true
		}
	}

	return nil
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
