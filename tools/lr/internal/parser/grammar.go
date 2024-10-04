// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	lexutil "github.com/pattyshack/gt/lexutil"
	io "io"
)

type LRSymbolId int

const (
	// char token '<' = LRSymbolId(60)
	// char token '>' = LRSymbolId(62)
	// char token '|' = LRSymbolId(124)
	// char token ';' = LRSymbolId(59)
	// char token '=' = LRSymbolId(61)
	LRTokenToken          = LRSymbolId(256)
	LRTypeToken           = LRSymbolId(257)
	LRStartToken          = LRSymbolId(258)
	LRRuleDefToken        = LRSymbolId(259)
	LRLabelToken          = LRSymbolId(260)
	LRSectionMarkerToken  = LRSymbolId(261)
	LRCharacterToken      = LRSymbolId(262)
	LRIdentifierToken     = LRSymbolId(263)
	LRSectionContentToken = LRSymbolId(264)
)

type LRLocation = lexutil.Location

type LRToken interface {
	Id() LRSymbolId
	Loc() LRLocation
}

type LRGenericSymbol struct {
	LRSymbolId
	StartPos LRLocation
}

func (t LRGenericSymbol) Id() LRSymbolId { return t.LRSymbolId }

func (t LRGenericSymbol) Loc() LRLocation { return t.StartPos }

type LRLexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return LRGenericSymbol
	Next() (LRToken, error)

	CurrentLocation() LRLocation
}

type LRGrammarReducer interface {
	// 26:20: grammar -> ...
	ToGrammar(Defs_ []Definition, AdditionalSections_ []*AdditionalSection) (*Grammar, error)
}

type LRAdditionalSectionsReducer interface {
	// 29:4: additional_sections -> add: ...
	AddToAdditionalSections(AdditionalSections_ []*AdditionalSection, AdditionalSection_ *AdditionalSection) ([]*AdditionalSection, error)

	// 30:4: additional_sections -> nil: ...
	NilToAdditionalSections() ([]*AdditionalSection, error)
}

type LRAdditionalSectionReducer interface {
	// 32:41: additional_section -> ...
	ToAdditionalSection(SectionMarker_ LRGenericSymbol, Identifier_ *Token, SectionContent_ *Token) (*AdditionalSection, error)
}

type LRDefsReducer interface {
	// 35:4: defs -> add: ...
	AddToDefs(Defs_ []Definition, Def_ Definition) ([]Definition, error)

	// 36:4: defs -> add_explicit: ...
	AddExplicitToDefs(Defs_ []Definition, Def_ Definition, char LRGenericSymbol) ([]Definition, error)

	// 37:4: defs -> def: ...
	DefToDefs(Def_ Definition) ([]Definition, error)

	// 38:4: defs -> explicit_def: ...
	ExplicitDefToDefs(Def_ Definition, char LRGenericSymbol) ([]Definition, error)
}

type LRDefReducer interface {
	// 42:4: def -> term_decl: ...
	TermDeclToDef(Rword_ LRGenericSymbol, char LRGenericSymbol, Identifier_ *Token, char2 LRGenericSymbol, NonemptyIdOrCharList_ []*Token) (Definition, error)

	// 43:4: def -> untyped_term_decl: ...
	UntypedTermDeclToDef(Rword_ LRGenericSymbol, NonemptyIdOrCharList_ []*Token) (Definition, error)

	// 45:4: def -> start_decl: ...
	StartDeclToDef(Start_ LRGenericSymbol, NonemptyIdentList_ []*Token) (Definition, error)

	// 46:4: def -> rule: ...
	RuleToDef(Rule_ *Rule) (Definition, error)
}

