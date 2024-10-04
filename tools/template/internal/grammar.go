// Auto-generated from source: grammar.lr

package template

import (
	fmt "fmt"
	lexutil "github.com/pattyshack/gt/lexutil"
	io "io"
)

type SymbolId int

const (
	SectionMarkerToken = SymbolId(256)
	PackageToken       = SymbolId(257)
	ImportToken        = SymbolId(258)
	TemplateDeclToken  = SymbolId(259)
	ForToken           = SymbolId(260)
	SwitchToken        = SymbolId(261)
	CaseToken          = SymbolId(262)
	IfToken            = SymbolId(263)
	ElseIfToken        = SymbolId(264)
	DefaultToken       = SymbolId(265)
	ElseToken          = SymbolId(266)
	EndToken           = SymbolId(267)
	TextToken          = SymbolId(268)
	SubstitutionToken  = SymbolId(269)
	EmbedToken         = SymbolId(270)
	CopySectionToken   = SymbolId(271)
	CommentToken       = SymbolId(272)
	ContinueToken      = SymbolId(273)
	BreakToken         = SymbolId(274)
	ReturnToken        = SymbolId(275)
	ErrorToken         = SymbolId(276)
)

type Location = lexutil.Location

type Token interface {
	Id() SymbolId
	Loc() Location
}

type GenericSymbol struct {
	SymbolId
	StartPos Location
}

func (t GenericSymbol) Id() SymbolId { return t.SymbolId }

func (t GenericSymbol) Loc() Location { return t.StartPos }

type Lexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return GenericSymbol
	Next() (Token, error)

	CurrentLocation() Location
}

type FileReducer interface {
	// 22:8: file -> ...
	ToFile(Package_ *Value, OptionalImports_ *Value, TemplateDecl_ *TemplateDeclaration, SectionMarker_ GenericSymbol, Body_ []Statement) (*File, error)
}

type OptionalImportsReducer interface {

	// 30:4: optional_imports -> nil: ...
	NilToOptionalImports() (*Value, error)
}

type BodyReducer interface {
	// 33:4: body -> add: ...
	AddToBody(Body_ []Statement, Statement_ Statement) ([]Statement, error)

	// 34:4: body -> nil: ...
	NilToBody() ([]Statement, error)
}

type ForReducer interface {
	// 53:7: for -> ...
	ToFor(For_ *Value, Body_ []Statement, End_ *TToken) (Statement, error)
}

type SwitchReducer interface {
	// 57:4: switch -> with_whitespace: ...
	WithWhitespaceToSwitch(Switch_ *Value, Text_ *Atom, CaseList_ []*Branch, OptionalDefault_ *Branch, End_ *TToken) (Statement, error)

	// 58:4: switch -> without_whitespace: ...
	WithoutWhitespaceToSwitch(Switch_ *Value, CaseList_ []*Branch, OptionalDefault_ *Branch, End_ *TToken) (Statement, error)
}

type CaseListReducer interface {
	// 61:4: case_list -> add: ...
	AddToCaseList(CaseList_ []*Branch, Case_ *Value, Body_ []Statement) ([]*Branch, error)

	// 62:4: case_list -> case: ...
	CaseToCaseList(Case_ *Value, Body_ []Statement) ([]*Branch, error)
}

type OptionalDefaultReducer interface {
	// 65:4: optional_default -> default: ...
	DefaultToOptionalDefault(Default_ *TToken, Body_ []Statement) (*Branch, error)

	// 66:4: optional_default -> nil: ...
	NilToOptionalDefault() (*Branch, error)
}

type IfReducer interface {
	// 68:6: if -> ...
	ToIf(If_ *Value, Body_ []Statement, ElseIfList_ []*Branch, OptionalElse_ *Branch, End_ *TToken) (Statement, error)
}

type ElseIfListReducer interface {
	// 71:4: else_if_list -> add: ...
	AddToElseIfList(ElseIfList_ []*Branch, ElseIf_ *Value, Body_ []Statement) ([]*Branch, error)

	// 72:4: else_if_list -> nil: ...
	NilToElseIfList() ([]*Branch, error)
}

