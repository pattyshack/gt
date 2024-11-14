package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
)

const (
	textPeekSize    = 1024
	initialPeekSize = 64
)

var (
	directiveRegexp = regexp.MustCompile(`\[\[|\$`)

	directiveSymbols = map[string]SymbolId{
		"$": CopySectionToken,
		"#": CommentToken,
	}

	// NOTE: for, $, and # directives can be both parameterless and parametered

	parameterlessOnlyDirectives = map[string]struct{}{
		"end":      struct{}{},
		"default":  struct{}{},
		"else":     struct{}{},
		"continue": struct{}{},
		"break":    struct{}{},
		"return":   struct{}{},
	}

	parameteredOnlyDirectives = map[string]struct{}{
		"switch":  struct{}{},
		"case":    struct{}{},
		"if":      struct{}{},
		"else if": struct{}{},
		"error":   struct{}{},
		"embed":   struct{}{},
	}

	NonIdentifierChars = map[byte]struct{}{}
)

func init() {
	for _, c := range "`~!@#$%^&*()-=+{[]}\\|;:'\",<.>/? \t\n" {
		NonIdentifierChars[byte(c)] = struct{}{}
	}
}

type BodyToken interface {
	Token

	// When true, and the previous statement is text, remove the whitespaces
	// in the text that are adjacent to this statement, potentially up to and
	// including the previous line's newline character.
	TrimLeadingWhitespaces() bool

	// When true, and the next statement is text, remove the whitespaces
	// in the text taht are adjacent to this statement, potentially up to and
	// including the current line's newline character.
	TrimTrailingWhitespaces() bool
}

type rawBodyLexer struct {
	reader lexutil.BufferedByteLocationReader

	internPool *stringutil.InternPool

	lexutil.ConstantSymbols[SymbolId]
}

func newRawBodyLexer(
	reader lexutil.BufferedByteLocationReader,
	internPool *stringutil.InternPool,
) *rawBodyLexer {
	return &rawBodyLexer{
		reader:          reader,
		internPool:      internPool,
		ConstantSymbols: lexutil.NewConstantSymbols(directiveSymbols, internPool),
	}
}

func (lexer *rawBodyLexer) CurrentLocation() lexutil.Location {
	return lexer.reader.Location
}

func (lexer *rawBodyLexer) Next() (BodyToken, error) {
	token, err := lexer.maybeTokenizeText()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeDirective()
	if token != nil || err != nil {
		return token, err
	}

	return nil, lexutil.NewLocationError(
		lexer.reader.Location,
		"Unexpected character")
}

func (lexer *rawBodyLexer) maybeTokenizeText() (BodyToken, error) {
	loc := lexer.reader.Location
	buffer := bytes.NewBuffer(nil)

	foundEnd := false
	for !foundEnd {
		bytes, err := lexer.reader.Peek(textPeekSize)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF && len(bytes) == 0 {
			if buffer.Len() == 0 {
				return nil, io.EOF
			}

			break
		}

		index := directiveRegexp.FindIndex(bytes)

		endIdx := 0
		if index == nil {
			// The whole chunk does not contain any directive

			endIdx = len(bytes)
			if bytes[endIdx-1] == '[' && len(bytes) == textPeekSize {
				// NOTE: There's a chance that the last character in this
				// chunk belongs to a different token.  Don't consume the
				// '[' and read more bytes to verify.
				endIdx -= 1
			}
		} else {
			endIdx = index[0]

			if bytes[endIdx] == '[' { // found directive / end of text
				if bytes[endIdx+1] != '[' { // sanity check
					panic("Programming error")
				}

				foundEnd = true
			} else if bytes[endIdx] == '$' {
				if endIdx+1 < len(bytes)-1 {
					if bytes[endIdx+1] == '$' { // escaped $
						endIdx += 2
					} else {
						foundEnd = true
					}
				} else {
					if len(bytes) == textPeekSize {
						// NOTE: There's a chance that the last character
						// in // this chunk is an escaped '$'.  Don't
						// consume the '$' and read more bytes to verify.
					} else {
						// last
						foundEnd = true
					}
				}

			} else {
				panic("Programming error")
			}
		}

		if endIdx > 0 {
			_, err = buffer.Write(bytes[:endIdx])
			if err != nil {
				return nil, err
			}

			_, err = lexer.reader.Read(bytes[:endIdx])
			if err != nil {
				panic(err) // should never happen
			}
		}
	}

	if buffer.Len() == 0 {
		return nil, nil
	}

	value := string(buffer.Bytes())

	// unescape $
	value = strings.ReplaceAll(value, "$$", "$")

	// "escape" ` since the text block is output as raw string literal.
	// Note that there's no clean way to escape ` in raw string.
	// https://github.com/golang/go/issues/24475
	value = strings.ReplaceAll(value, "`", "`+\"`\"+`")

	pos := lexutil.NewStartEndPos(loc, lexer.reader.Location)
	return NewAtom(TextToken, pos, value, false, false), nil
}

