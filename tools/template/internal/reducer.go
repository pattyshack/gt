package template

import (
	"fmt"

	"github.com/pattyshack/gt/lexutil"
)

type ReducerImpl struct{}

var _ Reducer = ReducerImpl{}

func (ReducerImpl) ToFile(
	pkg *Value,
	imports *Value,
	template *TemplateDeclaration,
	sectionMarker GenericSymbol,
	body []Statement) (
	*File,
	error) {

	importsVal := ""
	if imports != nil {
		importsVal = imports.Value
	}
	return &File{
		PackageName:         pkg.Value,
		Imports:             importsVal,
		TemplateDeclaration: template,
		Body:                body,
	}, nil
}

func (ReducerImpl) NilToOptionalImports() (*Value, error) {
	return nil, nil
}

func (ReducerImpl) AddToBody(
	body []Statement,
	statement Statement) (
	[]Statement,
	error) {

	return append(body, statement), nil
}

func (ReducerImpl) NilToBody() ([]Statement, error) {
	return nil, nil
}

func (ReducerImpl) ToFor(
	for_ *Value,
	body []Statement,
	end *TToken) (
	Statement,
	error) {

	return &For{
		Branch{for_, body},
	}, nil
}

func (ReducerImpl) WithWhitespaceToSwitch(
	switch_ *Value,
	whitespace *Atom,
	cases []*Branch,
	default_ *Branch,
	end *TToken) (
	Statement,
	error) {

	for _, char := range whitespace.Value {
		if !lexutil.IsWhitespace(char) {
			return nil, fmt.Errorf(
				"Text found between [[switch]] and [[case]] (%s)",
				whitespace.Loc())
		}
	}

	return &Switch{
		switch_,
		cases,
		default_,
	}, nil
}

func (ReducerImpl) WithoutWhitespaceToSwitch(
	switch_ *Value,
	cases []*Branch,
	default_ *Branch,
	end *TToken) (
	Statement,
	error) {

	return &Switch{
		switch_,
		cases,
		default_,
	}, nil
}

func (ReducerImpl) AddToCaseList(
	cases []*Branch,
	predicate *Value,
	body []Statement) (
	[]*Branch,
	error) {

	return append(cases, &Branch{predicate, body}), nil
}

func (ReducerImpl) CaseToCaseList(
	predicate *Value,
	body []Statement) (
	[]*Branch,
	error) {

	return []*Branch{&Branch{predicate, body}}, nil
}

func (ReducerImpl) DefaultToOptionalDefault(
	default_ *TToken,
	body []Statement) (
	*Branch,
	error) {

	return &Branch{nil, body}, nil
}

func (ReducerImpl) NilToOptionalDefault() (*Branch, error) {
	return nil, nil
}

func (ReducerImpl) ToIf(
	predicate *Value,
	body []Statement,
	elseIfs []*Branch,
	else_ *Branch,
	end *TToken) (
	Statement,
	error) {

	return &If{
		Branch{predicate, body},
		elseIfs,
		else_,
	}, nil
}

func (ReducerImpl) AddToElseIfList(
	elseIfs []*Branch,
	predicate *Value,
	body []Statement) (
	[]*Branch,
	error) {

	return append(elseIfs, &Branch{predicate, body}), nil
}

func (ReducerImpl) NilToElseIfList() ([]*Branch, error) {
	return []*Branch{}, nil
}

func (ReducerImpl) ElseToOptionalElse(
	else_ *TToken,
	body []Statement) (
	*Branch,
	error) {

	return &Branch{nil, body}, nil
}

func (ReducerImpl) NilToOptionalElse() (*Branch, error) {
	return nil, nil
}
