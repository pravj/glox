// Package ast implements the Abstarct-Syntax-Tree interface of the Lox language.
package ast

import (
	"github.com/pravj/glox/scanner/token"
)

// interface to facilitate "visitor pattern"
type Visitor interface {
	visitBinaryExpression(expr *Binary)

	visitGroupExpression(expr *Group)

	visitLiteralExpression(expr *Literal)

	visitUnaryExpression(expr *Unary)
}

// base struct for expressions
type Expression struct {
}

func (expr *Expression) accept(v Visitor) {
}

type Binary struct {
	Left Expression

	Operator token.TokenType

	Right Expression

	*Expression
}

func (expr *Binary) accept(v Visitor) {
	v.visitBinaryExpression(expr)
}

type Group struct {
	Expr Expression

	*Expression
}

func (expr *Group) accept(v Visitor) {
	v.visitGroupExpression(expr)
}

type Literal struct {
	Value interface{}

	LiteralType string

	*Expression
}

func (expr *Literal) accept(v Visitor) {
	v.visitLiteralExpression(expr)
}

type Unary struct {
	Operator token.TokenType

	Right Expression

	*Expression
}

func (expr *Unary) accept(v Visitor) {
	v.visitUnaryExpression(expr)
}
