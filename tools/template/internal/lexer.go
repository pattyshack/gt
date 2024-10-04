package template

import (
	"io"
	"io/ioutil"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
)

var (
	sectionMarker = map[string]struct{}{
		"%%\n": struct{}{},
		"%%":   struct{}{},
	}

	importMarker = map[string]struct{}{
		"(": struct{}{},
	}

	templateDeclMarker = map[string]struct{}{
		"{": struct{}{},
	}
)

type currentLexer interface {
	Next() (Token, error)
}

type LexerImpl struct {
	reader lexutil.BufferedByteLocationReader

	currentLexer

	internPool *stringutil.InternPool
}

func NewLexer(filename string, input io.Reader) (Lexer, error) {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	content = stripHeaderComments(content)

	reader := lexutil.NewBufferedByteLocationReaderFromSlice(
		filename,
		content)
	internPool := stringutil.NewInternPool()

	return &LexerImpl{
		reader: reader,
		currentLexer: &headerLexer{
			reader:        reader,
			internPool:    internPool,
			sectionMarker: lexutil.NewConstantSymbols(sectionMarker, internPool),
			importMarker:  lexutil.NewConstantSymbols(importMarker, internPool),
			templateDeclMarker: lexutil.NewConstantSymbols(
				templateDeclMarker,
				internPool),
		},
		internPool: internPool,
	}, nil
}

func (lexer *LexerImpl) CurrentLocation() Location {
	return lexer.reader.Location
}

func (lexer *LexerImpl) Next() (Token, error) {
	token, err := lexer.currentLexer.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() == SectionMarkerToken {
		lexer.currentLexer = &bodyLexer{
			raw: newRawBodyLexer(lexer.reader, lexer.internPool),
		}
	}

	return token, nil
}

type headerLexer struct {
	reader     lexutil.BufferedByteLocationReader
	internPool *stringutil.InternPool

	sectionMarker      lexutil.ConstantSymbols[struct{}]
	importMarker       lexutil.ConstantSymbols[struct{}]
	templateDeclMarker lexutil.ConstantSymbols[struct{}]
}

func (lexer *headerLexer) Next() (Token, error) {
	err := lexutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	val, loc, err := lexutil.MaybeTokenizeIdentifier(
		lexer.reader,
		lexer.internPool)
	if err != nil {
		return nil, err
	}

	switch string(val) {
	case "package":
		return lexer.tokenizePackage(loc)
	case "import":
		return lexer.tokenizeImport(loc)
	case "template":
		return lexer.tokenizeTemplateDecl(loc)
	case "":
		// try to tokenize symbol below
	default:
		return nil, lexutil.NewLocationError(
			loc,
			"Unexpected IDENTIFIER %s",
			string(val))
	}

	symbolStr, _, loc, err := lexer.sectionMarker.MaybeTokenizeSymbol(
		lexer.reader)
	if err != nil {
		return nil, err
	}

	if symbolStr != "" {
		return GenericSymbol{
			SymbolId: SectionMarkerToken,
			StartPos: loc,
		}, nil
	}

	return nil, lexutil.NewLocationError(
		lexer.reader.Location,
		"Unexpected character")
}

func (lexer *headerLexer) tokenizePackage(pkgLoc Location) (Token, error) {
	err := lexutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	val, _, err := lexutil.MaybeTokenizeIdentifier(
		lexer.reader,
		lexer.internPool)
	if err != nil {
		return nil, err
	}

	if val != "" {
		return NewValue(PackageToken, pkgLoc, val, false, false), nil
	}

	return nil, lexutil.NewLocationError(
		lexer.reader.Location,
		"Unexpected character")
}

