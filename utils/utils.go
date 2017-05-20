// Package utils implements the utility functions for the interpreter.
package utils

import (
	"github.com/pravj/glox/scanner/token"
)

// CheckError panics if there is an error.
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// IsDigit returns true if the input character represents a digit (0-9).
func IsDigit(character string) bool {
	return (character >= "0" && character <= "9")
}

// IsAlpha returns true if the input character represents an alphabet.
// The following set of characters are allowed, (a-z, A-Z, _).
func IsAlpha(c string) bool {
	return ((c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c == "_"))
}

// IsAlphaNumeric returns true if the input character is either digit or alphabet.
func IsAlphaNumeric(character string) bool {
	return (IsDigit(character) || IsAlpha(character))
}

// KeywordMap returns a map of all the language keywords (tokens).
func KeywordMap() map[string]token.TokenType {
	keywords := make(map[string]token.TokenType)

	keywords["and"] = token.AND
	keywords["class"] = token.CLASS
	keywords["else"] = token.ELSE
	keywords["false"] = token.FALSE
	keywords["for"] = token.FOR
	keywords["fun"] = token.FUN
	keywords["if"] = token.IF
	keywords["nil"] = token.NIL
	keywords["or"] = token.OR
	keywords["print"] = token.PRINT
	keywords["return"] = token.RETURN
	keywords["super"] = token.SUPER
	keywords["this"] = token.THIS
	keywords["true"] = token.TRUE
	keywords["var"] = token.VAR
	keywords["while"] = token.WHILE

	return keywords
}
