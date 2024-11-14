package parser

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pattyshack/gt/lexutil"
	"github.com/pattyshack/gt/stringutil"
)

const (
	initialPeekSize = 64
)

type rawLexer struct {
	reader lexutil.BufferedByteLocationReader

	lexutil.ConstantSymbols[LRSymbolId]

	internPool *stringutil.InternPool
}

func newRawLexer(filename string, reader io.Reader) *rawLexer {
	markersAndSymbols := map[string]LRSymbolId{
		TokenMarker: LRTokenToken,
		TypeMarker:  LRTypeToken,
		"%start":    LRStartToken,
		"<":         '<',
		">":         '>',
		"|":         '|',
		";":         ';',
		"=":         '=',

		"%%": LRSectionMarkerToken,

		"->": Arrow,
		":":  ':',
	}

	pool := stringutil.NewInternPool()
	return &rawLexer{
		reader: lexutil.NewBufferedByteLocationReader(
			filename,
			reader,
			1024*1024),
		ConstantSymbols: lexutil.NewConstantSymbols(markersAndSymbols, pool),
		internPool:      pool,
	}
}

func (lexer *rawLexer) CurrentLocation() lexutil.Location {
	return lexer.reader.Location
}

func (lexer *rawLexer) Next() (LRToken, error) {
	err := lexutil.StripLeadingWhitespacesAndComments(lexer.reader)
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
	token, err := lexer.MaybeTokenizeSymbol(lexer.reader)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, nil
	}

	return *token, nil
}

func (lexer *rawLexer) maybeTokenizeCharacter() (LRToken, error) {
	token, errMsg, err := lexutil.MaybeTokenizeRuneLiteral(
		lexer.reader,
		8,
		lexer.internPool,
		LRCharacterToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, nil
	}

	if errMsg != "" {
		return nil, lexutil.NewLocationError(token.StartPos, errMsg)
	}

	return token, nil
}

func (lexer *rawLexer) maybeTokenizeIdentifier() (LRToken, error) {
	// NOTE: cannot return this directly since nil token gets wrapped by LRToken,
	// which then becomes "non-nil" ...
	token, err := lexutil.MaybeTokenizeIdentifier(
		lexer.reader,
		initialPeekSize,
		lexer.internPool,
		LRIdentifierToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, nil
	}

	return token, nil
}

func (lexer *rawLexer) maybeTokenizeSectionContent() (LRToken, error) {
	peek, _ := lexer.reader.Peek(1)
	if string(peek) != "{" {
		return nil, nil
	}

	token := &Token{
		SymbolId: LRSectionContentToken,
		Value:    "",
	}
	token.StartPos = lexer.reader.Location

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

type Lexer struct {
	base     *rawLexer
	buffered *lexutil.BufferedReader[LRToken]
}

// This merges LRIdentifierSymbol Arrow token pairs into a single RULE_DEF
// token and LRIdentifierSymbol Colon token pairs into a single LABEL token.
func (lexer *Lexer) Next() (LRToken, error) {
	tokens, err := lexer.buffered.Peek(5)
	if len(tokens) < 1 {
		return nil, err
	}

	curr := tokens[0]
	if len(tokens) < 2 || curr.Id() != LRIdentifierToken {
		lexer.buffered.Discard(1)
		return curr, nil
	}

	next := tokens[1]
	if next.Id() == Arrow {
		lexer.buffered.Discard(2)

		return &RuleDef{
			Name: curr.(*Token),
		}, nil
	} else if next.Id() == ':' {
		curr.(*Token).SymbolId = LRLabelToken

		lexer.buffered.Discard(2)
		return curr, nil
	} else if next.Id() == '<' &&
		len(tokens) == 5 &&
		tokens[2].Id() == LRIdentifierToken &&
		tokens[3].Id() == '>' &&
		tokens[4].Id() == Arrow {

		def := &RuleDef{
			Name:      curr.(*Token),
			ValueType: tokens[2].(*Token),
		}

		lexer.buffered.Discard(5)
		return def, nil
	}

	lexer.buffered.Discard(1)
	return curr, nil
}

func (lexer *Lexer) CurrentLocation() lexutil.Location {
	tokens, err := lexer.buffered.Peek(1)
	if err != nil || len(tokens) == 0 {
		return lexer.base.reader.Location
	}

	return tokens[0].Loc()
}

func NewLexer(filename string, reader io.Reader) LRLexer {
	base := newRawLexer(filename, reader)
	return &Lexer{
		base:     base,
		buffered: lexutil.NewBufferedReader(lexutil.NewLexerReader(base), 1000),
	}
}
