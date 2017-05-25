// Package scanner implements the library to handle the lexical-analysis.
package scanner

import (
	"fmt"
	"strings"

	"github.com/pravj/glox/errors"
	"github.com/pravj/glox/scanner/token"
	"github.com/pravj/glox/utils"
)

/*
Scanner represents a lexical scanner that does lexical-analysis on the source.

It holds the basic metadata like the input to scanner and in-process state.
*/
type Scanner struct {
	// source code input to the scanner
	source string

	// slice that holds the lines of source code.
	lines []string

	// slice that holds the tokens.
	tokens []token.Token

	// map strings to language reserved tokens (keywords).
	keywords map[string]token.TokenType

	// the starting position of in-process scan.
	// in-process scan means the hunt for the next lexeme after finding a delimiter (whitespace).
	start int

	// the current position of in-process can.
	current int

	// the current line number for the in-process scan.
	line int
}

// NewScanner returns a new scanner instance (*Scanner).
func NewScanner(source string) *Scanner {
	return &Scanner{
		source:   source,
		lines:    strings.Split(source, "\n"),
		tokens:   make([]token.Token, 0),
		keywords: utils.KeywordMap(),
		line:     1}
}

/*
ScanTokens keeps scanning the source code untils it find the EOF delimiter.
It returns a slice of tokens that represents the entire source code.
*/
func (s *Scanner) ScanTokens() []token.Token {
	for !s.scanComplete() {
		s.start = s.current
		s.scanToken()
	}

	// append the terminal token (End Of File)
	s.tokens = append(s.tokens, token.New(token.EOF, "", s.line))

	return s.tokens
}

// scanToken infers the type of lexical token from its string representation.
func (s *Scanner) scanToken() {
	nextChar := s.nextCharacter()

	switch nextChar {
	case "(":
		s.addToken(token.LEFT_PAREN)
		break
	case ")":
		s.addToken(token.RIGHT_PAREN)
		break
	case "{":
		s.addToken(token.LEFT_BRACE)
		break
	case "}":
		s.addToken(token.RIGHT_BRACE)
		break
	case ",":
		s.addToken(token.COMMA)
		break
	case ";":
		s.addToken(token.SEMICOLON)
		break
	case ".":
		s.addToken(token.DOT)
		break
	case "+":
		s.addToken(token.PLUS)
		break
	case "-":
		s.addToken(token.MINUS)
		break
	case "*":
		s.addToken(token.STAR)
		break
	case " ":
		break
	case "\r":
		break
	case "\t":
		break
	case "\n":
		s.line++
		break
	case "!":
		s.addConditionalToken("=", token.BANG_EQUAL, token.BANG)
		break
	case "=":
		s.addConditionalToken("=", token.EQUAL_EQUAL, token.EQUAL)
		break
	case "<":
		s.addConditionalToken("=", token.LESS_EQUAL, token.LESS)
		break
	case ">":
		s.addConditionalToken("=", token.GREATER_EQUAL, token.GREATER)
		break
	case "/":
		if s.matchCharacter("/") {
			for s.lookahead(0) != "\n" && !s.scanComplete() {
				s.nextCharacter()
			}
		} else {
			s.addToken(token.SLASH)
		}
		break
	case "\"":
		s.scanStringLiteral()
		break
	default:
		if utils.IsDigitCharacter(nextChar) {
			s.scanNumberLiteral()
		} else if utils.IsAlphaCharacter(nextChar) {
			s.scanIdentifier()
		} else {
			errors.ReportError(s.line, s.lines[s.line-1], fmt.Sprintf("Unexpected token %v", nextChar))
		}
	}
}

// scanIdentifier scans the prospective identifier-type token.
func (s *Scanner) scanIdentifier() {
	for utils.IsAlphaNumericCharacter(s.lookahead(0)) {
		s.nextCharacter()
	}

	text := string(s.source[s.start:s.current])
	if _, present := s.keywords[text]; present {
		s.addToken(s.keywords[text])
	} else {
		s.addToken(token.IDENTIFIER)
	}
}

// scanStringLiteral scans the prospective string-type token.
func (s *Scanner) scanStringLiteral() {
	for s.lookahead(0) != "\"" && !s.scanComplete() {
		if s.lookahead(0) == "\n" {
			s.line++
		}

		s.nextCharacter()
	}

	if s.scanComplete() {
		errors.ReportError(s.line, s.lines[s.line-1], fmt.Sprintf("Unterminated string"))
		return
	}

	// the closing quote
	s.nextCharacter()

	// trim the surrounding quotes
	// TODO: deal with the literal value for a token (use interface)
	// fmt.Println(string(s.source[s.start+1:s.current-1]))
	s.addToken(token.STRING)
}

// scanNumberLiteral scans the prospective number-type token.
func (s *Scanner) scanNumberLiteral() {
	for utils.IsDigitCharacter(s.lookahead(0)) {
		s.nextCharacter()
	}

	if s.lookahead(0) == "." && utils.IsDigitCharacter(s.lookahead(1)) {
		s.nextCharacter()

		for utils.IsDigitCharacter(s.lookahead(0)) {
			s.nextCharacter()
		}
	}

	s.addToken(token.NUMBER)
}

// addToken appends the given token to the scanner's tokens.
func (s *Scanner) addToken(tokenType token.TokenType) {
	lexemeStringValue := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.New(tokenType, lexemeStringValue, s.line))
}

// addConditionalToken appends one of the tokens to the scanner's tokens.
func (s *Scanner) addConditionalToken(expectedChar string, trueToken token.TokenType, falseToken token.TokenType) {
	// if the next character is as expected, it will add the 'trueToken', else 'falseToken'
	if s.matchCharacter(expectedChar) {
		s.addToken(trueToken)
	} else {
		s.addToken(falseToken)
	}
}

// nextCharacter returns the current character in source code, then it moves on.
func (s *Scanner) nextCharacter() string {
	s.current++
	return string(s.source[s.current-1])
}

// matchCharacter checks if the next character is as expected, returns true if successful.
// It consumes the next character only if it matches as expected.
func (s *Scanner) matchCharacter(expectedChar string) bool {
	if s.scanComplete() {
		return false
	}

	if string(s.source[s.current]) != expectedChar {
		return false
	}

	s.current++
	return true
}

// lookahead returns the next character on a given position from the current location of scanner.
func (s *Scanner) lookahead(length int) string {
	if s.current+length >= len(s.source) {
		return "\000"
	}

	return string(s.source[s.current+length])
}

// scanComplete checks if the scanner has finished entire source code, returns true if successful.
func (s *Scanner) scanComplete() bool {
	return s.current >= len(s.source)
}