type OptionalElseReducer interface {
	// 75:4: optional_else -> else: ...
	ElseToOptionalElse(Else_ *TToken, Body_ []Statement) (*Branch, error)

	// 76:4: optional_else -> nil: ...
	NilToOptionalElse() (*Branch, error)
}

type Reducer interface {
	FileReducer
	OptionalImportsReducer
	BodyReducer
	ForReducer
	SwitchReducer
	CaseListReducer
	OptionalDefaultReducer
	IfReducer
	ElseIfListReducer
	OptionalElseReducer
}

type ParseErrorHandler interface {
	Error(nextToken Token, parseStack _Stack) error
}

type DefaultParseErrorHandler struct{}

func (DefaultParseErrorHandler) Error(nextToken Token, stack _Stack) error {
	return lexutil.NewLocationError(
		nextToken.Loc(),
		"syntax error: unexpected symbol %s. expecting %v",
		nextToken.Id(),
		ExpectedTerminals(stack[len(stack)-1].StateId))
}

func ExpectedTerminals(id _StateId) []SymbolId {
	switch id {
	case _State1:
		return []SymbolId{PackageToken}
	case _State2:
		return []SymbolId{_EndMarker}
	case _State4:
		return []SymbolId{TemplateDeclToken}
	case _State5:
		return []SymbolId{SectionMarkerToken}
	case _State10:
		return []SymbolId{CaseToken, TextToken}
	case _State11:
		return []SymbolId{ForToken, SwitchToken, IfToken, EndToken, TextToken, SubstitutionToken, EmbedToken, CopySectionToken, CommentToken, ContinueToken, BreakToken, ReturnToken, ErrorToken}
	case _State14:
		return []SymbolId{CaseToken}
	case _State21:
		return []SymbolId{EndToken}
	case _State24:
		return []SymbolId{EndToken}
	case _State25:
		return []SymbolId{EndToken}
	}

	return nil
}

