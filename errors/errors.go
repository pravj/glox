package errors

import (
  "fmt"
)

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func ReportError(line int, where string, message string) {
  fmt.Printf("Error: %v\n", message)
  fmt.Printf("  %v | %v\n", line, where)
  //fmt.Printf("%v^--\n", strings.Repeat(" ", 5 + column + len(strconv.Itoa(line))))
}