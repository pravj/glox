// Package errors implements the library to handle compile time errors.
// It also implements the error reporting by the interpreter.
package errors

import (
	"fmt"
)

// ReportError points out to the error in the source code.
func ReportError(line int, where string, message string) {
	fmt.Printf("Error: %v\n", message)
	fmt.Printf("  %v | %v\n", line, where)
	//fmt.Printf("%v^--\n", strings.Repeat(" ", 5 + column + len(strconv.Itoa(line))))
}
