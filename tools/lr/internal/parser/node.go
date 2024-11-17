package parser

import (
	"strings"

	"github.com/pattyshack/gt/parseutil"
)

type Definition interface {
	Loc() parseutil.Location
	String() string
}

type Token = parseutil.TokenValue[LRSymbolId]

type StartDeclaration struct {
	parseutil.StartEndPos

	Ids []*Token
}

func NewStartDeclaration(
	start parseutil.TokenValue[LRSymbolId],
	ids []*Token) *StartDeclaration {

	pos := start.StartEndPos
	if len(ids) > 0 {
		pos.EndPos = ids[len(ids)-1].EndPos
	}

	return &StartDeclaration{
		StartEndPos: pos,
		Ids:         ids,
	}
}

func (sd *StartDeclaration) String() string {
	names := []string{}
	for _, id := range sd.Ids {
		names = append(names, id.Value)
	}
	return "%start " + strings.Join(names, " ")
}

type TermDeclaration struct {
	parseutil.StartEndPos

	TermType parseutil.TokenValue[LRSymbolId]

	IsTerminal bool

	ValueType *Token

	Terms []*Token
}

func NewTermDeclaration(
	termType parseutil.TokenValue[LRSymbolId],
	valueType *Token,
	terms []*Token) *TermDeclaration {

	pos := termType.StartEndPos
	if valueType != nil {
		pos.EndPos = valueType.EndPos
	}
	if len(terms) > 0 {
		pos.EndPos = terms[len(terms)-1].EndPos
	}

	return &TermDeclaration{
		StartEndPos: pos,
		TermType:    termType,
		IsTerminal:  termType.Id() == LRTokenToken,
		ValueType:   valueType,
		Terms:       terms,
	}
}

func (td *TermDeclaration) String() string {
	result := TypeMarker
	if td.IsTerminal {
		result = TokenMarker
	}

	if td.ValueType != nil {
		result += " <" + td.ValueType.Value + ">"
	}

	for _, term := range td.Terms {
		result += " " + term.Value
	}

	return result
}

type Clause struct {
	Label *Token // optional
	Body  []*Token

	Passthrough bool

	// set by NewRule
	parseutil.Location
	Parent *Rule

	SortId int
}

func NewClause(label *Token, body []*Token, passthrough bool) *Clause {
	return &Clause{
		Label:       label,
		Body:        body,
		Passthrough: passthrough,
	}
}

func (clause *Clause) String() string {
	result := "(none):"
	if clause.Label != nil {
		result = clause.Label.Value + ":"
	}

	for _, item := range clause.Body {
		result += " " + item.Value
	}

	return result
}

type RuleDef struct {
	parseutil.StartEndPos

	Name      *Token
	ValueType *Token
}

func (RuleDef) Id() LRSymbolId {
	return LRRuleDefToken
}

func (def *RuleDef) Loc() parseutil.Location {
	return def.Name.Loc()
}

func (def *RuleDef) End() parseutil.Location {
	return def.ValueType.Loc()
}

type Rule struct {
	*RuleDef
	Clauses []*Clause
}

func NewRule(name *RuleDef, clauses []*Clause) *Rule {
	rule := &Rule{
		RuleDef: name,
		Clauses: clauses,
	}

	for _, clause := range clauses {
		loc := name.Loc()
		if clause.Label != nil {
			loc = clause.Label.StartPos
		} else if len(clause.Body) > 0 {
			loc = clause.Body[0].StartPos
		}

		clause.Location = loc
		clause.Parent = rule
	}

	return rule
}

func (r *Rule) Loc() parseutil.Location {
	return r.Name.StartPos
}

func (r *Rule) String() string {
	result := r.Name.Value + " ->"
	for i, clause := range r.Clauses {
		result += "\n    " + clause.String()
		if i < len(r.Clauses)-1 {
			result += " |"
		}
	}

	return result
}

type AdditionalSection struct {
	Name    *Token
	Content *Token
}

func NewAdditionalSection(name *Token, content *Token) *AdditionalSection {
	return &AdditionalSection{
		Name:    name,
		Content: content,
	}
}

type Grammar struct {
	Definitions []Definition

	AdditionalSections []*AdditionalSection
}

func NewGrammar(
	defs []Definition,
	additionalSections []*AdditionalSection) *Grammar {

	return &Grammar{defs, additionalSections}
}
