package parseutil

import (
	"fmt"
)

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

type Token[SymbolId any] interface {
	Locatable
	Id() SymbolId
}

type TokenValue[SymbolId any] struct {
	SymbolId SymbolId
	StartEndPos

	Value string

	SubType LiteralSubType // Only set by certain literal tokens
}

func (tv TokenValue[SymbolId]) Id() SymbolId {
	return tv.SymbolId
}

func (tv TokenValue[SymbolId]) Val() string {
	return tv.Value
}

func (tv TokenValue[SymbolId]) String() string {
	return fmt.Sprintf("%v (%s)", tv.SymbolId, tv.Value)
}

type TokenCount[SymbolId any] struct {
	SymbolId SymbolId
	StartEndPos

	Count int
}

func (tc TokenCount[SymbolId]) Id() SymbolId {
	return tc.SymbolId
}

func (tc TokenCount[SymbolId]) String() string {
	return fmt.Sprintf("%v (%d)", tc.SymbolId, tc.Count)
}

type TokenError[SymbolId any] struct {
	SymbolId SymbolId
	StartEndPos

	Error error
}

func (te TokenError[SymbolId]) Id() SymbolId {
	return te.SymbolId
}

func (te TokenError[SymbolId]) String() string {
	return fmt.Sprintf("%v (%s): %s", te.SymbolId, te.Error, te.StartPos)
}
