package lexutil

import (
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pattyshack/gt/stringutil"
)

// Peek for identifier of the form
//
//	(unicode-letter | '_') (unicode-letter | unicode-number | '_')*
//
// skipping over the ignorePrefixBytes.
func PeekIdentifier(
	reader BufferedByteLocationReader,
	ignorePrefixBytes int,
	initialPeekWindowSize int,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numIdentifierBytes := 0

	for hasMore {
		peeked, err := reader.Peek(ignorePrefixBytes + peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		if len(peeked) < ignorePrefixBytes+numIdentifierBytes {
			panic("should never happen")
		}

		remaining := peeked[ignorePrefixBytes+numIdentifierBytes:]

		for len(remaining) > 0 {
			utf8Char, size := utf8.DecodeRune(remaining)
			if utf8Char == utf8.RuneError {
				if len(remaining) < 4 && hasMore {
					// The rune may have been chopped off in the middle by Peek.  Read
					// more bytes and try again.
					break
				} else {
					// Encountered a real invalid utf8 byte
					hasMore = false
					break
				}
			}

			if size == 0 {
				panic("should never happen")
			}

			if unicode.IsLetter(utf8Char) ||
				utf8Char == '_' ||
				(numIdentifierBytes > 0 && unicode.IsNumber(utf8Char)) {

				numIdentifierBytes += size
				remaining = remaining[size:]
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	return numIdentifierBytes, nil
}

// If the reader's leading bytes are an identifier, read the identifer bytes
// and return the identifier value.  Otherwise, return ""
func MaybeTokenizeIdentifier(
	reader BufferedByteLocationReader,
	internPool *stringutil.InternPool,
) (
	string,
	Location,
	error,
) {
	size, err := PeekIdentifier(reader, 0, 32)
	if err != nil {
		return "", Location{}, err
	}

	if size == 0 {
		return "", Location{}, nil
	}

	loc := reader.Location

	bytes, err := reader.Peek(size)
	if err != nil || len(bytes) != size {
		panic("should never happen")
	}
	value := internPool.InternBytes(bytes)

	reader.Discard(size)

	return value, loc, nil
}
