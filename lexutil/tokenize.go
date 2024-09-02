package lexutil

import (
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pattyshack/gt/stringutil"
)

func IsWhitespace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

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
	numBytes := 0

	for hasMore {
		peeked, err := reader.Peek(ignorePrefixBytes + peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		if len(peeked) < ignorePrefixBytes+numBytes {
			panic("should never happen")
		}

		remaining := peeked[ignorePrefixBytes+numBytes:]

		for len(remaining) > 0 {
			utf8Char, size := utf8.DecodeRune(remaining)
			if utf8Char == utf8.RuneError {
				if len(remaining) < 4 && hasMore {
					// The rune may have been chopped off in the middle by Peek.  Read
					// more bytes and try again.
					break
				} else {
					// Encountered a real invalid utf8 byte
					return numBytes, nil
				}
			}

			if size == 0 {
				panic("should never happen")
			}

			if unicode.IsLetter(utf8Char) ||
				utf8Char == '_' ||
				(numBytes > 0 && unicode.IsNumber(utf8Char)) {

				numBytes += size
				remaining = remaining[size:]
			} else {
				return numBytes, nil
			}
		}

		peekSize *= 2
	}

	return numBytes, nil
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

// Peek for spaces of the form
//
//	(' ' | '\t')+
func PeekSpaces(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		for numBytes < len(peeked) {
			char := peeked[numBytes]
			if char == ' ' || char == '\t' {
				numBytes++
			} else {
				return numBytes, nil
			}
		}

		peekSize *= 2
	}

	return numBytes, nil
}

// Peek for newlines of the form
//
//	(\n | \r\n)+
//
// Note that \r not paired with \n is not considered a newline.
func PeekNewlines(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
) (
	int,
	int,
	bool,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0
	numNewlines := 0

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, 0, false, err
		}

		for numBytes < len(peeked) {
			char := peeked[numBytes]
			if char == '\n' {
				numBytes++
				numNewlines++

			} else if char == '\r' {
				if numBytes+1 >= len(peeked) {
					if hasMore {
						// read more bytes
						break
					} else { // '\r' not paired with '\n'
						return numBytes, numNewlines, true, nil
					}
				}

				if peeked[numBytes+1] == '\n' {
					numBytes += 2
					numNewlines++
				} else { // '\r' not paired with '\n'
					return numBytes, numNewlines, true, nil
				}
			} else {
				return numBytes, numNewlines, false, nil
			}
		}

		peekSize *= 2
	}

	return numBytes, numNewlines, false, nil
}

// Peek for comment of the form
//
//	//<optional comment string>
//
// The comment is terminated by '\n', '\r', or EOF. '\n' / '\r' / EOF is not
// part of the comment.
func PeekLineComment(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0 // "//" already peeked by peekNextToken

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		if len(peeked) < 2 {
			if hasMore {
				// read more bytes
				peekSize *= 2
				continue
			} else { // This can't be a line comment.  Not enough bytes
				return 0, nil
			}
		}

		if numBytes == 0 {
			if peeked[0] == '/' && peeked[1] == '/' {
				numBytes += 2
			} else { // Not a line comment
				return 0, nil
			}
		}

		for numBytes < len(peeked) {
			char := peeked[numBytes]
			if char == '\n' || char == '\r' {
				return numBytes, nil
			} else {
				numBytes++
			}
		}

		peekSize *= 2
	}

	return numBytes, nil
}

// Peek for block comment of the form
//
//	/*<optional comment string>*/
//
// If scoped is false, the comment terminate on the first encounter of "*/".
// If scoped is true, the comment terminate when the block comment is out of
// scope.
func PeekBlockComment(
	reader BufferedByteLocationReader,
	scoped bool,
	initialPeekWindowSize int,
) (
	int,
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0
	scopeLevel := 0

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, 0, err
		}

		if len(peeked) < 2 {
			if hasMore {
				// read more bytes
				peekSize *= 2
				continue
			} else { // This can't be a block comment.  Not enough bytes
				return 0, 0, nil
			}
		}

		if numBytes == 0 {
			if peeked[0] == '/' && peeked[1] == '*' {
				numBytes += 2
				scopeLevel = 1
			} else { // Not a block comment
				return 0, 0, nil
			}
		}

		for numBytes < len(peeked) {
			if numBytes+1 >= len(peeked) {
				if hasMore {
					// read more bytes
					break
				} else {
					// reached EOF without finding block comment's closing delimiter.
					numBytes++
					return numBytes, scopeLevel, nil
				}
			}

			char := peeked[numBytes]
			if char == '/' && peeked[numBytes+1] == '*' {
				numBytes += 2
				scopeLevel += 1

			} else if char == '*' && peeked[numBytes+1] == '/' {
				numBytes += 2
				scopeLevel -= 1

				if !scoped || scopeLevel == 0 {
					return numBytes, 0, nil
				}
			} else {
				numBytes++
			}
		}

		peekSize *= 2
	}

	return numBytes, scopeLevel, nil
}

// Strip all leading whitespaces.
func StripLeadingWhitespaces(
	reader BufferedByteLocationReader,
) error {
	modified := true
	for modified {
		modified = false

		num, err := PeekSpaces(reader, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}

		num, _, _, err = PeekNewlines(reader, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}
	}

	return nil
}

// Strip all leading whitespaces, // line comments, and scoped /**/ block
// comments.
func StripLeadingWhitespacesAndComments(
	reader BufferedByteLocationReader,
) error {
	modified := true
	for modified {
		modified = false

		num, err := PeekSpaces(reader, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}

		num, _, _, err = PeekNewlines(reader, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}

		num, err = PeekLineComment(reader, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}

		num, _, err = PeekBlockComment(reader, true, 32)
		if err != nil {
			return err
		}

		if num > 0 {
			reader.Discard(num)
			modified = true
		}
	}

	return nil
}

// Constant (non-keyword) symbol peeker.
type ConstantSymbols[T any] struct {
	symbols         map[string]T
	maxSymbolLength int

	*stringutil.InternPool
}

func NewConstantSymbols[T any](
	symbols map[string]T,
	internPool *stringutil.InternPool,
) ConstantSymbols[T] {
	s := ConstantSymbols[T]{
		symbols:    symbols,
		InternPool: internPool,
	}

	for symbol, _ := range symbols {
		internPool.Intern(symbol)
		if len(symbol) > s.maxSymbolLength {
			s.maxSymbolLength = len(symbol)
		}
	}

	return s
}

func (symbols ConstantSymbols[T]) PeekSymbol(
	reader BufferedByteLocationReader,
) (
	string,
	T,
	bool,
	error,
) {
	peeked, err := reader.Peek(symbols.maxSymbolLength)
	if err != nil && err != io.EOF {
		var defaultEntry T
		return "", defaultEntry, false, err
	}

	for len(peeked) > 0 {
		symbolStr, ok := symbols.GetInternBytes(peeked)
		if ok {
			entry, ok := symbols.symbols[symbolStr]
			if ok {
				return symbolStr, entry, ok, nil
			}
		}
		peeked = peeked[:len(peeked)-1]
	}

	var defaultEntry T
	return "", defaultEntry, false, nil
}

func (symbols ConstantSymbols[T]) MaybeTokenizeSymbol(
	reader BufferedByteLocationReader,
) (
	string,
	T,
	Location,
	error,
) {
	symbolStr, entry, found, err := symbols.PeekSymbol(reader)
	if err != nil || !found {
		return "", entry, Location{}, err
	}

	loc := reader.Location

	_, err = reader.Discard(len(symbolStr))
	if err != nil {
		panic("should never happen")
	}

	return symbolStr, entry, loc, nil
}