func (lexer *headerLexer) tokenizeImport(importLoc Location) (Token, error) {
	err := lexutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	symbolStr, _, _, err := lexer.importMarker.MaybeTokenizeSymbol(
		lexer.reader)
	if err != nil {
		return nil, err
	}

	if symbolStr == "" {
		return nil, lexutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	value, _, err := readDirective(lexer.reader, 0, ")")
	if err != nil {
		return nil, err
	}

	value = value[:len(value)-1]

	return NewValue(ImportToken, importLoc, string(value), false, false), nil
}

func (lexer *headerLexer) tokenizeTemplateDecl(
	declLoc Location) (
	Token,
	error) {

	err := lexutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	templateName, _, err := lexutil.MaybeTokenizeIdentifier(
		lexer.reader,
		lexer.internPool)
	if err != nil {
		return nil, err
	}

	if templateName == "" {
		return nil, lexutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	err = lexutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	lcurl, _, _, err := lexer.templateDeclMarker.MaybeTokenizeSymbol(
		lexer.reader)
	if err != nil {
		return nil, err
	}

	if lcurl == "" {
		return nil, lexutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	body, loc, err := readDirective(lexer.reader, 0, "}")
	if err != nil {
		return nil, err
	}

	body = body[:len(body)-1]

	declReader := lexutil.NewBufferedByteLocationReaderFromSlice("", body)
	declReader.Location = loc

	args := []Argument{}
	for {
		err := lexutil.StripLeadingWhitespaces(declReader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		head, err := declReader.Peek(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(head) == 0 {
			break
		}

		line, loc, err := readDirective(declReader, 0, "\n")
		if err != nil {
			return nil, err
		}

		line = line[:len(line)-1]
		lineReader := lexutil.NewBufferedByteLocationReaderFromSlice("", line)
		lineReader.Location = loc

		argName, _, err := lexutil.MaybeTokenizeIdentifier(
			lineReader,
			lexer.internPool)
		if err != nil {
			return nil, err
		}

		if argName == "" {
			return nil, lexutil.NewLocationError(
				lineReader.Location,
				"Expecting argument name")
		}

		err = lexutil.StripLeadingWhitespaces(lineReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		typeName, err := ioutil.ReadAll(lineReader)
		if err != nil {
			return nil, err
		}

		if len(typeName) == 0 {
			return nil, lexutil.NewLocationError(
				lineReader.Location,
				"Expecting argument type")
		}

		args = append(args, Argument{string(argName), string(typeName)})
	}

	return NewTemplateDeclaration(declLoc, templateName, args), nil
}

// Strip all comments from the header section of template.  Note that
// the content is modify in place.
func stripHeaderComments(template []byte) []byte {
	curr := 0
	shifted := 0

	skip := func() {
		curr += 1
	}

	shift := func() {
		template[shifted] = template[curr]
		shifted += 1
		curr += 1
	}

	// XXX maybe handle `` string and '' char as well
	inString := false
	inLineComment := false
	inBlockComment := false

	for curr < len(template) {
		char := template[curr]

		if inString {
			shift()

			if char == '\\' {
				if curr < len(template) {
					shift() // shift escaped character
				}
			} else if char == '"' {
				inString = false
			}
		} else if inLineComment {
			if char == '\n' { // preserve '\n' in the code
				inLineComment = false
			} else {
				skip()
			}
		} else if inBlockComment {
			if char == '\n' {
				// We need to preserve \n to ensure token locations are correct
				shift()
			} else {
				skip()
			}

			if char == '*' &&
				curr < len(template) &&
				template[curr] == '/' {

				skip()
				inBlockComment = false
			}
		} else {
			if char == '/' && curr+1 < len(template) {
				if template[curr+1] == '/' {
					skip()
					skip()
					inLineComment = true
				} else if template[curr+1] == '*' {
					skip()
					skip()
					inBlockComment = true
				} else {
					shift()
				}
			} else if char == '"' {
				shift()
				inString = true
			} else if char == '%' &&
				curr+1 < len(template) &&
				template[curr+1] == '%' {

				// Reached section marker.  Leave the body unmodified.
				for curr < len(template) {
					shift()
				}
			} else {
				shift()
			}
		}
	}

	return template[:shifted]
}

type bodyLexer struct {
	raw *rawBodyLexer

	lookAhead    []BodyToken
	lookAheadErr error
}

func (lexer *bodyLexer) fillTokensAndMaybeTrimWhitespaces() {
	if lexer.lookAheadErr != nil {
		return
	}

	for len(lexer.lookAhead) < 2 {
		token, err := lexer.raw.Next()
		if err != nil {
			lexer.lookAheadErr = err
			return
		}

		lexer.lookAhead = append(lexer.lookAhead, token)
	}

	if lexer.lookAhead[0].Id() == TextToken &&
		lexer.lookAhead[1].TrimLeadingWhitespaces() {

		text := lexer.lookAhead[0].(*Atom)
		length := len(text.Value)

		for length > 0 {
			char := text.Value[length-1]
			if char == ' ' || char == '\t' {
				length -= 1
				continue
			}

			if char == '\n' {
				length -= 1
			}

			// windows style \r\n.  Don't know why I even bothered checking ...
			if char == '\r' {
				length -= 1
			}

			break
		}

		text.Value = text.Value[:length]
	}

	if lexer.lookAhead[0].TrimTrailingWhitespaces() &&
		lexer.lookAhead[1].Id() == TextToken {

		text := lexer.lookAhead[1].(*Atom)
		start := 0

		for start < len(text.Value) {
			char := text.Value[start]
			if char == ' ' || char == '\t' || char == '\r' {
				start += 1
				continue
			}

			if char == '\n' {
				start += 1
			}

			break
		}

		text.Value = text.Value[start:]
	}
}

func (lexer *bodyLexer) Next() (Token, error) {
	lexer.fillTokensAndMaybeTrimWhitespaces()

	if len(lexer.lookAhead) > 0 {
		ret := lexer.lookAhead[0]
		lexer.lookAhead = lexer.lookAhead[1:]
		return ret, nil
	}

	if lexer.lookAheadErr == nil {
		panic("Programming error")
	}

	return nil, lexer.lookAheadErr
}
