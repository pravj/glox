// Package ast implements the Abstarct-Syntax-Tree interface of the Lox language.
package ast

import (
	"bytes"
	"fmt"
)

// PrinterAST represents the AST printer structure
type PrinterAST struct {
}

func NewPrinterAST() *PrinterAST {
	return &PrinterAST{}
}

// visitBinaryExpression
func (p *PrinterAST) visitBinaryExpression(expr Binary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

// visitGroupExpression
func (p *PrinterAST) visitGroupExpression(expr Group) interface{} {
	return p.parenthesize("group", expr.Expr)
}

// visitLiteralExpression
func (p *PrinterAST) visitLiteralExpression(expr Literal) interface{} {
	return expr.Value
}

// visitUnaryExpression
func (p *PrinterAST) visitUnaryExpression(expr Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

// prettyPrint properly formats the input expression with parentheses and whitespace.
func (p *PrinterAST) PrettyPrint(expr Expression) string {
	return expr.Accept(p).(string)
}

// parenthesize takes a name and a list of subexpressions and wraps them in parentheses.
func (p *PrinterAST) parenthesize(name string, expressions ...Expression) string {
	var strBuffer bytes.Buffer

	strBuffer.WriteString("(")
	strBuffer.WriteString(name)
	for _, expression := range expressions {
		strBuffer.WriteString(" ")
		fmt.Println(expression.Accept(p))
		strBuffer.WriteString(expression.Accept(p).(string))
	}
	strBuffer.WriteString(")")

	return strBuffer.String()
}