func Parse(lexer Lexer, reducer Reducer) (*File, error) {

	return ParseWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
) (
	*File,
	error,
) {
	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal *File
		return errRetVal, err
	}
	return item.File, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _Parse(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
	startState _StateId,
) (
	*_StackItem,
	error,
) {
	stateStack := _Stack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_StackItem{startState, nil},
	}

	symbolStack := &_PseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _ActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _ShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _ReduceAction {
			var reduceSymbol *Symbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack,
				nextSymbol.Loc())
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _ShiftAndReduceAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}

			var reduceSymbol *Symbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack,
				nextSymbol.Loc())
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _AcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i SymbolId) String() string {
	switch i {
	case _EndMarker:
		return "$"
	case _WildcardMarker:
		return "*"
	case SectionMarkerToken:
		return "SECTION_MARKER"
	case PackageToken:
		return "PACKAGE"
	case ImportToken:
		return "IMPORT"
	case TemplateDeclToken:
		return "TEMPLATE_DECL"
	case ForToken:
		return "FOR"
	case SwitchToken:
		return "SWITCH"
	case CaseToken:
		return "CASE"
	case IfToken:
		return "IF"
	case ElseIfToken:
		return "ELSE_IF"
	case DefaultToken:
		return "DEFAULT"
	case ElseToken:
		return "ELSE"
	case EndToken:
		return "END"
	case TextToken:
		return "TEXT"
	case SubstitutionToken:
		return "SUBSTITUTION"
	case EmbedToken:
		return "EMBED"
	case CopySectionToken:
		return "COPY_SECTION"
	case CommentToken:
		return "COMMENT"
	case ContinueToken:
		return "CONTINUE"
	case BreakToken:
		return "BREAK"
	case ReturnToken:
		return "RETURN"
	case ErrorToken:
		return "ERROR"
	case FileType:
		return "file"
	case OptionalImportsType:
		return "optional_imports"
	case BodyType:
		return "body"
	case StatementType:
		return "statement"
	case AtomType:
		return "atom"
	case ForType:
		return "for"
	case SwitchType:
		return "switch"
	case CaseListType:
		return "case_list"
	case OptionalDefaultType:
		return "optional_default"
	case IfType:
		return "if"
	case ElseIfListType:
		return "else_if_list"
	case OptionalElseType:
		return "optional_else"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	FileType            = SymbolId(277)
	OptionalImportsType = SymbolId(278)
	BodyType            = SymbolId(279)
	StatementType       = SymbolId(280)
	AtomType            = SymbolId(281)
	ForType             = SymbolId(282)
	SwitchType          = SymbolId(283)
	CaseListType        = SymbolId(284)
	OptionalDefaultType = SymbolId(285)
	IfType              = SymbolId(286)
	ElseIfListType      = SymbolId(287)
	OptionalElseType    = SymbolId(288)
)

type _ActionType int

const (
	// NOTE: error action is implicit
	_ShiftAction          = _ActionType(0)
	_ReduceAction         = _ActionType(1)
	_ShiftAndReduceAction = _ActionType(2)
	_AcceptAction         = _ActionType(3)
)

func (i _ActionType) String() string {
	switch i {
	case _ShiftAction:
		return "shift"
	case _ReduceAction:
		return "reduce"
	case _ShiftAndReduceAction:
		return "shift-and-reduce"
	case _AcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _ReduceType int

const (
	_ReduceToFile                    = _ReduceType(1)
	_ReduceImportToOptionalImports   = _ReduceType(2)
	_ReduceNilToOptionalImports      = _ReduceType(3)
	_ReduceAddToBody                 = _ReduceType(4)
	_ReduceNilToBody                 = _ReduceType(5)
	_ReduceAtomToStatement           = _ReduceType(6)
	_ReduceForToStatement            = _ReduceType(7)
	_ReduceSwitchToStatement         = _ReduceType(8)
	_ReduceIfToStatement             = _ReduceType(9)
	_ReduceTextToAtom                = _ReduceType(10)
	_ReduceSubstitutionToAtom        = _ReduceType(11)
	_ReduceEmbedToAtom               = _ReduceType(12)
	_ReduceCopySectionToAtom         = _ReduceType(13)
	_ReduceCommentToAtom             = _ReduceType(14)
	_ReduceContinueToAtom            = _ReduceType(15)
	_ReduceBreakToAtom               = _ReduceType(16)
	_ReduceReturnToAtom              = _ReduceType(17)
	_ReduceErrorToAtom               = _ReduceType(18)
	_ReduceToFor                     = _ReduceType(19)
	_ReduceWithWhitespaceToSwitch    = _ReduceType(20)
	_ReduceWithoutWhitespaceToSwitch = _ReduceType(21)
	_ReduceAddToCaseList             = _ReduceType(22)
	_ReduceCaseToCaseList            = _ReduceType(23)
	_ReduceDefaultToOptionalDefault  = _ReduceType(24)
	_ReduceNilToOptionalDefault      = _ReduceType(25)
	_ReduceToIf                      = _ReduceType(26)
	_ReduceAddToElseIfList           = _ReduceType(27)
	_ReduceNilToElseIfList           = _ReduceType(28)
	_ReduceElseToOptionalElse        = _ReduceType(29)
	_ReduceNilToOptionalElse         = _ReduceType(30)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceToFile:
		return "ToFile"
	case _ReduceImportToOptionalImports:
		return "ImportToOptionalImports"
	case _ReduceNilToOptionalImports:
		return "NilToOptionalImports"
	case _ReduceAddToBody:
		return "AddToBody"
	case _ReduceNilToBody:
		return "NilToBody"
	case _ReduceAtomToStatement:
		return "AtomToStatement"
	case _ReduceForToStatement:
		return "ForToStatement"
	case _ReduceSwitchToStatement:
		return "SwitchToStatement"
	case _ReduceIfToStatement:
		return "IfToStatement"
	case _ReduceTextToAtom:
		return "TextToAtom"
	case _ReduceSubstitutionToAtom:
		return "SubstitutionToAtom"
	case _ReduceEmbedToAtom:
		return "EmbedToAtom"
	case _ReduceCopySectionToAtom:
		return "CopySectionToAtom"
	case _ReduceCommentToAtom:
		return "CommentToAtom"
	case _ReduceContinueToAtom:
		return "ContinueToAtom"
	case _ReduceBreakToAtom:
		return "BreakToAtom"
	case _ReduceReturnToAtom:
		return "ReturnToAtom"
	case _ReduceErrorToAtom:
		return "ErrorToAtom"
	case _ReduceToFor:
		return "ToFor"
	case _ReduceWithWhitespaceToSwitch:
		return "WithWhitespaceToSwitch"
	case _ReduceWithoutWhitespaceToSwitch:
		return "WithoutWhitespaceToSwitch"
	case _ReduceAddToCaseList:
		return "AddToCaseList"
	case _ReduceCaseToCaseList:
		return "CaseToCaseList"
	case _ReduceDefaultToOptionalDefault:
		return "DefaultToOptionalDefault"
	case _ReduceNilToOptionalDefault:
		return "NilToOptionalDefault"
	case _ReduceToIf:
		return "ToIf"
	case _ReduceAddToElseIfList:
		return "AddToElseIfList"
	case _ReduceNilToElseIfList:
		return "NilToElseIfList"
	case _ReduceElseToOptionalElse:
		return "ElseToOptionalElse"
	case _ReduceNilToOptionalElse:
		return "NilToOptionalElse"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _StateId int

func (id _StateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_State1  = _StateId(1)
	_State2  = _StateId(2)
	_State3  = _StateId(3)
	_State4  = _StateId(4)
	_State5  = _StateId(5)
	_State6  = _StateId(6)
	_State7  = _StateId(7)
	_State8  = _StateId(8)
	_State9  = _StateId(9)
	_State10 = _StateId(10)
	_State11 = _StateId(11)
	_State12 = _StateId(12)
	_State13 = _StateId(13)
	_State14 = _StateId(14)
	_State15 = _StateId(15)
	_State16 = _StateId(16)
	_State17 = _StateId(17)
	_State18 = _StateId(18)
	_State19 = _StateId(19)
	_State20 = _StateId(20)
	_State21 = _StateId(21)
	_State22 = _StateId(22)
	_State23 = _StateId(23)
	_State24 = _StateId(24)
	_State25 = _StateId(25)
	_State26 = _StateId(26)
	_State27 = _StateId(27)
	_State28 = _StateId(28)
	_State29 = _StateId(29)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Atom         *Atom
	Branch       *Branch
	Branches     []*Branch
	File         *File
	Statement    Statement
	Statements   []Statement
	TemplateDecl *TemplateDeclaration
	Token        *TToken
	Value        *Value
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case TextToken, SubstitutionToken, EmbedToken, CopySectionToken, CommentToken, ContinueToken, BreakToken, ReturnToken, ErrorToken:
		val, ok := token.(*Atom)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *Atom",
				token.Id())
		}
		symbol.Atom = val
	case _EndMarker, SectionMarkerToken:
		val, ok := token.(GenericSymbol)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting GenericSymbol",
				token.Id())
		}
		symbol.Generic_ = val
	case TemplateDeclToken:
		val, ok := token.(*TemplateDeclaration)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *TemplateDeclaration",
				token.Id())
		}
		symbol.TemplateDecl = val
	case DefaultToken, ElseToken, EndToken:
		val, ok := token.(*TToken)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *TToken",
				token.Id())
		}
		symbol.Token = val
	case PackageToken, ImportToken, ForToken, SwitchToken, CaseToken, IfToken, ElseIfToken:
		val, ok := token.(*Value)
		if !ok {
			return nil, lexutil.NewLocationError(
				token.Loc(),
				"invalid value type for token %s. "+
					"expecting *Value",
				token.Id())
		}
		symbol.Value = val
	default:
		return nil, lexutil.NewLocationError(
			token.Loc(),
			"unexpected token type: %s",
			token.Id())
	}
	return symbol, nil
}