type LRNonemptyIdentListReducer interface {
	// 53:4: nonempty_ident_list -> add: ...
	AddToNonemptyIdentList(NonemptyIdentList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 54:4: nonempty_ident_list -> ident: ...
	IdentToNonemptyIdentList(Identifier_ *Token) ([]*Token, error)
}

type LRNonemptyIdOrCharListReducer interface {
	// 57:4: nonempty_id_or_char_list -> add_id: ...
	AddIdToNonemptyIdOrCharList(NonemptyIdOrCharList_ []*Token, Identifier_ *Token) ([]*Token, error)

	// 58:4: nonempty_id_or_char_list -> add_char: ...
	AddCharToNonemptyIdOrCharList(NonemptyIdOrCharList_ []*Token, Character_ *Token) ([]*Token, error)

	// 59:4: nonempty_id_or_char_list -> id: ...
	IdToNonemptyIdOrCharList(Identifier_ *Token) ([]*Token, error)

	// 60:4: nonempty_id_or_char_list -> char: ...
	CharToNonemptyIdOrCharList(Character_ *Token) ([]*Token, error)
}

type LRIdOrCharListReducer interface {
	// 63:4: id_or_char_list -> list: ...
	ListToIdOrCharList(NonemptyIdOrCharList_ []*Token) ([]*Token, error)

	// 64:4: id_or_char_list -> nil: ...
	NilToIdOrCharList() ([]*Token, error)
}

type LRRuleReducer interface {
	// 66:14: rule -> ...
	ToRule(RuleDef_ *RuleDef, Clauses_ []*Clause) (*Rule, error)
}

type LRClauseReducer interface {
	// 69:2: clause -> passthrough_id: ...
	PassthroughIdToClause(char LRGenericSymbol, Identifier_ *Token) (*Clause, error)

	// 70:2: clause -> passthrough_char: ...
	PassthroughCharToClause(char LRGenericSymbol, Character_ *Token) (*Clause, error)

	// 71:2: clause -> unlabeled: ...
	UnlabeledToClause(IdOrCharList_ []*Token) (*Clause, error)

	// 72:2: clause -> labeled: ...
	LabeledToClause(Label_ *Token, IdOrCharList_ []*Token) (*Clause, error)
}

type LRClausesReducer interface {
	// 75:4: clauses -> add: ...
	AddToClauses(Clauses_ []*Clause, char LRGenericSymbol, Clause_ *Clause) ([]*Clause, error)

	// 76:4: clauses -> clause: ...
	ClauseToClauses(Clause_ *Clause) ([]*Clause, error)
}

type LRReducer interface {
	LRGrammarReducer
	LRAdditionalSectionsReducer
	LRAdditionalSectionReducer
	LRDefsReducer
	LRDefReducer
	LRNonemptyIdentListReducer
	LRNonemptyIdOrCharListReducer
	LRIdOrCharListReducer
	LRRuleReducer
	LRClauseReducer
	LRClausesReducer
}

type LRParseErrorHandler interface {
	Error(nextToken LRToken, parseStack _LRStack) error
}

type LRDefaultParseErrorHandler struct{}

func (LRDefaultParseErrorHandler) Error(nextToken LRToken, stack _LRStack) error {
	return lexutil.NewLocationError(
		nextToken.Loc(),
		"syntax error: unexpected symbol %s. expecting %v",
		nextToken.Id(),
		LRExpectedTerminals(stack[len(stack)-1].StateId))
}

func LRExpectedTerminals(id _LRStateId) []LRSymbolId {
	switch id {
	case _LRState1:
		return []LRSymbolId{LRTokenToken, LRTypeToken, LRStartToken, LRRuleDefToken}
	case _LRState2:
		return []LRSymbolId{_LREndMarker}
	case _LRState4:
		return []LRSymbolId{LRIdentifierToken}
	case _LRState7:
		return []LRSymbolId{'<', LRCharacterToken, LRIdentifierToken}
	case _LRState8:
		return []LRSymbolId{LRCharacterToken, LRIdentifierToken}
	case _LRState15:
		return []LRSymbolId{LRIdentifierToken}
	case _LRState18:
		return []LRSymbolId{LRIdentifierToken}
	case _LRState19:
		return []LRSymbolId{'>'}
	case _LRState20:
		return []LRSymbolId{LRSectionContentToken}
	case _LRState21:
		return []LRSymbolId{LRCharacterToken, LRIdentifierToken}
	}

	return nil
}

func LRParse(lexer LRLexer, reducer LRReducer) (*Grammar, error) {

	return LRParseWithCustomErrorHandler(
		lexer,
		reducer,
		LRDefaultParseErrorHandler{})
}

func LRParseWithCustomErrorHandler(
	lexer LRLexer,
	reducer LRReducer,
	errHandler LRParseErrorHandler,
) (
	*Grammar,
	error,
) {
	item, err := _LRParse(lexer, reducer, errHandler, _LRState1)
	if err != nil {
		var errRetVal *Grammar
		return errRetVal, err
	}
	return item.Grammar, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _LRParse(
	lexer LRLexer,
	reducer LRReducer,
	errHandler LRParseErrorHandler,
	startState _LRStateId,
) (
	*_LRStackItem,
	error,
) {
	stateStack := _LRStack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_LRStackItem{startState, nil},
	}

	symbolStack := &_LRPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _LRActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _LRShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _LRReduceAction {
			var reduceSymbol *LRSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack,
				nextSymbol.Loc())
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _LRShiftAndReduceAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}

			var reduceSymbol *LRSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack,
				nextSymbol.Loc())
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _LRAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i LRSymbolId) String() string {
	switch i {
	case _LREndMarker:
		return "$"
	case _LRWildcardMarker:
		return "*"
	case LRTokenToken:
		return "TOKEN"
	case LRTypeToken:
		return "TYPE"
	case LRStartToken:
		return "START"
	case LRRuleDefToken:
		return "RULE_DEF"
	case LRLabelToken:
		return "LABEL"
	case '<':
		return "'<'"
	case '>':
		return "'>'"
	case '|':
		return "'|'"
	case ';':
		return "';'"
	case '=':
		return "'='"
	case LRSectionMarkerToken:
		return "SECTION_MARKER"
	case LRCharacterToken:
		return "CHARACTER"
	case LRIdentifierToken:
		return "IDENTIFIER"
	case LRSectionContentToken:
		return "SECTION_CONTENT"
	case LRGrammarType:
		return "grammar"
	case LRAdditionalSectionsType:
		return "additional_sections"
	case LRAdditionalSectionType:
		return "additional_section"
	case LRDefsType:
		return "defs"
	case LRDefType:
		return "def"
	case LRRwordType:
		return "rword"
	case LRNonemptyIdentListType:
		return "nonempty_ident_list"
	case LRNonemptyIdOrCharListType:
		return "nonempty_id_or_char_list"
	case LRIdOrCharListType:
		return "id_or_char_list"
	case LRRuleType:
		return "rule"
	case LRClauseType:
		return "clause"
	case LRClausesType:
		return "clauses"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_LREndMarker      = LRSymbolId(0)
	_LRWildcardMarker = LRSymbolId(-1)

	LRGrammarType              = LRSymbolId(270)
	LRAdditionalSectionsType   = LRSymbolId(271)
	LRAdditionalSectionType    = LRSymbolId(272)
	LRDefsType                 = LRSymbolId(273)
	LRDefType                  = LRSymbolId(274)
	LRRwordType                = LRSymbolId(275)
	LRNonemptyIdentListType    = LRSymbolId(276)
	LRNonemptyIdOrCharListType = LRSymbolId(277)
	LRIdOrCharListType         = LRSymbolId(278)
	LRRuleType                 = LRSymbolId(279)
	LRClauseType               = LRSymbolId(280)
	LRClausesType              = LRSymbolId(281)
)

type _LRActionType int

const (
	// NOTE: error action is implicit
	_LRShiftAction          = _LRActionType(0)
	_LRReduceAction         = _LRActionType(1)
	_LRShiftAndReduceAction = _LRActionType(2)
	_LRAcceptAction         = _LRActionType(3)
)

func (i _LRActionType) String() string {
	switch i {
	case _LRShiftAction:
		return "shift"
	case _LRReduceAction:
		return "reduce"
	case _LRShiftAndReduceAction:
		return "shift-and-reduce"
	case _LRAcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _LRReduceType int

const (
	_LRReduceToGrammar                     = _LRReduceType(1)
	_LRReduceAddToAdditionalSections       = _LRReduceType(2)
	_LRReduceNilToAdditionalSections       = _LRReduceType(3)
	_LRReduceToAdditionalSection           = _LRReduceType(4)
	_LRReduceAddToDefs                     = _LRReduceType(5)
	_LRReduceAddExplicitToDefs             = _LRReduceType(6)
	_LRReduceDefToDefs                     = _LRReduceType(7)
	_LRReduceExplicitDefToDefs             = _LRReduceType(8)
	_LRReduceTermDeclToDef                 = _LRReduceType(9)
	_LRReduceUntypedTermDeclToDef          = _LRReduceType(10)
	_LRReduceStartDeclToDef                = _LRReduceType(11)
	_LRReduceRuleToDef                     = _LRReduceType(12)
	_LRReduceTokenToRword                  = _LRReduceType(13)
	_LRReduceTypeToRword                   = _LRReduceType(14)
	_LRReduceAddToNonemptyIdentList        = _LRReduceType(15)
	_LRReduceIdentToNonemptyIdentList      = _LRReduceType(16)
	_LRReduceAddIdToNonemptyIdOrCharList   = _LRReduceType(17)
	_LRReduceAddCharToNonemptyIdOrCharList = _LRReduceType(18)
	_LRReduceIdToNonemptyIdOrCharList      = _LRReduceType(19)
	_LRReduceCharToNonemptyIdOrCharList    = _LRReduceType(20)
	_LRReduceListToIdOrCharList            = _LRReduceType(21)
	_LRReduceNilToIdOrCharList             = _LRReduceType(22)
	_LRReduceToRule                        = _LRReduceType(23)
	_LRReducePassthroughIdToClause         = _LRReduceType(24)
	_LRReducePassthroughCharToClause       = _LRReduceType(25)
	_LRReduceUnlabeledToClause             = _LRReduceType(26)
	_LRReduceLabeledToClause               = _LRReduceType(27)
	_LRReduceAddToClauses                  = _LRReduceType(28)
	_LRReduceClauseToClauses               = _LRReduceType(29)
)

func (i _LRReduceType) String() string {
	switch i {
	case _LRReduceToGrammar:
		return "ToGrammar"
	case _LRReduceAddToAdditionalSections:
		return "AddToAdditionalSections"
	case _LRReduceNilToAdditionalSections:
		return "NilToAdditionalSections"
	case _LRReduceToAdditionalSection:
		return "ToAdditionalSection"
	case _LRReduceAddToDefs:
		return "AddToDefs"
	case _LRReduceAddExplicitToDefs:
		return "AddExplicitToDefs"
	case _LRReduceDefToDefs:
		return "DefToDefs"
	case _LRReduceExplicitDefToDefs:
		return "ExplicitDefToDefs"
	case _LRReduceTermDeclToDef:
		return "TermDeclToDef"
	case _LRReduceUntypedTermDeclToDef:
		return "UntypedTermDeclToDef"
	case _LRReduceStartDeclToDef:
		return "StartDeclToDef"
	case _LRReduceRuleToDef:
		return "RuleToDef"
	case _LRReduceTokenToRword:
		return "TokenToRword"
	case _LRReduceTypeToRword:
		return "TypeToRword"
	case _LRReduceAddToNonemptyIdentList:
		return "AddToNonemptyIdentList"
	case _LRReduceIdentToNonemptyIdentList:
		return "IdentToNonemptyIdentList"
	case _LRReduceAddIdToNonemptyIdOrCharList:
		return "AddIdToNonemptyIdOrCharList"
	case _LRReduceAddCharToNonemptyIdOrCharList:
		return "AddCharToNonemptyIdOrCharList"
	case _LRReduceIdToNonemptyIdOrCharList:
		return "IdToNonemptyIdOrCharList"
	case _LRReduceCharToNonemptyIdOrCharList:
		return "CharToNonemptyIdOrCharList"
	case _LRReduceListToIdOrCharList:
		return "ListToIdOrCharList"
	case _LRReduceNilToIdOrCharList:
		return "NilToIdOrCharList"
	case _LRReduceToRule:
		return "ToRule"
	case _LRReducePassthroughIdToClause:
		return "PassthroughIdToClause"
	case _LRReducePassthroughCharToClause:
		return "PassthroughCharToClause"
	case _LRReduceUnlabeledToClause:
		return "UnlabeledToClause"
	case _LRReduceLabeledToClause:
		return "LabeledToClause"
	case _LRReduceAddToClauses:
		return "AddToClauses"
	case _LRReduceClauseToClauses:
		return "ClauseToClauses"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _LRStateId int

func (id _LRStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_LRState1  = _LRStateId(1)
	_LRState2  = _LRStateId(2)
	_LRState3  = _LRStateId(3)
	_LRState4  = _LRStateId(4)
	_LRState5  = _LRStateId(5)
	_LRState6  = _LRStateId(6)
	_LRState7  = _LRStateId(7)
	_LRState8  = _LRStateId(8)
	_LRState9  = _LRStateId(9)
	_LRState10 = _LRStateId(10)
	_LRState11 = _LRStateId(11)
	_LRState12 = _LRStateId(12)
	_LRState13 = _LRStateId(13)
	_LRState14 = _LRStateId(14)
	_LRState15 = _LRStateId(15)
	_LRState16 = _LRStateId(16)
	_LRState17 = _LRStateId(17)
	_LRState18 = _LRStateId(18)
	_LRState19 = _LRStateId(19)
	_LRState20 = _LRStateId(20)
	_LRState21 = _LRStateId(21)
	_LRState22 = _LRStateId(22)
)

type LRSymbol struct {
	SymbolId_ LRSymbolId

	Generic_ LRGenericSymbol

	AdditionalSection  *AdditionalSection
	AdditionalSections []*AdditionalSection
	Clause             *Clause
	Clauses            []*Clause
	Definition         Definition
	Definitions        []Definition
	Grammar            *Grammar
	Rule               *Rule
	RuleDef            *RuleDef
	Token              *Token
	Tokens             []*Token
}

func NewSymbol(token LRToken) (*LRSymbol, error) {
	symbol, ok := token.(*LRSymbol)
	if ok {
		return symbol, nil
	}

	symbol = &LRSymbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _LREndMarker, LRTokenToken, LRTypeToken, LRStartToken, '<', '>', '|', ';', '=', LRSectionMarkerToken:
		val, ok := token.(LRGenericSymbol)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting LRGenericSymbol",
				token.Id())
		}
		symbol.Generic_ = val
	case LRRuleDefToken:
		val, ok := token.(*RuleDef)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *RuleDef",
				token.Id())
		}
		symbol.RuleDef = val
	case LRLabelToken, LRCharacterToken, LRIdentifierToken, LRSectionContentToken:
		val, ok := token.(*Token)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *Token",
				token.Id())
		}
		symbol.Token = val
	default:
		return nil, lexutil.NewLocationError(
			token.Loc(),
			"unexpected token type: %s",
			token.Id())
	}
	return symbol, nil
}

