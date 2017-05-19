package scanner

import "github.com/pravj/glox/scanner/token"
import "github.com/pravj/glox/errors"
import "fmt"
import "strings"
import "github.com/pravj/glox/utils"

type Scanner struct {
	source string
	lines  []string
	tokens []token.Token

  keywords map[string]token.TokenType

	start   int
	current int
	line    int
}

func New(source string) *Scanner {
	return &Scanner{
    source: source,
    lines: strings.Split(source, "\n"),
    tokens: make([]token.Token, 0),
    keywords: utils.KeywordMap(),
    line: 1}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.scanComplete() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.New(token.EOF, "", s.line))
	fmt.Println(s.tokens)
	return s.tokens
}

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
      for (s.lookahead(0) != "\n" && !s.scanComplete()) {
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
    if utils.IsDigit(nextChar) {
      s.scanNumberLiteral()
    } else if utils.IsAlpha(nextChar) {
      s.scanIdentifier()
    } else {
      errors.ReportError(s.line, s.lines[s.line-1], fmt.Sprintf("Unexpected token %v", nextChar))
    }
	}
}

func (s *Scanner) scanIdentifier() {
  for utils.IsAlphaNumeric(s.lookahead(0)) {
    s.nextCharacter()
  }

  text := string(s.source[s.start:s.current])
  if _, present := s.keywords[text]; present {
    s.addToken(s.keywords[text])
  } else {
    s.addToken(token.IDENTIFIER)
  }
}

func (s *Scanner) scanStringLiteral() {
  for (s.lookahead(0) != "\"" && !s.scanComplete()) {
    if s.lookahead(0) == "\n" { s.line++ }

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

func (s *Scanner) scanNumberLiteral() {
  for utils.IsDigit(s.lookahead(0)) {
    s.nextCharacter()
  }

  if (s.lookahead(0) == "." && utils.IsDigit(s.lookahead(1))) {
    s.nextCharacter()

    for utils.IsDigit(s.lookahead(0)) {
      s.nextCharacter()
    }
  }

  s.addToken(token.NUMBER)
}

func (s *Scanner) addToken(tokenType token.TokenType) {
	lexemeStringValue := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.New(tokenType, lexemeStringValue, s.line))
}

func (s *Scanner) addConditionalToken(expectedChar string, trueToken token.TokenType, falseToken token.TokenType) {
  if s.matchCharacter(expectedChar) {
    s.addToken(trueToken)
  } else {
    s.addToken(falseToken)
  }
}

func (s *Scanner) nextCharacter() string {
	s.current++
	return string(s.source[s.current-1])
}

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

func (s *Scanner) lookahead(length int) string {
  if s.current + length >= len(s.source) {
    return "\000"
  }

  return string(s.source[s.current + length])
}

func (s *Scanner) scanComplete() bool {
	return s.current >= len(s.source)
}
