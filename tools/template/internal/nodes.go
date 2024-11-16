package template

import (
	"github.com/pattyshack/gt/parseutil"
)

// Note: in case of text, we need to modify the value and escape ` and /
// correctly

// in lexer, we can enforce scoping character.  If we prevent unpaired ), we
// can ensure that ($(stmt.Value)) always is an valid expression

type Statement interface {
	IsStatement()

	parseutil.Token[SymbolId]
}

type TToken struct {
	parseutil.TokenValue[SymbolId]

	// When true, and the previous statement is text, remove the whitespaces
	// in the text that are adjacent to this statement, potentially up to and
	// including the previous line's newline character.
	trimLeadingWhitespaces bool

	// When true, and the next statement is text, remove the whitespaces
	// in the text taht are adjacent to this statement, potentially up to and
	// including the current line's newline character.
	trimTrailingWhitespaces bool
}

func NewTToken(
	id SymbolId,
	pos parseutil.StartEndPos,
	trimLeading bool,
	trimTrailing bool) *TToken {

	return &TToken{
		TokenValue: parseutil.TokenValue[SymbolId]{
			SymbolId:    id,
			StartEndPos: pos,
			Value:       "",
		},
		trimLeadingWhitespaces:  trimLeading,
		trimTrailingWhitespaces: trimTrailing,
	}
}

func (token *TToken) TrimLeadingWhitespaces() bool {
	return token.trimLeadingWhitespaces
}

func (token *TToken) TrimTrailingWhitespaces() bool {
	return token.trimTrailingWhitespaces
}

type Value struct {
	*TToken

	Value string
}

func NewValue(
	id SymbolId,
	pos parseutil.StartEndPos,
	val string,
	trimLeading bool,
	trimTrailing bool) *Value {

	return &Value{NewTToken(id, pos, trimLeading, trimTrailing), val}
}

type Atom struct {
	*TToken

	Value string
}

func NewAtom(
	id SymbolId,
	pos parseutil.StartEndPos,
	val string,
	trimLeading bool,
	trimTrailing bool) *Atom {

	return &Atom{NewTToken(id, pos, trimLeading, trimTrailing), val}
}

func (Atom) IsStatement() {}

type Branch struct {
	Predicate *Value
	Body      []Statement
}

type For struct {
	parseutil.StartEndPos

	Branch
}

func (For) IsStatement() {}

func (For) Id() SymbolId { return ForType }

type Switch struct {
	parseutil.StartEndPos

	Switch  *Value
	Cases   []*Branch
	Default *Branch
}

func (Switch) IsStatement() {}

func (Switch) Id() SymbolId { return SwitchType }

type If struct {
	parseutil.StartEndPos

	If      Branch
	ElseIfs []*Branch
	Else    *Branch
}

func (If) IsStatement() {}

func (If) Id() SymbolId { return IfType }

type Argument struct {
	Name string
	Type string
}

type TemplateDeclaration struct {
	parseutil.TokenValue[SymbolId]

	TemplateName string
	Arguments    []Argument
}

func NewTemplateDeclaration(
	pos parseutil.StartEndPos,
	name string,
	args []Argument) *TemplateDeclaration {

	return &TemplateDeclaration{
		TokenValue: parseutil.TokenValue[SymbolId]{
			SymbolId:    TemplateDeclToken,
			StartEndPos: pos,
			Value:       "",
		},
		TemplateName: name,
		Arguments:    args,
	}
}

type File struct {
	PackageName string

	Imports string

	*TemplateDeclaration

	Body []Statement
}
