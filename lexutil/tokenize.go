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
	numSpaceBytes := 0

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		for numSpaceBytes < len(peeked) {
			char := peeked[numSpaceBytes]
			if char == ' ' || char == '\t' {
				numSpaceBytes++
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	return numSpaceBytes, nil
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
	numNewlineBytes := 0
	numNewlines := 0
	foundInvalidNewline := false

	for hasMore {
		peeked, err := reader.Peek(peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, 0, false, err
		}

		for numNewlineBytes < len(peeked) {
			char := peeked[numNewlineBytes]
			if char == '\n' {
				numNewlineBytes++
				numNewlines++

			} else if char == '\r' {
				if numNewlineBytes+1 >= len(peeked) {
					// If hasMore is true, then we may have read half of the \r\n pair.
					// The next outer loop iteration will take care of this. If hasMore
					// is false, then this is a lone \r and it should not be included
					// as part of this token.
					break
				}

				if peeked[numNewlineBytes+1] == '\n' {
					numNewlineBytes += 2
					numNewlines++
				} else { // '\r' not paired with '\n'
					foundInvalidNewline = true
					hasMore = false
					break
				}
			} else {
				hasMore = false
				break
			}
		}

		peekSize *= 2
	}

	return numNewlineBytes, numNewlines, foundInvalidNewline, nil
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
	numCommentBytes := 0 // "//" already peeked by peekNextToken

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

		if numCommentBytes == 0 {
			if peeked[0] == '/' && peeked[1] == '/' {
				numCommentBytes += 2
			} else { // Not a line comment
				return 0, nil
			}
		}

		for numCommentBytes < len(peeked) {
			char := peeked[numCommentBytes]
			if char == '\n' || char == '\r' {
				hasMore = false
				break
			} else {
				numCommentBytes++
			}
		}

		peekSize *= 2
	}

	return numCommentBytes, nil
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
	numCommentBytes := 0
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

		if numCommentBytes == 0 {
			if peeked[0] == '/' && peeked[1] == '*' {
				numCommentBytes += 2
				scopeLevel = 1
			} else { // Not a block comment
				return 0, 0, nil
			}
		}

		for numCommentBytes < len(peeked) {
			if numCommentBytes+1 >= len(peeked) {
				if hasMore {
					// read more bytes
					break
				} else {
					// reached EOF without finding block comment's closing delimiter.
					numCommentBytes++
					break
				}
			}

			char := peeked[numCommentBytes]
			if char == '/' && peeked[numCommentBytes+1] == '*' {
				numCommentBytes += 2
				scopeLevel += 1

			} else if char == '*' && peeked[numCommentBytes+1] == '/' {
				numCommentBytes += 2
				scopeLevel -= 1

				if !scoped || scopeLevel == 0 {
					hasMore = false
					break
				}
			} else {
				numCommentBytes++
			}
		}

		peekSize *= 2
	}

	return numCommentBytes, scopeLevel, nil
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
