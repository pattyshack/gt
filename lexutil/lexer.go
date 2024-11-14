package lexutil

type Lexer[T any] interface {
	Next() (T, error)
	CurrentLocation() Location
}

type LexerReader[T any] struct {
	Lexer[T]
}

func (lexer *LexerReader[T]) Read(buffer []T) (int, error) {
	if len(buffer) == 0 {
		return 0, nil
	}

	t, err := lexer.Next()
	if err != nil {
		return 0, err
	}

	buffer[0] = t
	return 1, nil
}

func NewLexerReader[T any](lexer Lexer[T]) *LexerReader[T] {
	return &LexerReader[T]{
		Lexer: lexer,
	}
}

// Discard all trimSymbol tokens.
type TrimLexer[SymbolId comparable] struct {
	base     Lexer[Token[SymbolId]]
	buffered *BufferedReader[Token[SymbolId]]

	trimSymbol SymbolId
}

func NewTrimLexer[SymbolId comparable](
	base Lexer[Token[SymbolId]],
	trimSymbol SymbolId,
) Lexer[Token[SymbolId]] {
	return &TrimLexer[SymbolId]{
		base:       base,
		buffered:   NewBufferedReader(NewLexerReader(base), 10),
		trimSymbol: trimSymbol,
	}
}

func (lexer *TrimLexer[SymbolId]) CurrentLocation() Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *TrimLexer[SymbolId]) Next() (Token[SymbolId], error) {
	for {
		token, err := lexer.buffered.Next()
		if err != nil {
			return nil, err
		}

		if token.Id() == lexer.trimSymbol {
			continue
		}

		return token, nil
	}
}

// Merge all adjacent mergeSymbol *TokenCount.
type MergeTokenCountLexer[SymbolId comparable] struct {
	base     Lexer[Token[SymbolId]]
	buffered *BufferedReader[Token[SymbolId]]

	mergeSymbol SymbolId
}

func NewMergeTokenCountLexer[SymbolId comparable](
	base Lexer[Token[SymbolId]],
	mergeSymbol SymbolId,
) Lexer[Token[SymbolId]] {
	return &MergeTokenCountLexer[SymbolId]{
		base:        base,
		buffered:    NewBufferedReader(NewLexerReader(base), 10),
		mergeSymbol: mergeSymbol,
	}
}

func (lexer *MergeTokenCountLexer[SymbolId]) CurrentLocation() Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *MergeTokenCountLexer[SymbolId]) Next() (Token[SymbolId], error) {
	token, err := lexer.buffered.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() != lexer.mergeSymbol {
		return token, nil
	}

	tokenCount := token.(*TokenCount[SymbolId])

	for {
		peeked, err := lexer.buffered.Peek(1)
		if err != nil || len(peeked) == 0 {
			break
		}

		if peeked[0].Id() != lexer.mergeSymbol {
			break
		}

		next := peeked[0].(*TokenCount[SymbolId])

		tokenCount.Count += next.Count
		tokenCount.EndPos = next.EndPos

		_, err = lexer.buffered.Discard(1)
		if err != nil {
			panic("should never happen")
		}
	}

	return token, nil
}