func (s *LRSymbol) Id() LRSymbolId {
	return s.SymbolId_
}

func (s *LRSymbol) Loc() LRLocation {
	type locator interface{ Loc() LRLocation }
	switch s.SymbolId_ {
	case LRAdditionalSectionType:
		loc, ok := interface{}(s.AdditionalSection).(locator)
		if ok {
			return loc.Loc()
		}
	case LRAdditionalSectionsType:
		loc, ok := interface{}(s.AdditionalSections).(locator)
		if ok {
			return loc.Loc()
		}
	case LRClauseType:
		loc, ok := interface{}(s.Clause).(locator)
		if ok {
			return loc.Loc()
		}
	case LRClausesType:
		loc, ok := interface{}(s.Clauses).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefType:
		loc, ok := interface{}(s.Definition).(locator)
		if ok {
			return loc.Loc()
		}
	case LRDefsType:
		loc, ok := interface{}(s.Definitions).(locator)
		if ok {
			return loc.Loc()
		}
	case LRGrammarType:
		loc, ok := interface{}(s.Grammar).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleType:
		loc, ok := interface{}(s.Rule).(locator)
		if ok {
			return loc.Loc()
		}
	case LRRuleDefToken:
		loc, ok := interface{}(s.RuleDef).(locator)
		if ok {
			return loc.Loc()
		}
	case LRLabelToken, LRCharacterToken, LRIdentifierToken, LRSectionContentToken:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Loc()
		}
	case LRNonemptyIdentListType, LRNonemptyIdOrCharListType, LRIdOrCharListType:
		loc, ok := interface{}(s.Tokens).(locator)
		if ok {
			return loc.Loc()
		}
	}
	return s.Generic_.Loc()
}