func (lexer *rawBodyLexer) tokenizeNonSubstituteDirective() (BodyToken, error) {
	content, loc, err := readDirective(lexer.reader, 2, "]]")
	if err != nil {
		return nil, err
	}
	content = content[2 : len(content)-2]

	trimLeading := false
	if len(content) > 0 && content[0] == '-' {
		trimLeading = true
		content = content[1:]
	}

	trimTrailing := false
	if len(content) > 0 && content[len(content)-1] == '-' {
		trimTrailing = true
		content = content[:len(content)-1]
	}

	directiveReader := lexutil.NewBufferedByteLocationReaderFromSlice("", content)
	directiveReader.Location = loc

	err = lexutil.StripLeadingWhitespaces(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	token, err := lexer.MaybeTokenizeSymbol(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if token != nil {
		content, err = ioutil.ReadAll(directiveReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		pos := lexutil.NewStartEndPos(loc, lexer.reader.Location)
		return NewAtom(
			SymbolId(token.SymbolId),
			pos,
			string(content),
			trimLeading,
			trimTrailing), nil
	}

	idToken, err := lexutil.MaybeTokenizeIdentifier(
		directiveReader,
		initialPeekSize,
		lexer.internPool,
		"")
	if err != nil && err != io.EOF {
		return nil, err
	}

	id := ""
	if idToken != nil {
		id = idToken.Value
	}

	err = lexutil.StripLeadingWhitespaces(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	param := ""
	// check for "else if" compound identifier
	if id == "else" {
		secondToken, err := lexutil.MaybeTokenizeIdentifier(
			directiveReader,
			initialPeekSize,
			lexer.internPool,
			"")
		if err != nil && err != io.EOF {
			return nil, err
		}

		second := ""
		if secondToken != nil {
			second = secondToken.Value
		}

		if second == "if" {
			id = "else if"

			err = lexutil.StripLeadingWhitespaces(directiveReader)
			if err != nil && err != io.EOF {
				return nil, err
			}
		} else {
			param = second
		}
	}

	content, err = ioutil.ReadAll(directiveReader)
	if err != nil && err != io.EOF {
		return nil, err
	}

	param += string(content)

	_, ok := parameterlessOnlyDirectives[id]
	if ok && len(param) > 0 {
		return nil, lexutil.NewLocationError(
			loc,
			"unexpected parameter specified in [[%s]] directive",
			id)
	}

	_, ok = parameteredOnlyDirectives[id]
	if ok && param == "" {
		return nil, lexutil.NewLocationError(
			loc,
			"expected parameter not specified in [[%s]] directive",
			id)
	}

	pos := lexutil.NewStartEndPos(loc, lexer.reader.Location)
	switch id {
	case "":
		return nil, lexutil.NewLocationError(
			loc,
			"invalid directive. directive type not specified")
	case "end":
		return NewTToken(EndToken, pos, trimLeading, trimTrailing), nil
	case "default":
		return NewTToken(DefaultToken, pos, trimLeading, trimTrailing), nil
	case "else":
		return NewTToken(ElseToken, pos, trimLeading, trimTrailing), nil
	case "for":
		return NewValue(ForToken, pos, param, trimLeading, trimTrailing), nil
	case "switch":
		return NewValue(SwitchToken, pos, param, trimLeading, trimTrailing), nil
	case "case":
		return NewValue(CaseToken, pos, param, trimLeading, trimTrailing), nil
	case "if":
		return NewValue(IfToken, pos, param, trimLeading, trimTrailing), nil
	case "else if":
		return NewValue(ElseIfToken, pos, param, trimLeading, trimTrailing), nil
	case "continue":
		return NewAtom(
			ContinueToken,
			pos,
			param,
			trimLeading,
			trimTrailing), nil
	case "break":
		return NewAtom(BreakToken, pos, param, trimLeading, trimTrailing), nil
	case "return":
		return NewAtom(ReturnToken, pos, param, trimLeading, trimTrailing), nil
	case "error":
		return NewAtom(ErrorToken, pos, param, trimLeading, trimTrailing), nil
	case "embed":
		return NewAtom(EmbedToken, pos, param, trimLeading, trimTrailing), nil
	}

	return nil, lexutil.NewLocationError(
		loc,
		"invalid directive. unknown directive type %s",
		id)
}

func (lexer *rawBodyLexer) maybeTokenizeDirective() (BodyToken, error) {
	content, err := lexer.reader.Peek(2)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if len(content) < 2 {
		return nil, nil
	}

	if content[0] == '[' {
		if content[1] != '[' {
			panic("Programming error")
		}

		return lexer.tokenizeNonSubstituteDirective()

	} else if content[0] == '$' {
		_, ok := NonIdentifierChars[content[1]]
		if !ok { // $<identifier>
			loc := lexer.reader.Location

			_, err := lexer.reader.Discard(1)
			if err != nil {
				panic(err) // Should never happen
			}

			value, err := lexutil.MaybeTokenizeIdentifier(
				lexer.reader,
				initialPeekSize,
				lexer.internPool,
				"")
			if err != nil {
				return nil, err
			}

			val := ""
			if value != nil {
				val = value.Value
			}

			pos := lexutil.NewStartEndPos(loc, lexer.reader.Location)
			return NewAtom(
				SubstitutionToken,
				pos,
				val,
				false,
				false), nil

		} else if content[1] == '(' {

			content, loc, err := readDirective(lexer.reader, 2, ")")
			if err != nil {
				return nil, err
			}

			value := string(content[2 : len(content)-1])
			pos := lexutil.NewStartEndPos(loc, lexer.reader.Location)
			return NewAtom(SubstitutionToken, pos, value, false, false), nil

		} else if content[1] == '$' {
			panic("Programming error")
		} else {
			return nil, lexutil.NewLocationError(
				lexer.reader.Location,
				"invalid substitute directive")
		}
	}

	panic("Programming error")
}

// Search for the first occurence of terminal starting from startIdx.  The
// terminal must start with one of the following characters: ']' ')' '}' '\n'.
// This respect golang scoping, string, char and comments, i.e.,
// {} [] () “ "" ” /**/ //
func readDirective(
	reader lexutil.BufferedByteLocationReader,
	startIdx int,
	terminal string,
) (
	[]byte,
	lexutil.Location,
	error,
) {

	if terminal == "" {
		return nil, lexutil.Location{}, lexutil.NewLocationError(
			reader.Location,
			"Invalid terminal")
	}

	if terminal[0] != '}' &&
		terminal[0] != ')' &&
		terminal[0] != ']' &&
		terminal[0] != '\n' {
		return nil, lexutil.Location{}, lexutil.NewLocationError(
			reader.Location,
			"Invalid terminal: "+terminal)
	}

	currentScope := []string{"#root"}

	peekRange := 64
	prevLen := 0
	checkIdx := startIdx

	bytes, err := reader.Peek(peekRange)
	if err != nil && err != io.EOF {
		return nil, lexutil.Location{}, lexutil.LocationError{
			Loc: reader.Location,
			Err: err,
		}
	}

	if len(bytes) <= startIdx {
		return nil, lexutil.Location{}, lexutil.NewLocationError(
			reader.Location,
			"lex error: \"%s\" not found",
			terminal)
	}

	for {
		needMoreBytes := false
		for !needMoreBytes && checkIdx < len(bytes) {
			char := bytes[checkIdx]

			top := currentScope[len(currentScope)-1]
			switch top {
			case "#root", "{", "[", "(": // nestable scopes
				switch char {
				case '{', '[', '(', '`', '"', '\'':
					currentScope = append(currentScope, string(char))
				case '/':
					if len(bytes) <= checkIdx+1 {
						needMoreBytes = true
						continue
					} else if bytes[checkIdx+1] == '*' ||
						bytes[checkIdx+1] == '/' {

						currentScope = append(currentScope, string(bytes[checkIdx:checkIdx+2]))
						checkIdx += 1
					}
				case '}', ']', ')', '\n':
					if (top == "(" && char == ')') ||
						(top == "{" && char == '}') ||
						(top == "[" && char == ']') {

						currentScope = currentScope[:len(currentScope)-1]
						checkIdx += 1
						continue
					}

					if top == "#root" {
						if len(bytes) <= checkIdx+len(terminal)-1 {
							needMoreBytes = true
							continue
						}

						if terminal == string(bytes[checkIdx:checkIdx+len(terminal)]) {
							loc := reader.Location

							bytes = make([]byte, checkIdx+len(terminal))
							n, err := reader.Read(bytes)
							if len(bytes) != n || err != nil {
								panic(err) // should never happen
							}

							return bytes, loc, nil
						}
					}

					if char != '\n' {
						bytes = bytes[:checkIdx]

						n, err := reader.Read(bytes)
						if len(bytes) != n || err != nil {
							panic(err) // should never happen
						}

						return nil, lexutil.Location{}, lexutil.NewLocationError(
							reader.Location,
							"lex error: no matching pair for %c",
							char)
					}
				}
			case "//":
				if char == '\n' {
					currentScope = currentScope[:len(currentScope)-1]
					// don't include '\n' since it could be part of the terminal
					checkIdx -= 1
				}
			case "/*":
				if char == '*' {
					if len(bytes) <= checkIdx+1 {
						needMoreBytes = true
						continue
					} else if bytes[checkIdx+1] == '/' {
						checkIdx += 1
						currentScope = currentScope[:len(currentScope)-1]
					}
				}
			case "`", "'", "\"":
				if char == '\\' {
					checkIdx += 1 // skip escaped char
				} else if char == top[0] {
					currentScope = currentScope[:len(currentScope)-1]
				}
			default:
				panic(fmt.Sprintf("Programming error: %v", currentScope))
			}

			if !needMoreBytes {
				checkIdx += 1
			}
		}

		peekRange *= 2
		prevLen = len(bytes)

		bytes, err = reader.Peek(peekRange)
		if err != nil && err != io.EOF {
			return nil, lexutil.Location{}, err
		}

		if prevLen == len(bytes) { // not found
			return nil, lexutil.Location{}, lexutil.NewLocationError(
				reader.Location,
				"lex error: \"%s\" not found",
				terminal)
		}
	}
}
