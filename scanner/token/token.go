// Package token defines constants (lexems) representing the lexical tokens of Lox
// programming language.
package token

import "fmt"

// TokenType is the set of lexical tokens of the Lox programming language.
type TokenType int

// TokenType set elements, all the lexems for the language
const (
	EOF TokenType = iota

	// single characters
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	SEMICOLON

	// single character operators
	BANG
	EQUAL
	GREATER
	LESS
	MINUS
	PLUS
	SLASH
	STAR

	// double character operators
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER_EQUAL
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

type Token struct {
	lexeme    string
	line      int
	tokenType TokenType
}

func New(tokenType TokenType, lexeme string, line int) Token {
	return Token{tokenType:tokenType, lexeme:lexeme, line:line}
}

func (t *Token) toString() string {
	return fmt.Sprintf("%v %v %v", t.tokenType, t.lexeme, t.line)
}