func (s *Symbol) Id() SymbolId {
	return s.SymbolId_
}

func (s *Symbol) Loc() Location {
	type locator interface{ Loc() Location }
	switch s.SymbolId_ {
	case TextToken, SubstitutionToken, EmbedToken, CopySectionToken, CommentToken, ContinueToken, BreakToken, ReturnToken, ErrorToken:
		loc, ok := interface{}(s.Atom).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalDefaultType, OptionalElseType:
		loc, ok := interface{}(s.Branch).(locator)
		if ok {
			return loc.Loc()
		}
	case CaseListType, ElseIfListType:
		loc, ok := interface{}(s.Branches).(locator)
		if ok {
			return loc.Loc()
		}
	case FileType:
		loc, ok := interface{}(s.File).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementType, AtomType, ForType, SwitchType, IfType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case BodyType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case TemplateDeclToken:
		loc, ok := interface{}(s.TemplateDecl).(locator)
		if ok {
			return loc.Loc()
		}
	case DefaultToken, ElseToken, EndToken:
		loc, ok := interface{}(s.Token).(locator)
		if ok {
			return loc.Loc()
		}
	case PackageToken, ImportToken, ForToken, SwitchToken, CaseToken, IfToken, ElseIfToken, OptionalImportsType:
		loc, ok := interface{}(s.Value).(locator)
		if ok {
			return loc.Loc()
		}
	}
	return s.Generic_.Loc()
}

