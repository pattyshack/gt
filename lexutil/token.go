package lexutil

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

type StartEndPos struct {
	StartPos Location
	EndPos   Location
}

func NewStartEndPos(start Location, end Location) StartEndPos {
	return StartEndPos{
		StartPos: start,
		EndPos:   end,
	}
}

func (sep StartEndPos) StartEnd() StartEndPos {
	return sep
}

func (sep StartEndPos) Loc() Location {
	return sep.StartPos
}

func (sep StartEndPos) End() Location {
	return sep.EndPos
}

type Token[SymbolId any] interface {
	Id() SymbolId
	Loc() Location
	End() Location
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
	return fmt.Sprintf("%s: %v (%s)", tv.StartPos, tv.SymbolId, tv.Value)
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
	return fmt.Sprintf("%s: %v (%d)", tc.StartPos, tc.SymbolId, tc.Count)
}
