package parser

import (
	"bytes"
	"fmt"
	"io"
	"sort"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
	"github.com/pattyshack/gt/tools/lr/parseutil"
)

var (
	keywordsAndSymbols = map[string]LRSymbolId{
		TokenKeyword: LRTokenToken,
		TypeKeyword:  LRTypeToken,
		"%start":     LRStartToken,
		"<":          '<',
		">":          '>',
		"|":          '|',
		";":          ';',

		"%%": LRSectionMarkerToken,

		"->": Arrow,
		":":  ':',
	}
)

type rawLexer struct {
	reader lexutil.BufferedByteLocationReader

	sorted parseutil.Symbols

	internPool *stringutil.InternPool
}

func newRawLexer(filename string, reader io.Reader) *rawLexer {
	sorted := parseutil.Symbols{}
	for val, id := range keywordsAndSymbols {
		sorted = append(sorted, parseutil.Symbol{val, int(id)})
	}
	sort.Sort(sorted)

	return &rawLexer{
		reader: lexutil.NewBufferedByteLocationReader(
			filename,
			reader,
			1024*1024),
		sorted:     sorted,
		internPool: stringutil.NewInternPool(),
	}
}

func (lexer *rawLexer) Next() (LRToken, error) {
	err := parseutil.StripLeadingWhitespacesAndComments(lexer.reader)
	if err != nil {
		return nil, err
	}

	_, err = lexer.reader.Peek(1)
	if err != nil {
		return nil, err
	}

	var token LRToken
	token, err = lexer.maybeTokenizeKeywordOrSymbol()
	if token != nil || err != nil {
		return token, err
	}
	token, err = lexer.maybeTokenizeCharacter()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeIdentifier()
	if token != nil || err != nil {
		return token, err
	}

	token, err = lexer.maybeTokenizeSectionContent()
	if token != nil || err != nil {
		return token, err
	}

	return nil, fmt.Errorf("Unexpected character at %s", lexer.reader.Location)
}

func (lexer *rawLexer) maybeTokenizeKeywordOrSymbol() (LRToken, error) {
	symbol, loc, err := parseutil.MaybeTokenizeSymbol(
		lexer.reader,
		lexer.sorted)
	if err != nil {
		return nil, err
	}

	if symbol == nil {
		return nil, nil
	}

	return &LRGenericSymbol{
		LRSymbolId: LRSymbolId(symbol.Id),
		LRLocation: LRLocation(loc),
	}, nil
}

func (lexer *rawLexer) maybeTokenizeCharacter() (LRToken, error) {
	value, loc, err := parseutil.MaybeTokenizeCharacter(lexer.reader)
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	return &Token{
		LRLocation: LRLocation(loc),
		LRSymbolId: LRCharacterToken,
		Value:      value,
	}, nil
}

func (lexer *rawLexer) maybeTokenizeIdentifier() (LRToken, error) {
	value, loc, err := parseutil.MaybeTokenizeIdentifier(
		lexer.reader,
		lexer.internPool)
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, nil
	}

	return &Token{
		LRLocation: LRLocation(loc),
		LRSymbolId: LRIdentifierToken,
		Value:      value,
	}, nil
}

func (lexer *rawLexer) maybeTokenizeSectionContent() (LRToken, error) {
	peek, _ := lexer.reader.Peek(1)
	if string(peek) != "{" {
		return nil, nil
	}

	token := &Token{
		LRLocation: LRLocation(lexer.reader.Location),
		LRSymbolId: LRSectionContentToken,
		Value:      "",
	}

	n, err := lexer.reader.Read(peek)
	if n != 1 || err != nil {
		panic(err) // should never happen
	}

	buffer := bytes.NewBuffer(nil)

	singleByte := [1]byte{}
	for {
		peek, err = lexer.reader.Peek(3)
		if err != nil {
			return nil, err
		}

		if string(peek) == "}%%" {
			n, err := lexer.reader.Read(peek)
			if n != 3 || err != nil {
				panic(err) // should never happen
			}

			break
		}

		numRead, err := lexer.reader.Read(singleByte[:])
		if err != nil || numRead == 0 {
			panic(err)
		}

		buffer.WriteByte(singleByte[0])
	}

	token.Value = string(buffer.Bytes())
	return token, nil
}

// This merges LRIdentifierSymbol Arrow token pairs into a single RULE_DEF token and
// LRIdentifierSymbol Colon token pairs into a single LABEL token.
type tokenPairLexer struct {
	base *rawLexer

	nextToken LRToken
	nextErr   error
}

func (lexer *tokenPairLexer) Next() (LRToken, error) {
	if lexer.nextErr != nil {
		err := lexer.nextErr
		lexer.nextErr = nil

		return nil, err
	}

	curr := lexer.nextToken
	lexer.nextToken = nil

	var err error
	if curr == nil {
		curr, err = lexer.base.Next()
		if err != nil {
			return nil, err
		}
	}

	if curr.Id() != LRIdentifierToken {
		return curr, nil
	}

	next, err := lexer.base.Next()
	if err != nil {
		lexer.nextErr = err
		return curr, nil
	}

	if next.Id() == Arrow {
		curr.(*Token).LRSymbolId = LRRuleDefToken
		return curr, nil
	}

	if next.Id() == ':' {
		curr.(*Token).LRSymbolId = LRLabelToken
		return curr, nil
	}

	lexer.nextToken = next
	return curr, nil
}

func (lexer *tokenPairLexer) CurrentLocation() LRLocation {
	if lexer.nextToken != nil {
		return lexer.nextToken.Loc()
	}

	return LRLocation(lexer.base.reader.Location)
}

func NewLexer(filename string, reader io.Reader) LRLexer {
	return &tokenPairLexer{
		base:      newRawLexer(filename, reader),
		nextToken: nil,
		nextErr:   nil,
	}
}