type _PseudoSymbolStack struct {
	lexer Lexer
	top   []*Symbol
}

func (stack *_PseudoSymbolStack) Top() (*Symbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, lexutil.NewLocationError(
					stack.lexer.CurrentLocation(),
					"unexpected lex error: %s",
					err)
			}
			token = GenericSymbol{
				SymbolId: _EndMarker,
				StartPos: stack.lexer.CurrentLocation(),
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

func (stack *_PseudoSymbolStack) Push(symbol *Symbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_PseudoSymbolStack) Pop() (*Symbol, error) {
	if len(stack.top) == 0 {
		return nil, lexutil.NewLocationError(
			stack.lexer.CurrentLocation(),
			"internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _StackItem struct {
	StateId _StateId

	*Symbol
}

type _Stack []*_StackItem

type _Action struct {
	ActionType _ActionType

	ShiftStateId _StateId
	ReduceType   _ReduceType
}

func (act *_Action) ShiftItem(symbol *Symbol) *_StackItem {
	return &_StackItem{StateId: act.ShiftStateId, Symbol: symbol}
}

func (act *_Action) ReduceSymbol(
	reducer Reducer,
	stack _Stack,
	nextLoc Location,
) (
	_Stack,
	*Symbol,
	error,
) {
	var err error
	symbol := &Symbol{}
	switch act.ReduceType {
	case _ReduceToFile:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FileType
		symbol.File, err = reducer.ToFile(args[0].Value, args[1].Value, args[2].TemplateDecl, args[3].Generic_, args[4].Statements)
	case _ReduceImportToOptionalImports:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalImportsType
		//line grammar.lr:29:6
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNilToOptionalImports:
		symbol.SymbolId_ = OptionalImportsType
		symbol.Value, err = reducer.NilToOptionalImports()
	case _ReduceAddToBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BodyType
		symbol.Statements, err = reducer.AddToBody(args[0].Statements, args[1].Statement)
	case _ReduceNilToBody:
		symbol.SymbolId_ = BodyType
		symbol.Statements, err = reducer.NilToBody()
	case _ReduceAtomToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:37:6
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceForToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:38:6
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceSwitchToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:39:6
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceIfToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:40:6
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceTextToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:43:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceSubstitutionToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:44:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceEmbedToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:45:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceCopySectionToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:46:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceCommentToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:47:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceContinueToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:48:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceBreakToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:49:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceReturnToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:50:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceErrorToAtom:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomType
		//line grammar.lr:51:6
		symbol.Statement = args[0].Atom
		err = nil
	case _ReduceToFor:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ForType
		symbol.Statement, err = reducer.ToFor(args[0].Value, args[1].Statements, args[2].Token)
	case _ReduceWithWhitespaceToSwitch:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = SwitchType
		symbol.Statement, err = reducer.WithWhitespaceToSwitch(args[0].Value, args[1].Atom, args[2].Branches, args[3].Branch, args[4].Token)
	case _ReduceWithoutWhitespaceToSwitch:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = SwitchType
		symbol.Statement, err = reducer.WithoutWhitespaceToSwitch(args[0].Value, args[1].Branches, args[2].Branch, args[3].Token)
	case _ReduceAddToCaseList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CaseListType
		symbol.Branches, err = reducer.AddToCaseList(args[0].Branches, args[1].Value, args[2].Statements)
	case _ReduceCaseToCaseList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CaseListType
		symbol.Branches, err = reducer.CaseToCaseList(args[0].Value, args[1].Statements)
	case _ReduceDefaultToOptionalDefault:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Branch, err = reducer.DefaultToOptionalDefault(args[0].Token, args[1].Statements)
	case _ReduceNilToOptionalDefault:
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Branch, err = reducer.NilToOptionalDefault()
	case _ReduceToIf:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfType
		symbol.Statement, err = reducer.ToIf(args[0].Value, args[1].Statements, args[2].Branches, args[3].Branch, args[4].Token)
	case _ReduceAddToElseIfList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ElseIfListType
		symbol.Branches, err = reducer.AddToElseIfList(args[0].Branches, args[1].Value, args[2].Statements)
	case _ReduceNilToElseIfList:
		symbol.SymbolId_ = ElseIfListType
		symbol.Branches, err = reducer.NilToElseIfList()
	case _ReduceElseToOptionalElse:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalElseType
		symbol.Branch, err = reducer.ElseToOptionalElse(args[0].Token, args[1].Statements)
	case _ReduceNilToOptionalElse:
		symbol.SymbolId_ = OptionalElseType
		symbol.Branch, err = reducer.NilToOptionalElse()
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

type _ActionTableKey struct {
	_StateId
	SymbolId
}

type _ActionTableType struct{}

func (_ActionTableType) Get(
	stateId _StateId,
	symbolId SymbolId,
) (
	_Action,
	bool,
) {
	switch stateId {
	case _State1:
		switch symbolId {
		case PackageToken:
			return _Action{_ShiftAction, _State3, 0}, true
		case FileType:
			return _Action{_ShiftAction, _State2, 0}, true
		}
	case _State2:
		switch symbolId {
		case _EndMarker:
			return _Action{_AcceptAction, 0, 0}, true
		}
	case _State3:
		switch symbolId {
		case OptionalImportsType:
			return _Action{_ShiftAction, _State4, 0}, true
		case ImportToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceImportToOptionalImports}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalImports}, true
		}
	case _State4:
		switch symbolId {
		case TemplateDeclToken:
			return _Action{_ShiftAction, _State5, 0}, true
		}
	case _State5:
		switch symbolId {
		case SectionMarkerToken:
			return _Action{_ShiftAction, _State6, 0}, true
		}
	case _State6:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State7, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State7:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceToFile}, true
		}
	case _State8:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State11, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State9:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State12, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State10:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case TextToken:
			return _Action{_ShiftAction, _State14, 0}, true
		case CaseListType:
			return _Action{_ShiftAction, _State15, 0}, true
		}
	case _State11:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case EndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToFor}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true
		}
	case _State12:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case ElseIfListType:
			return _Action{_ShiftAction, _State16, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToElseIfList}, true
		}
	case _State13:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State17, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State14:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State13, 0}, true
		case CaseListType:
			return _Action{_ShiftAction, _State18, 0}, true
		}
	case _State15:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case OptionalDefaultType:
			return _Action{_ShiftAction, _State21, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}, true
		}
	case _State16:
		switch symbolId {
		case ElseIfToken:
			return _Action{_ShiftAction, _State23, 0}, true
		case ElseToken:
			return _Action{_ShiftAction, _State22, 0}, true
		case OptionalElseType:
			return _Action{_ShiftAction, _State24, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalElse}, true
		}
	case _State17:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceCaseToCaseList}, true
		}
	case _State18:
		switch symbolId {
		case CaseToken:
			return _Action{_ShiftAction, _State19, 0}, true
		case DefaultToken:
			return _Action{_ShiftAction, _State20, 0}, true
		case OptionalDefaultType:
			return _Action{_ShiftAction, _State25, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}, true
		}
	case _State19:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State26, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State20:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State27, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State21:
		switch symbolId {
		case EndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWithoutWhitespaceToSwitch}, true
		}
	case _State22:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State28, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State23:
		switch symbolId {
		case BodyType:
			return _Action{_ShiftAction, _State29, 0}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceNilToBody}, true
		}
	case _State24:
		switch symbolId {
		case EndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceToIf}, true
		}
	case _State25:
		switch symbolId {
		case EndToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceWithWhitespaceToSwitch}, true
		}
	case _State26:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddToCaseList}, true
		}
	case _State27:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceDefaultToOptionalDefault}, true
		}
	case _State28:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceElseToOptionalElse}, true
		}
	case _State29:
		switch symbolId {
		case ForToken:
			return _Action{_ShiftAction, _State8, 0}, true
		case SwitchToken:
			return _Action{_ShiftAction, _State10, 0}, true
		case IfToken:
			return _Action{_ShiftAction, _State9, 0}, true
		case TextToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceTextToAtom}, true
		case SubstitutionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSubstitutionToAtom}, true
		case EmbedToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceEmbedToAtom}, true
		case CopySectionToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCopySectionToAtom}, true
		case CommentToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceCommentToAtom}, true
		case ContinueToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceContinueToAtom}, true
		case BreakToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceBreakToAtom}, true
		case ReturnToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceReturnToAtom}, true
		case ErrorToken:
			return _Action{_ShiftAndReduceAction, 0, _ReduceErrorToAtom}, true
		case StatementType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAddToBody}, true
		case AtomType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceAtomToStatement}, true
		case ForType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceForToStatement}, true
		case SwitchType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceSwitchToStatement}, true
		case IfType:
			return _Action{_ShiftAndReduceAction, 0, _ReduceIfToStatement}, true

		default:
			return _Action{_ReduceAction, 0, _ReduceAddToElseIfList}, true
		}
	}

	return _Action{}, false
}

