package scanner

import (
	//"fmt"
	"testing"

	"github.com/pravj/glox/scanner/token"
)

func TestScanSingleToken(t *testing.T) {
	tkn := "234"
	s := NewScanner(tkn)
	s.ScanTokens()

	if s.tokens[0].TypeOfToken != token.NUMBER {
		t.Errorf("Incorrect token type for \"%v\". Expected %v, Found %v.", tkn, token.NUMBER, s.tokens[0].TypeOfToken)
	}

	/*
	  tkn = "2ab4"
	  s = NewScanner(tkn)
	  s.ScanTokens()

	  if s.tokens[0].TypeOfToken != token.IDENTIFIER {
	    t.Errorf("Incorrect token type for \"%v\". Expected %v, Found %v.", tkn, token.IDENTIFIER, s.tokens[0].TypeOfToken)
	  }
	*/

	tkn = "123.45"
	s = NewScanner(tkn)
	s.ScanTokens()

	if s.tokens[0].TypeOfToken != token.NUMBER {
		t.Errorf("Incorrect token type for \"%v\". Expected %v, Found %v.", tkn, token.NUMBER, s.tokens[0].TypeOfToken)
	}

	tkn = "\"123.45\""
	s = NewScanner(tkn)
	s.ScanTokens()

	if s.tokens[0].TypeOfToken != token.STRING {
		t.Errorf("Incorrect token type for \"%v\". Expected %v, Found %v.", tkn, token.STRING, s.tokens[0].TypeOfToken)
	}
}
