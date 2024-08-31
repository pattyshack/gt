package parseutil

import (
	"unicode/utf8"

	"github.com/pattyshack/gt/lexutil"
)

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
