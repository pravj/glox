// Package ast implements the Abstarct-Syntax-Tree interface of the Lox language.
package ast

import (
	"github.com/pravj/glox/scanner/token"
)

// interface to facilitate "visitor pattern"
type Visitor interface {
	visitBinaryExpression(expr Binary) interface{}

	visitGroupExpression(expr Group) interface{}

	visitLiteralExpression(expr Literal) interface{}

	visitUnaryExpression(expr Unary) interface{}
}

// base interface for expressions
type Expression interface {
	Accept(v Visitor) interface{}
}

type Binary struct {
	Left Expression

	Operator token.Token

	Right Expression
}

func NewBinaryExpression(left Expression, operator token.Token, right Expression) *Binary {
	return &Binary{Left: left, Operator: operator, Right: right}
}

func (expr *Binary) Accept(v Visitor) interface{} {
	return v.visitBinaryExpression(*expr)
}

type Group struct {
	Expr Expression
}

func NewGroupExpression(expr Expression) *Group {
	return &Group{Expr: expr}
}

func (expr *Group) Accept(v Visitor) interface{} {
	return v.visitGroupExpression(*expr)
}

type Literal struct {
	Value interface{}

	LiteralType string
}

func NewLiteralExpression(value interface{}, literalType string) *Literal {
	return &Literal{Value: value, LiteralType: literalType}
}

func (expr *Literal) Accept(v Visitor) interface{} {
	return v.visitLiteralExpression(*expr)
}

type Unary struct {
	Operator token.Token

	Right Expression
}

func NewUnaryExpression(operator token.Token, right Expression) *Unary {
	return &Unary{Operator: operator, Right: right}
}

func (expr *Unary) Accept(v Visitor) interface{} {
	return v.visitUnaryExpression(*expr)
}
