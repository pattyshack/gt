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
type TrimTokenLexer[SymbolId comparable] struct {
	base Lexer[Token[SymbolId]]

	trimSymbol SymbolId
}

func NewTrimTokenLexer[SymbolId comparable](
	base Lexer[Token[SymbolId]],
	trimSymbol SymbolId,
) Lexer[Token[SymbolId]] {
	return &TrimTokenLexer[SymbolId]{
		base:       base,
		trimSymbol: trimSymbol,
	}
}

func (lexer *TrimTokenLexer[SymbolId]) CurrentLocation() Location {
	return lexer.base.CurrentLocation()
}

func (lexer *TrimTokenLexer[SymbolId]) Next() (Token[SymbolId], error) {
	for {
		token, err := lexer.base.Next()
		if err != nil {
			return nil, err
		}

		if token.Id() == lexer.trimSymbol {
			continue
		}

		return token, nil
	}
}

type MergeFunc[SymbolId comparable] func(
	Token[SymbolId],
	*BufferedReader[Token[SymbolId]],
) Token[SymbolId]

// Merge all adjacent mergeSymbol Token using the provided merge function.
// The merge function must discard all consumed tokens.
type MergeTokenLexer[SymbolId comparable] struct {
	base     Lexer[Token[SymbolId]]
	buffered *BufferedReader[Token[SymbolId]]

	mergeSymbol SymbolId
	merge       MergeFunc[SymbolId]
}

func NewMergeTokenLexer[SymbolId comparable](
	base Lexer[Token[SymbolId]],
	mergeSymbol SymbolId,
	merge MergeFunc[SymbolId],
	bufferSize int,
) Lexer[Token[SymbolId]] {
	return &MergeTokenLexer[SymbolId]{
		base:        base,
		buffered:    NewBufferedReader(NewLexerReader(base), bufferSize),
		mergeSymbol: mergeSymbol,
		merge:       merge,
	}
}

func (lexer *MergeTokenLexer[SymbolId]) CurrentLocation() Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *MergeTokenLexer[SymbolId]) Next() (Token[SymbolId], error) {
	token, err := lexer.buffered.Next()
	if err != nil {
		return nil, err
	}

	if token.Id() != lexer.mergeSymbol {
		return token, nil
	}

	return lexer.merge(token, lexer.buffered), nil
}

// Merge all adjacent mergeSymbol TokenCount.
func NewMergeTokenCountLexer[SymbolId comparable](
	base Lexer[Token[SymbolId]],
	mergeSymbol SymbolId,
) Lexer[Token[SymbolId]] {
	return NewMergeTokenLexer(
		base,
		mergeSymbol,
		func(
			token Token[SymbolId],
			buffered *BufferedReader[Token[SymbolId]],
		) Token[SymbolId] {
			count := token.(*TokenCount[SymbolId])
			for {
				peeked, err := buffered.Peek(1)
				if err != nil || len(peeked) == 0 || peeked[0].Id() != mergeSymbol {
					return token
				}

				next := peeked[0].(*TokenCount[SymbolId])
				count.Count += next.Count
				count.EndPos = next.EndPos

				_, err = buffered.Discard(1)
				if err != nil {
					panic("should not happen")
				}
			}
		},
		10)
}