type _LRPseudoSymbolStack struct {
	lexer LRLexer
	top   []*LRSymbol
}

func (stack *_LRPseudoSymbolStack) Top() (*LRSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, lexutil.NewLocationError(
					stack.lexer.CurrentLocation(),
					"unexpected lex error: %s",
					err)
			}
			token = LRGenericSymbol{
				LRSymbolId: _LREndMarker,
				StartPos:   stack.lexer.CurrentLocation(),
			}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_LRPseudoSymbolStack) Push(symbol *LRSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_LRPseudoSymbolStack) Pop() (*LRSymbol, error) {
	if len(stack.top) == 0 {
		return nil, lexutil.NewLocationError(
			stack.lexer.CurrentLocation(),
			"internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _LRStackItem struct {
	StateId _LRStateId

	*LRSymbol
}

type _LRStack []*_LRStackItem

type _LRAction struct {
	ActionType _LRActionType

	ShiftStateId _LRStateId
	ReduceType   _LRReduceType
}

func (act *_LRAction) ShiftItem(symbol *LRSymbol) *_LRStackItem {
	return &_LRStackItem{StateId: act.ShiftStateId, LRSymbol: symbol}
}

func (act *_LRAction) ReduceSymbol(
	reducer LRReducer,
	stack _LRStack,
	nextLoc LRLocation,
) (
	_LRStack,
	*LRSymbol,
	error,
) {
	var err error
	symbol := &LRSymbol{}
	switch act.ReduceType {
	case _LRReduceToGrammar:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRGrammarType
		symbol.Grammar, err = reducer.ToGrammar(args[0].Definitions, args[1].AdditionalSections)
	case _LRReduceAddToAdditionalSections:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.AddToAdditionalSections(args[0].AdditionalSections, args[1].AdditionalSection)
	case _LRReduceNilToAdditionalSections:
		symbol.SymbolId_ = LRAdditionalSectionsType
		symbol.AdditionalSections, err = reducer.NilToAdditionalSections()
	case _LRReduceToAdditionalSection:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRAdditionalSectionType
		symbol.AdditionalSection, err = reducer.ToAdditionalSection(args[0].Generic_, args[1].Token, args[2].Token)
	case _LRReduceAddToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddToDefs(args[0].Definitions, args[1].Definition)
	case _LRReduceAddExplicitToDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.AddExplicitToDefs(args[0].Definitions, args[1].Definition, args[2].Generic_)
	case _LRReduceDefToDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.DefToDefs(args[0].Definition)
	case _LRReduceExplicitDefToDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefsType
		symbol.Definitions, err = reducer.ExplicitDefToDefs(args[0].Definition, args[1].Generic_)
	case _LRReduceTermDeclToDef:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.TermDeclToDef(args[0].Generic_, args[1].Generic_, args[2].Token, args[3].Generic_, args[4].Tokens)
	case _LRReduceUntypedTermDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.UntypedTermDeclToDef(args[0].Generic_, args[1].Tokens)
	case _LRReduceStartDeclToDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.StartDeclToDef(args[0].Generic_, args[1].Tokens)
	case _LRReduceRuleToDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRDefType
		symbol.Definition, err = reducer.RuleToDef(args[0].Rule)
	case _LRReduceTokenToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		//line grammar.lr:49:6
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _LRReduceTypeToRword:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRRwordType
		//line grammar.lr:50:6
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _LRReduceAddToNonemptyIdentList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.AddToNonemptyIdentList(args[0].Tokens, args[1].Token)
	case _LRReduceIdentToNonemptyIdentList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdentListType
		symbol.Tokens, err = reducer.IdentToNonemptyIdentList(args[0].Token)
	case _LRReduceAddIdToNonemptyIdOrCharList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.AddIdToNonemptyIdOrCharList(args[0].Tokens, args[1].Token)
	case _LRReduceAddCharToNonemptyIdOrCharList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.AddCharToNonemptyIdOrCharList(args[0].Tokens, args[1].Token)
	case _LRReduceIdToNonemptyIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.IdToNonemptyIdOrCharList(args[0].Token)
	case _LRReduceCharToNonemptyIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRNonemptyIdOrCharListType
		symbol.Tokens, err = reducer.CharToNonemptyIdOrCharList(args[0].Token)
	case _LRReduceListToIdOrCharList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRIdOrCharListType
		symbol.Tokens, err = reducer.ListToIdOrCharList(args[0].Tokens)
	case _LRReduceNilToIdOrCharList:
		symbol.SymbolId_ = LRIdOrCharListType
		symbol.Tokens, err = reducer.NilToIdOrCharList()
	case _LRReduceToRule:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRRuleType
		symbol.Rule, err = reducer.ToRule(args[0].RuleDef, args[1].Clauses)
	case _LRReducePassthroughIdToClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRClauseType
		symbol.Clause, err = reducer.PassthroughIdToClause(args[0].Generic_, args[1].Token)
	case _LRReducePassthroughCharToClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRClauseType
		symbol.Clause, err = reducer.PassthroughCharToClause(args[0].Generic_, args[1].Token)
	case _LRReduceUnlabeledToClause:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRClauseType
		symbol.Clause, err = reducer.UnlabeledToClause(args[0].Tokens)
	case _LRReduceLabeledToClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LRClauseType
		symbol.Clause, err = reducer.LabeledToClause(args[0].Token, args[1].Tokens)
	case _LRReduceAddToClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LRClausesType
		symbol.Clauses, err = reducer.AddToClauses(args[0].Clauses, args[1].Generic_, args[2].Clause)
	case _LRReduceClauseToClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LRClausesType
		symbol.Clauses, err = reducer.ClauseToClauses(args[0].Clause)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = lexutil.NewLocationError(
			nextLoc,
			"unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _LRActionTableKey struct {
	_LRStateId
	LRSymbolId
}

type _LRActionTableType struct{}

func (_LRActionTableType) Get(
	stateId _LRStateId,
	symbolId LRSymbolId,
) (
	_LRAction,
	bool,
) {
	switch stateId {
	case _LRState1:
		switch symbolId {
		case LRStartToken:
			return _LRAction{_LRShiftAction, _LRState4, 0}, true
		case LRRuleDefToken:
			return _LRAction{_LRShiftAction, _LRState3, 0}, true
		case LRGrammarType:
			return _LRAction{_LRShiftAction, _LRState2, 0}, true
		case LRDefsType:
			return _LRAction{_LRShiftAction, _LRState6, 0}, true
		case LRDefType:
			return _LRAction{_LRShiftAction, _LRState5, 0}, true
		case LRRwordType:
			return _LRAction{_LRShiftAction, _LRState7, 0}, true
		case LRTokenToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceTokenToRword}, true
		case LRTypeToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceTypeToRword}, true
		case LRRuleType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceRuleToDef}, true
		}
	case _LRState2:
		switch symbolId {
		case _LREndMarker:
			return _LRAction{_LRAcceptAction, 0, 0}, true
		}
	case _LRState3:
		switch symbolId {
		case LRLabelToken:
			return _LRAction{_LRShiftAction, _LRState9, 0}, true
		case '=':
			return _LRAction{_LRShiftAction, _LRState8, 0}, true
		case LRNonemptyIdOrCharListType:
			return _LRAction{_LRShiftAction, _LRState11, 0}, true
		case LRClausesType:
			return _LRAction{_LRShiftAction, _LRState10, 0}, true
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}, true
		case LRIdOrCharListType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceUnlabeledToClause}, true
		case LRClauseType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceClauseToClauses}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceNilToIdOrCharList}, true
		}
	case _LRState4:
		switch symbolId {
		case LRNonemptyIdentListType:
			return _LRAction{_LRShiftAction, _LRState12, 0}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdentToNonemptyIdentList}, true
		}
	case _LRState5:
		switch symbolId {
		case ';':
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceExplicitDefToDefs}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceDefToDefs}, true
		}
	case _LRState6:
		switch symbolId {
		case LRStartToken:
			return _LRAction{_LRShiftAction, _LRState4, 0}, true
		case LRRuleDefToken:
			return _LRAction{_LRShiftAction, _LRState3, 0}, true
		case LRAdditionalSectionsType:
			return _LRAction{_LRShiftAction, _LRState13, 0}, true
		case LRDefType:
			return _LRAction{_LRShiftAction, _LRState14, 0}, true
		case LRRwordType:
			return _LRAction{_LRShiftAction, _LRState7, 0}, true
		case LRTokenToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceTokenToRword}, true
		case LRTypeToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceTypeToRword}, true
		case LRRuleType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceRuleToDef}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceNilToAdditionalSections}, true
		}
	case _LRState7:
		switch symbolId {
		case '<':
			return _LRAction{_LRShiftAction, _LRState15, 0}, true
		case LRNonemptyIdOrCharListType:
			return _LRAction{_LRShiftAction, _LRState16, 0}, true
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}, true
		}
	case _LRState8:
		switch symbolId {
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReducePassthroughCharToClause}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReducePassthroughIdToClause}, true
		}
	case _LRState9:
		switch symbolId {
		case LRNonemptyIdOrCharListType:
			return _LRAction{_LRShiftAction, _LRState11, 0}, true
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}, true
		case LRIdOrCharListType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceLabeledToClause}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceNilToIdOrCharList}, true
		}
	case _LRState10:
		switch symbolId {
		case '|':
			return _LRAction{_LRShiftAction, _LRState17, 0}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceToRule}, true
		}
	case _LRState11:
		switch symbolId {
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddIdToNonemptyIdOrCharList}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceListToIdOrCharList}, true
		}
	case _LRState12:
		switch symbolId {
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddToNonemptyIdentList}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceStartDeclToDef}, true
		}
	case _LRState13:
		switch symbolId {
		case LRSectionMarkerToken:
			return _LRAction{_LRShiftAction, _LRState18, 0}, true
		case LRAdditionalSectionType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddToAdditionalSections}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceToGrammar}, true
		}
	case _LRState14:
		switch symbolId {
		case ';':
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddExplicitToDefs}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceAddToDefs}, true
		}
	case _LRState15:
		switch symbolId {
		case LRIdentifierToken:
			return _LRAction{_LRShiftAction, _LRState19, 0}, true
		}
	case _LRState16:
		switch symbolId {
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddIdToNonemptyIdOrCharList}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceUntypedTermDeclToDef}, true
		}
	case _LRState17:
		switch symbolId {
		case LRLabelToken:
			return _LRAction{_LRShiftAction, _LRState9, 0}, true
		case '=':
			return _LRAction{_LRShiftAction, _LRState8, 0}, true
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}, true
		case LRNonemptyIdOrCharListType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceListToIdOrCharList}, true
		case LRIdOrCharListType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceUnlabeledToClause}, true
		case LRClauseType:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddToClauses}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceNilToIdOrCharList}, true
		}
	case _LRState18:
		switch symbolId {
		case LRIdentifierToken:
			return _LRAction{_LRShiftAction, _LRState20, 0}, true
		}
	case _LRState19:
		switch symbolId {
		case '>':
			return _LRAction{_LRShiftAction, _LRState21, 0}, true
		}
	case _LRState20:
		switch symbolId {
		case LRSectionContentToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceToAdditionalSection}, true
		}
	case _LRState21:
		switch symbolId {
		case LRNonemptyIdOrCharListType:
			return _LRAction{_LRShiftAction, _LRState22, 0}, true
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceIdToNonemptyIdOrCharList}, true
		}
	case _LRState22:
		switch symbolId {
		case LRCharacterToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddCharToNonemptyIdOrCharList}, true
		case LRIdentifierToken:
			return _LRAction{_LRShiftAndReduceAction, 0, _LRReduceAddIdToNonemptyIdOrCharList}, true

		default:
			return _LRAction{_LRReduceAction, 0, _LRReduceTermDeclToDef}, true
		}
	}

	return _LRAction{}, false
}

