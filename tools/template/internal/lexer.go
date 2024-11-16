package template

import (
	"io"
	"io/ioutil"

	"github.com/pattyshack/gt/parseutil"
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
	Next() (parseutil.Token[SymbolId], error)
}

type LexerImpl struct {
	reader parseutil.BufferedByteLocationReader

	currentLexer

	internPool *stringutil.InternPool
}

func NewLexer(
	filename string,
	input io.Reader,
) (
	parseutil.Lexer[parseutil.Token[SymbolId]],
	error,
) {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	content = stripHeaderComments(content)

	reader := parseutil.NewBufferedByteLocationReaderFromSlice(
		filename,
		content)
	internPool := stringutil.NewInternPool()

	return &LexerImpl{
		reader: reader,
		currentLexer: &headerLexer{
			reader:        reader,
			internPool:    internPool,
			sectionMarker: parseutil.NewConstantSymbols(sectionMarker, internPool),
			importMarker:  parseutil.NewConstantSymbols(importMarker, internPool),
			templateDeclMarker: parseutil.NewConstantSymbols(
				templateDeclMarker,
				internPool),
		},
		internPool: internPool,
	}, nil
}

func (lexer *LexerImpl) CurrentLocation() parseutil.Location {
	return lexer.reader.Location
}

func (lexer *LexerImpl) Next() (parseutil.Token[SymbolId], error) {
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
	reader     parseutil.BufferedByteLocationReader
	internPool *stringutil.InternPool

	sectionMarker      parseutil.ConstantSymbols[struct{}]
	importMarker       parseutil.ConstantSymbols[struct{}]
	templateDeclMarker parseutil.ConstantSymbols[struct{}]
}

func (lexer *headerLexer) Next() (parseutil.Token[SymbolId], error) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	token, err := parseutil.MaybeTokenizeIdentifier(
		lexer.reader,
		initialPeekSize,
		lexer.internPool,
		struct{}{})
	if err != nil {
		return nil, err
	}

	value := ""
	if token != nil {
		value = token.Value
	}

	switch value {
	case "package":
		return lexer.tokenizePackage(token.StartEndPos)
	case "import":
		return lexer.tokenizeImport(token.StartEndPos)
	case "template":
		return lexer.tokenizeTemplateDecl(token.StartEndPos)
	case "":
		// try to tokenize symbol below
	default:
		return nil, parseutil.NewLocationError(
			token.StartPos,
			"Unexpected IDENTIFIER %s",
			value)
	}

	token, err = lexer.sectionMarker.MaybeTokenizeSymbol(lexer.reader)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, parseutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	return parseutil.TokenValue[SymbolId]{
		SymbolId:    SectionMarkerToken,
		StartEndPos: token.StartEndPos,
	}, nil
}

func (lexer *headerLexer) tokenizePackage(
	pkgPos parseutil.StartEndPos,
) (
	parseutil.Token[SymbolId],
	error,
) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	token, err := parseutil.MaybeTokenizeIdentifier(
		lexer.reader,
		initialPeekSize,
		lexer.internPool,
		PackageToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, parseutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	pkgPos.EndPos = token.EndPos
	return NewValue(PackageToken, pkgPos, token.Value, false, false), nil
}

func (lexer *headerLexer) tokenizeImport(
	importPos parseutil.StartEndPos,
) (
	parseutil.Token[SymbolId],
	error,
) {
	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	token, err := lexer.importMarker.MaybeTokenizeSymbol(
		lexer.reader)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, parseutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	value, _, err := readDirective(lexer.reader, 0, ")")
	if err != nil {
		return nil, err
	}

	value = value[:len(value)-1]

	importPos.EndPos = lexer.reader.Location
	return NewValue(ImportToken, importPos, string(value), false, false), nil
}

func (lexer *headerLexer) tokenizeTemplateDecl(
	declPos parseutil.StartEndPos,
) (
	parseutil.Token[SymbolId],
	error,
) {

	err := parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	template, err := parseutil.MaybeTokenizeIdentifier(
		lexer.reader,
		initialPeekSize,
		lexer.internPool,
		"")
	if err != nil {
		return nil, err
	}

	if template == nil {
		return nil, parseutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	err = parseutil.StripLeadingWhitespaces(lexer.reader)
	if err != nil {
		return nil, err
	}

	lcurl, err := lexer.templateDeclMarker.MaybeTokenizeSymbol(
		lexer.reader)
	if err != nil {
		return nil, err
	}

	if lcurl == nil {
		return nil, parseutil.NewLocationError(
			lexer.reader.Location,
			"Unexpected character")
	}

	body, loc, err := readDirective(lexer.reader, 0, "}")
	if err != nil {
		return nil, err
	}

	body = body[:len(body)-1]

	declReader := parseutil.NewBufferedByteLocationReaderFromSlice("", body)
	declReader.Location = loc

	args := []Argument{}
	for {
		err := parseutil.StripLeadingWhitespaces(declReader)
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
		lineReader := parseutil.NewBufferedByteLocationReaderFromSlice("", line)
		lineReader.Location = loc

		argName, err := parseutil.MaybeTokenizeIdentifier(
			lineReader,
			initialPeekSize,
			lexer.internPool,
			"")
		if err != nil {
			return nil, err
		}

		if argName == nil {
			return nil, parseutil.NewLocationError(
				lineReader.Location,
				"Expecting argument name")
		}

		err = parseutil.StripLeadingWhitespaces(lineReader)
		if err != nil && err != io.EOF {
			return nil, err
		}

		typeName, err := ioutil.ReadAll(lineReader)
		if err != nil {
			return nil, err
		}

		if len(typeName) == 0 {
			return nil, parseutil.NewLocationError(
				lineReader.Location,
				"Expecting argument type")
		}

		args = append(args, Argument{string(argName.Value), string(typeName)})
	}

	declPos.EndPos = lexer.reader.Location
	return NewTemplateDeclaration(declPos, template.Value, args), nil
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

func (lexer *bodyLexer) Next() (parseutil.Token[SymbolId], error) {
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
