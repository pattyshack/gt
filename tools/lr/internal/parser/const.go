package parser

const (
	TokenMarker = "%token"
	TypeMarker  = "%type"

	// intermediate tokens used by the lexer, not directly consumed by the
	// parser.
	Arrow = -2
)
