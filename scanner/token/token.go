// Package token defines (lexems) lexical tokens for the Lox programming language.
package token

import "fmt"

// TokenType represents a set of lexical tokens of the Lox programming language.
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

// Token represents metadata about a lexical token
type Token struct {
	// type of the lexical token
	TypeOfToken TokenType

	// string represents of the lexical token
	Lexeme string

	// line number for the token
	line int
}

// New returns a new Token
func New(typeOfToken TokenType, lexeme string, line int) Token {
	return Token{TypeOfToken: typeOfToken, Lexeme: lexeme, line: line}
}

// toString returns the string representation of the token
func (t *Token) toString() string {
	return fmt.Sprintf("%v %v %v", t.TypeOfToken, t.Lexeme, t.line)
}