var _LRActionTable = _LRActionTableType{}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.grammar
    Non-kernel Items:
      grammar:.defs additional_sections
      defs:.defs def
      defs:.defs def ';'
      defs:.def
      defs:.def ';'
      def:.rword '<' IDENTIFIER '>' nonempty_id_or_char_list
      def:.rword nonempty_id_or_char_list
      def:.START nonempty_ident_list
      def:.rule
      rword:.TOKEN
      rword:.TYPE
      rule:.RULE_DEF clauses
    Reduce:
      (nil)
    ShiftAndReduce:
      TOKEN -> [rword]
      TYPE -> [rword]
      rule -> [def]
    Goto:
      START -> State 4
      RULE_DEF -> State 3
      grammar -> State 2
      defs -> State 6
      def -> State 5
      rword -> State 7

  State 2:
    Kernel Items:
      #accept: ^ grammar., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 3:
    Kernel Items:
      rule: RULE_DEF.clauses
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
      id_or_char_list:.nonempty_id_or_char_list
      id_or_char_list:., *
      clause:.'=' IDENTIFIER
      clause:.'=' CHARACTER
      clause:.id_or_char_list
      clause:.LABEL id_or_char_list
      clauses:.clauses '|' clause
      clauses:.clause
    Reduce:
      * -> [id_or_char_list]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
      id_or_char_list -> [clause]
      clause -> [clauses]
    Goto:
      LABEL -> State 9
      '=' -> State 8
      nonempty_id_or_char_list -> State 11
      clauses -> State 10

  State 4:
    Kernel Items:
      def: START.nonempty_ident_list
    Non-kernel Items:
      nonempty_ident_list:.nonempty_ident_list IDENTIFIER
      nonempty_ident_list:.IDENTIFIER
    Reduce:
      (nil)
    ShiftAndReduce:
      IDENTIFIER -> [nonempty_ident_list]
    Goto:
      nonempty_ident_list -> State 12

  State 5:
    Kernel Items:
      defs: def., *
      defs: def.';'
    Reduce:
      * -> [defs]
    ShiftAndReduce:
      ';' -> [defs]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      grammar: defs.additional_sections
      defs: defs.def
      defs: defs.def ';'
    Non-kernel Items:
      additional_sections:.additional_sections additional_section
      additional_sections:., *
      def:.rword '<' IDENTIFIER '>' nonempty_id_or_char_list
      def:.rword nonempty_id_or_char_list
      def:.START nonempty_ident_list
      def:.rule
      rword:.TOKEN
      rword:.TYPE
      rule:.RULE_DEF clauses
    Reduce:
      * -> [additional_sections]
    ShiftAndReduce:
      TOKEN -> [rword]
      TYPE -> [rword]
      rule -> [def]
    Goto:
      START -> State 4
      RULE_DEF -> State 3
      additional_sections -> State 13
      def -> State 14
      rword -> State 7

  State 7:
    Kernel Items:
      def: rword.'<' IDENTIFIER '>' nonempty_id_or_char_list
      def: rword.nonempty_id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
    Reduce:
      (nil)
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
    Goto:
      '<' -> State 15
      nonempty_id_or_char_list -> State 16

  State 8:
    Kernel Items:
      clause: '='.IDENTIFIER
      clause: '='.CHARACTER
    Reduce:
      (nil)
    ShiftAndReduce:
      CHARACTER -> [clause]
      IDENTIFIER -> [clause]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      clause: LABEL.id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
      id_or_char_list:.nonempty_id_or_char_list
      id_or_char_list:., *
    Reduce:
      * -> [id_or_char_list]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
      id_or_char_list -> [clause]
    Goto:
      nonempty_id_or_char_list -> State 11

  State 10:
    Kernel Items:
      rule: RULE_DEF clauses., *
      clauses: clauses.'|' clause
    Reduce:
      * -> [rule]
    ShiftAndReduce:
      (nil)
    Goto:
      '|' -> State 17

  State 11:
    Kernel Items:
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
      id_or_char_list: nonempty_id_or_char_list., *
    Reduce:
      * -> [id_or_char_list]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      def: START nonempty_ident_list., *
      nonempty_ident_list: nonempty_ident_list.IDENTIFIER
    Reduce:
      * -> [def]
    ShiftAndReduce:
      IDENTIFIER -> [nonempty_ident_list]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      grammar: defs additional_sections., *
      additional_sections: additional_sections.additional_section
    Non-kernel Items:
      additional_section:.SECTION_MARKER IDENTIFIER SECTION_CONTENT
    Reduce:
      * -> [grammar]
    ShiftAndReduce:
      additional_section -> [additional_sections]
    Goto:
      SECTION_MARKER -> State 18

  State 14:
    Kernel Items:
      defs: defs def., *
      defs: defs def.';'
    Reduce:
      * -> [defs]
    ShiftAndReduce:
      ';' -> [defs]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      def: rword '<'.IDENTIFIER '>' nonempty_id_or_char_list
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 19

  State 16:
    Kernel Items:
      def: rword nonempty_id_or_char_list., *
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
    Reduce:
      * -> [def]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      clauses: clauses '|'.clause
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
      id_or_char_list:.nonempty_id_or_char_list
      id_or_char_list:., *
      clause:.'=' IDENTIFIER
      clause:.'=' CHARACTER
      clause:.id_or_char_list
      clause:.LABEL id_or_char_list
    Reduce:
      * -> [id_or_char_list]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
      nonempty_id_or_char_list -> [id_or_char_list]
      id_or_char_list -> [clause]
      clause -> [clauses]
    Goto:
      LABEL -> State 9
      '=' -> State 8

  State 18:
    Kernel Items:
      additional_section: SECTION_MARKER.IDENTIFIER SECTION_CONTENT
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      IDENTIFIER -> State 20

  State 19:
    Kernel Items:
      def: rword '<' IDENTIFIER.'>' nonempty_id_or_char_list
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      '>' -> State 21

  State 20:
    Kernel Items:
      additional_section: SECTION_MARKER IDENTIFIER.SECTION_CONTENT
    Reduce:
      (nil)
    ShiftAndReduce:
      SECTION_CONTENT -> [additional_section]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      def: rword '<' IDENTIFIER '>'.nonempty_id_or_char_list
    Non-kernel Items:
      nonempty_id_or_char_list:.nonempty_id_or_char_list IDENTIFIER
      nonempty_id_or_char_list:.nonempty_id_or_char_list CHARACTER
      nonempty_id_or_char_list:.IDENTIFIER
      nonempty_id_or_char_list:.CHARACTER
    Reduce:
      (nil)
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
    Goto:
      nonempty_id_or_char_list -> State 22

  State 22:
    Kernel Items:
      def: rword '<' IDENTIFIER '>' nonempty_id_or_char_list., *
      nonempty_id_or_char_list: nonempty_id_or_char_list.IDENTIFIER
      nonempty_id_or_char_list: nonempty_id_or_char_list.CHARACTER
    Reduce:
      * -> [def]
    ShiftAndReduce:
      CHARACTER -> [nonempty_id_or_char_list]
      IDENTIFIER -> [nonempty_id_or_char_list]
    Goto:
      (nil)

Number of states: 22
Number of shift actions: 27
Number of reduce actions: 13
Number of shift-and-reduce actions: 36
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 45
Number of unoptimized shift actions: 63
Number of unoptimized reduce actions: 245
*/
