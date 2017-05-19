package utils

import (
  "github.com/pravj/glox/scanner/token"
)

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func IsDigit(character string) bool {
  return (character >= "0" && character <= "9")
}

func IsAlpha(c string) bool {
  return ((c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c == "_"))
}

func IsAlphaNumeric(character string) bool {
  return (IsDigit(character) || IsAlpha(character))
}

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