var _ActionTable = _ActionTableType{}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.file
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      PACKAGE -> State 3
      file -> State 2

  State 2:
    Kernel Items:
      #accept: ^ file., $
    Reduce:
      $ -> [#accept]
    ShiftAndReduce:
      (nil)
    Goto:
      (nil)

  State 3:
    Kernel Items:
      file: PACKAGE.optional_imports TEMPLATE_DECL SECTION_MARKER body
    Reduce:
      * -> [optional_imports]
    ShiftAndReduce:
      IMPORT -> [optional_imports]
    Goto:
      optional_imports -> State 4

  State 4:
    Kernel Items:
      file: PACKAGE optional_imports.TEMPLATE_DECL SECTION_MARKER body
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      TEMPLATE_DECL -> State 5

  State 5:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL.SECTION_MARKER body
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      SECTION_MARKER -> State 6

  State 6:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL SECTION_MARKER.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 7

  State 7:
    Kernel Items:
      file: PACKAGE optional_imports TEMPLATE_DECL SECTION_MARKER body., *
      body: body.statement
    Reduce:
      * -> [file]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 8:
    Kernel Items:
      for: FOR.body END
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 11

  State 9:
    Kernel Items:
      if: IF.body else_if_list optional_else END
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 12

  State 10:
    Kernel Items:
      switch: SWITCH.TEXT case_list optional_default END
      switch: SWITCH.case_list optional_default END
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      CASE -> State 13
      TEXT -> State 14
      case_list -> State 15

  State 11:
    Kernel Items:
      body: body.statement
      for: FOR body.END
    Reduce:
      (nil)
    ShiftAndReduce:
      END -> [for]
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 12:
    Kernel Items:
      body: body.statement
      if: IF body.else_if_list optional_else END
    Reduce:
      * -> [else_if_list]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9
      else_if_list -> State 16

  State 13:
    Kernel Items:
      case_list: CASE.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 17

  State 14:
    Kernel Items:
      switch: SWITCH TEXT.case_list optional_default END
    Reduce:
      (nil)
    ShiftAndReduce:
      (nil)
    Goto:
      CASE -> State 13
      case_list -> State 18

  State 15:
    Kernel Items:
      switch: SWITCH case_list.optional_default END
      case_list: case_list.CASE body
    Reduce:
      * -> [optional_default]
    ShiftAndReduce:
      (nil)
    Goto:
      CASE -> State 19
      DEFAULT -> State 20
      optional_default -> State 21

  State 16:
    Kernel Items:
      if: IF body else_if_list.optional_else END
      else_if_list: else_if_list.ELSE_IF body
    Reduce:
      * -> [optional_else]
    ShiftAndReduce:
      (nil)
    Goto:
      ELSE_IF -> State 23
      ELSE -> State 22
      optional_else -> State 24

  State 17:
    Kernel Items:
      body: body.statement
      case_list: CASE body., *
    Reduce:
      * -> [case_list]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 18:
    Kernel Items:
      switch: SWITCH TEXT case_list.optional_default END
      case_list: case_list.CASE body
    Reduce:
      * -> [optional_default]
    ShiftAndReduce:
      (nil)
    Goto:
      CASE -> State 19
      DEFAULT -> State 20
      optional_default -> State 25

  State 19:
    Kernel Items:
      case_list: case_list CASE.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 26

  State 20:
    Kernel Items:
      optional_default: DEFAULT.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 27

  State 21:
    Kernel Items:
      switch: SWITCH case_list optional_default.END
    Reduce:
      (nil)
    ShiftAndReduce:
      END -> [switch]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      optional_else: ELSE.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 28

  State 23:
    Kernel Items:
      else_if_list: else_if_list ELSE_IF.body
    Reduce:
      * -> [body]
    ShiftAndReduce:
      (nil)
    Goto:
      body -> State 29

  State 24:
    Kernel Items:
      if: IF body else_if_list optional_else.END
    Reduce:
      (nil)
    ShiftAndReduce:
      END -> [if]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      switch: SWITCH TEXT case_list optional_default.END
    Reduce:
      (nil)
    ShiftAndReduce:
      END -> [switch]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      body: body.statement
      case_list: case_list CASE body., *
    Reduce:
      * -> [case_list]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 27:
    Kernel Items:
      body: body.statement
      optional_default: DEFAULT body., *
    Reduce:
      * -> [optional_default]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 28:
    Kernel Items:
      body: body.statement
      optional_else: ELSE body., *
    Reduce:
      * -> [optional_else]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

  State 29:
    Kernel Items:
      body: body.statement
      else_if_list: else_if_list ELSE_IF body., *
    Reduce:
      * -> [else_if_list]
    ShiftAndReduce:
      TEXT -> [atom]
      SUBSTITUTION -> [atom]
      EMBED -> [atom]
      COPY_SECTION -> [atom]
      COMMENT -> [atom]
      CONTINUE -> [atom]
      BREAK -> [atom]
      RETURN -> [atom]
      ERROR -> [atom]
      statement -> [body]
      atom -> [statement]
      for -> [statement]
      switch -> [statement]
      if -> [statement]
    Goto:
      FOR -> State 8
      SWITCH -> State 10
      IF -> State 9

Number of states: 29
Number of shift actions: 52
Number of reduce actions: 20
Number of shift-and-reduce actions: 117
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 138
Number of unoptimized shift actions: 334
Number of unoptimized reduce actions: 1243
*/
