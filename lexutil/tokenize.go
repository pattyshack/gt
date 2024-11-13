package lexutil

import (
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/pattyshack/gt/stringutil"
)

func IsBinaryDigit(char rune) bool {
	return char == '0' || char == '1'
}

func IsOctalDigit(char rune) bool {
	return '0' <= char && char <= '7'
}

func IsDecimalDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func IsHexadecimalDigit(char rune) bool {
	return '0' <= char && char <= '9' ||
		'A' <= char && char <= 'F' ||
		'a' <= char && char <= 'f'
}

func IsWhitespace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r'
}

func IsValidIdentifier(id string) bool {
	if id == "" {
		return false
	}

	for idx, utf8Char := range id {
		if !(unicode.IsLetter(utf8Char) ||
			utf8Char == '_' ||
			(idx > 0 && unicode.IsNumber(utf8Char))) {

			return false
		}
	}

	return true
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
func MaybeTokenizeIdentifier[SymbolId any](
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	internPool *stringutil.InternPool,
	identifierToken SymbolId,
) (
	*TokenValue[SymbolId],
	error,
) {
	size, err := PeekIdentifier(reader, 0, initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if size == 0 {
		return nil, nil
	}

	loc := reader.Location

	bytes, err := reader.Peek(size)
	if err != nil || len(bytes) != size {
		panic("should never happen")
	}
	value := internPool.InternBytes(bytes)

	_, err = reader.Discard(size)
	if err != nil {
		panic("should never happen")
	}

	return &TokenValue[SymbolId]{
		SymbolId:    identifierToken,
		StartEndPos: NewStartEndPos(loc, reader.Location),
		Value:       value,
	}, nil
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

func MaybeTokenizeSpaces[SymbolId any](
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	spacesToken SymbolId,
) (
	*TokenCount[SymbolId],
	error,
) {
	numBytes, err := PeekSpaces(reader, initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if numBytes == 0 {
		return nil, nil
	}

	loc := reader.Location
	_, err = reader.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	return &TokenCount[SymbolId]{
		SymbolId:    spacesToken,
		StartEndPos: NewStartEndPos(loc, reader.Location),
		Count:       numBytes,
	}, nil
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

func MaybeTokenizeNewlines[SymbolId any](
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	newlinesToken SymbolId,
) (
	*TokenCount[SymbolId],
	bool, // found and discarded invalid stand-alone \r
	error,
) {
	numBytes, numNewlines, foundInvalidNewline, err := PeekNewlines(
		reader,
		initialPeekWindowSize)
	if err != nil {
		return nil, false, err
	}

	loc := reader.Location
	if foundInvalidNewline {
		numBytes = 1
	}

	if numBytes == 0 {
		return nil, false, nil
	}

	_, err = reader.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	if foundInvalidNewline {
		return nil, true, nil
	}

	return &TokenCount[SymbolId]{
		SymbolId:    newlinesToken,
		StartEndPos: NewStartEndPos(loc, reader.Location),
		Count:       numNewlines,
	}, false, nil
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

func MaybeTokenizeLineComment[SymbolId any](
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	lineCommentToken SymbolId,
	preserveContent bool,
) (
	*TokenValue[SymbolId],
	error,
) {
	numBytes, err := PeekLineComment(reader, initialPeekWindowSize)
	if err != nil {
		return nil, err
	}

	if numBytes == 0 {
		return nil, nil
	}

	loc := reader.Location

	value := ""
	if preserveContent {
		peeked, err := reader.Peek(numBytes)
		if err != nil {
			panic("should never happen")
		}
		value = string(peeked)
	}

	_, err = reader.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	return &TokenValue[SymbolId]{
		SymbolId:    lineCommentToken,
		StartEndPos: NewStartEndPos(loc, reader.Location),
		Value:       value,
	}, nil
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

func MaybeTokenizeBlockComment[SymbolId any](
	reader BufferedByteLocationReader,
	scoped bool,
	initialPeekWindowSize int,
	blockCommentToken SymbolId,
	preserveContent bool,
) (
	*TokenValue[SymbolId],
	bool, // block comment not terminated
	error,
) {
	numBytes, scopeLevel, err := PeekBlockComment(
		reader,
		scoped,
		initialPeekWindowSize)
	if err != nil {
		return nil, false, err
	}

	if numBytes == 0 {
		return nil, false, nil
	}

	loc := reader.Location

	value := ""
	if preserveContent && (!scoped || scopeLevel == 0) {
		peeked, err := reader.Peek(numBytes)
		if err != nil {
			panic("should never happen")
		}
		value = string(peeked)
	}

	_, err = reader.Discard(numBytes)
	if err != nil {
		panic("should never happen")
	}

	notTerminated := scoped && scopeLevel > 0

	return &TokenValue[SymbolId]{
		SymbolId:    blockCommentToken,
		StartEndPos: NewStartEndPos(loc, reader.Location),
		Value:       value,
	}, notTerminated, nil
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

		num, scope, err := PeekBlockComment(reader, true, 32)
		if err != nil {
			return err
		}

		if scope > 0 {
			return fmt.Errorf("block comment not terminated %s", reader.Location)
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

type PeekStringResult struct {
	FoundStartMarker bool
	NumBytes         int
	ContentLength    int
	ErrorMsg         string
}

func PeekString(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	typeName string,
	skipLeadingBytes int,
	marker byte,
	markerLength int,
	allowMultiline bool,
	allowEscaped bool,
) (
	PeekStringResult,
	error,
) {
	peekSize := initialPeekWindowSize
	hasMore := true

	result := PeekStringResult{}

	checkStartMarker := true
	leadingBytes := skipLeadingBytes + markerLength
	for hasMore {
		peeked, err := reader.Peek(leadingBytes + peekSize)
		if len(peeked) > 0 && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return result, err
		}

		if checkStartMarker {
			if len(peeked) < leadingBytes {
				return result, nil
			}

			for _, b := range peeked[skipLeadingBytes:leadingBytes] {
				if b != marker {
					return result, nil
				}
			}

			result.FoundStartMarker = true
			result.NumBytes = leadingBytes
			checkStartMarker = false
		}

		remaining := peeked[result.NumBytes:]

		for len(remaining) > 0 {
			// Ensure we can process the longest rune content: \U[0-9a-fA-F]{8}
			if len(remaining) < 10 && hasMore {
				// read more bytes
				break
			}

			char := remaining[0]
			if char == marker {
				result.NumBytes++
				remaining = remaining[1:]

				count := 1
				for ; count < markerLength && len(remaining) > 0; count++ {
					if remaining[0] == marker {
						result.NumBytes++
						remaining = remaining[1:]
					} else {
						break
					}
				}

				if count == markerLength {
					return result, nil
				} else {
					result.ContentLength += count
				}
			} else if char == '\n' {
				if allowMultiline {
					result.NumBytes++
					remaining = remaining[1:]
				} else { // don't include the newline
					result.ErrorMsg = typeName + " not terminated"
					return result, nil
				}
			} else if allowEscaped && char == '\\' { // escape sequence
				result.NumBytes++
				remaining = remaining[1:]

				if len(remaining) == 0 {
					result.ErrorMsg = "invalid escaped character in " + typeName
					return result, nil
				}

				char := remaining[0]
				switch char {
				case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'', '"', '`':
					// valid escape
					// Note: if we replace '\'', '"', and '`' with marker, then this
					// would behave like golang's escape.
					result.NumBytes++
					result.ContentLength++
					remaining = remaining[1:]
				case '\n':
					if allowMultiline { // valid escape (line continuation)
						result.NumBytes++
						result.ContentLength++
						remaining = remaining[1:]
					} else { // don't include the newline as part of this token.
						result.ErrorMsg = "invalid escaped character in " + typeName
						return result, nil
					}
				default:
					result.NumBytes++
					remaining = remaining[1:]

					value := int(char - '0')
					verifyOctal := false
					isDigit := IsOctalDigit
					length := 2
					if IsOctalDigit(rune(char)) { // \[0-7]{3}
						// need to check the remaining 2 octal bytes, and
						verifyOctal = true
					} else if char == 'x' { // \x[0-9a-fA-F]{2}
						isDigit = IsHexadecimalDigit
					} else if char == 'u' { // \u[0-9a-fA-F]{4}
						isDigit = IsHexadecimalDigit
						length = 4
					} else if char == 'U' { // \U[0-9a-fA-F]{8}
						isDigit = IsHexadecimalDigit
						length = 8
					} else { // invalid escape
						result.ErrorMsg = "invalid escaped character in " + typeName
						return result, nil
					}

					count := 0
					for ; count < length && len(remaining) > 0; count++ {
						if isDigit(rune(remaining[0])) {
							if verifyOctal {
								value <<= 3
								value |= int(remaining[0] - '0')
							}

							result.NumBytes++
							remaining = remaining[1:]
						} else {
							break
						}
					}

					if count == length {
						result.ContentLength++
					} else {
						result.ErrorMsg = "invalid escaped unicode value in " + typeName
						return result, nil
					}

					if verifyOctal && value > 255 {
						result.ErrorMsg = fmt.Sprintf(
							"invalid escaped octal value (%d > 255)  in %s",
							value,
							typeName)
						return result, nil
					}
				}
			} else {
				utf8Char, size := utf8.DecodeRune(remaining)

				result.NumBytes += size
				remaining = remaining[size:]

				if utf8Char == utf8.RuneError {
					result.ErrorMsg = "invalid unicode rune in " + typeName
					return result, nil
				} else {
					result.ContentLength++
				}
			}
		}

		peekSize *= 2
	}

	result.ErrorMsg = typeName + " not terminated"
	return result, nil
}

func peekDigits(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	offset int,
	isDigit func(rune) bool,
	requireLeadingDigit bool,
) (
	int,
	error,
) {
	hasMore := true
	peekSize := initialPeekWindowSize
	numBytes := 0

	for hasMore {
		peeked, err := reader.Peek(offset + peekSize)
		if len(peeked) >= offset && err == io.EOF {
			hasMore = false
			err = nil
		}
		if err != nil {
			return 0, err
		}

		remaining := peeked[offset+numBytes:]
		for len(remaining) > 0 {
			char := remaining[0]
			if isDigit(rune(char)) {
				numBytes++
				remaining = remaining[1:]
			} else if char == '_' {
				if requireLeadingDigit && numBytes == 0 {
					return numBytes, nil
				}

				if len(remaining) < 2 {
					if hasMore {
						// read more bytes
						break
					} else { // the int is followed by a '_'
						return numBytes, nil
					}
				}

				if isDigit(rune(remaining[1])) {
					numBytes += 2
					remaining = remaining[2:]
				} else { // the int is followed by a '_'
					return numBytes, nil
				}
			} else {
				return numBytes, nil
			}
		}

		peekSize *= 2
	}

	return numBytes, nil
}

func peekFloat(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
	leadingIntBytes int,
	isHex bool, // false = decimal
) (
	int,
	error,
) {
	isDigit := IsDecimalDigit
	intPrefixLen := 0 // no prefix
	lowerExp := byte('e')
	upperExp := byte('E')
	if isHex {
		isDigit = IsHexadecimalDigit
		intPrefixLen = 2 // "0x" or "0X"
		lowerExp = 'p'
		upperExp = 'P'
	}

	// peek for (hexa)decimal points of the form DOT digits

	floatBytes := leadingIntBytes
	peeked, err := reader.Peek(floatBytes + 1)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, err
	}

	if len(peeked) <= floatBytes {
		// can't be a float without decimal point or exponent
		return 0, nil
	}

	if peeked[floatBytes] == '.' {
		floatBytes++

		digitBytes, err := peekDigits(
			reader,
			initialPeekWindowSize,
			floatBytes,
			isDigit,
			true)
		if err != nil {
			return 0, err
		}

		if digitBytes == 0 && leadingIntBytes == intPrefixLen {
			// ".", without leading and trailing digits, is not a valid float
			return 0, err
		}

		floatBytes += digitBytes
	}

	// peek for exponent of the form [eEpP][+-]?digits

	peeked, err = reader.Peek(floatBytes + 2)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return 0, err
	}

	if len(peeked) <= floatBytes {
		if isHex { // exponent is not optional for hexadecimal
			return 0, nil
		} else if floatBytes > leadingIntBytes {
			// a valid decimal float without exponent
			return floatBytes, nil
		}
	}

	exp := peeked[floatBytes]
	if exp == lowerExp || exp == upperExp {
		prefix := 1
		if len(peeked) == floatBytes+2 {
			char := peeked[floatBytes+1]
			if char == '+' || char == '-' {
				prefix = 2
			}
		}

		exponentDigits, err := peekDigits(
			reader,
			initialPeekWindowSize,
			floatBytes+prefix,
			isDigit,
			true)
		if err != nil {
			return 0, err
		}

		if exponentDigits > 0 {
			floatBytes += prefix + exponentDigits
		} else if isHex { // exponent is not optional for hexadecimal
			return 0, err
		}
	} else if isHex { // exponent is not optional for hexadecimal
		return 0, err
	}

	if floatBytes > leadingIntBytes {
		return floatBytes, nil
	}

	return 0, nil
}

type LiteralSubType string

func (t LiteralSubType) String() string { return string(t) }

const (
	DecimalInteger            = LiteralSubType("decimal integer")
	HexadecimalInteger        = LiteralSubType("hexadecimal integer")
	ZeroOPrefixedOctalInteger = LiteralSubType("0o-prefixed octal integer")
	ZeroPrefixedOctalInteger  = LiteralSubType("0-prefixed octal integer")
	BinaryInteger             = LiteralSubType("binary integer")

	DecimalFloat     = LiteralSubType("decimal float")
	HexadecimalFloat = LiteralSubType("hexadecimal float")

	SingleLineString    = LiteralSubType("single line string")
	MultiLineString     = LiteralSubType("mutli line string")
	RawSingleLineString = LiteralSubType("raw single line string")
	RawMultiLineString  = LiteralSubType("raw mutli line string")
)

type PeekIntegerOrFloatResult struct {
	NumBytes   int
	IsNegative bool
	IsFloat    bool
	SubType    LiteralSubType
	HasDigits  bool
}

func PeekIntegerOrFloat(
	reader BufferedByteLocationReader,
	initialPeekWindowSize int,
) (
	PeekIntegerOrFloatResult,
	error,
) {
	peeked, err := reader.Peek(3)
	if len(peeked) > 0 && err == io.EOF {
		err = nil
	}
	if err != nil {
		return PeekIntegerOrFloatResult{}, err
	}

	isNegative := false
	totalBytes := 0

	if peeked[0] == '-' {
		isNegative = true
		totalBytes = 1
		peeked = peeked[1:]
	}

	if len(peeked) == 0 {
		return PeekIntegerOrFloatResult{}, err
	}

	char := peeked[0]
	if char == '.' {
		totalFloatBytes, err := peekFloat(
			reader,
			initialPeekWindowSize,
			totalBytes,
			false)
		if err != nil {
			return PeekIntegerOrFloatResult{}, err
		}

		if totalFloatBytes == 0 {
			return PeekIntegerOrFloatResult{}, nil
		} else {
			return PeekIntegerOrFloatResult{
				NumBytes:   totalFloatBytes,
				IsNegative: isNegative,
				IsFloat:    true,
				SubType:    DecimalFloat,
				HasDigits:  true,
			}, nil
		}
	}

	if !IsDecimalDigit(rune(char)) {
		return PeekIntegerOrFloatResult{}, nil
	}

	subType := DecimalInteger
	hasDigits := true
	totalBytes++

	if char != '0' {
		numDigits, err := peekDigits(
			reader,
			initialPeekWindowSize,
			totalBytes,
			IsDecimalDigit,
			false)
		if err != nil {
			return PeekIntegerOrFloatResult{}, err
		}

		totalBytes += numDigits
	} else if len(peeked) > 1 {
		switch peeked[1] {
		case 'b', 'B':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				IsBinaryDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = BinaryInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		case 'o', 'O':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				IsOctalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = ZeroOPrefixedOctalInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		case 'x', 'X':
			totalBytes++
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				IsHexadecimalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			subType = HexadecimalInteger
			hasDigits = numDigits != 0
			totalBytes += numDigits
		default:
			numDigits, err := peekDigits(
				reader,
				initialPeekWindowSize,
				totalBytes,
				IsOctalDigit,
				false)
			if err != nil {
				return PeekIntegerOrFloatResult{}, err
			}

			if numDigits > 0 { // otherwise is a decimal "0"
				subType = ZeroPrefixedOctalInteger
				totalBytes += numDigits
			}
		}
	}

	result := PeekIntegerOrFloatResult{
		NumBytes:   totalBytes,
		IsNegative: isNegative,
		SubType:    subType,
		HasDigits:  hasDigits,
	}

	if subType == DecimalInteger || subType == HexadecimalInteger {
		isHex := subType == HexadecimalInteger
		totalFloatBytes, err := peekFloat(
			reader,
			initialPeekWindowSize,
			totalBytes,
			isHex)
		if err != nil {
			return result, nil // we still have a valid integer
		}

		if totalFloatBytes > 0 {
			result.NumBytes = totalFloatBytes
			result.IsFloat = true
			result.HasDigits = true
			result.SubType = DecimalFloat
			if isHex {
				result.SubType = HexadecimalFloat
			}
		}
	}

	return result, nil
}
