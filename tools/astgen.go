/*
Package main implements the toolchain for the Lox language.

- AST (Abstarct Syntax Tree) generator
*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/pravj/glox/utils"
)

// TemplateData represents the structure of template input data
type TemplateData struct {
	TypeAnnotations []string
}

// funcMap is a map of function names to functions
var funcMap template.FuncMap

/*
generateAST parses the given annotations
it produces the expression class for AST in a new file.
*/
func generateAST() {
	// top level hierarchy representing the base class for grammar productions
	baseClassName := "Expression"

	// name of the output file
	// strings.ToLower(baseClassName)

	typeAnnotations := []string{
		"Binary : left Expression, operator token.Token, right Expression",
		"Group : expr Expression",
		"Literal : value interface{}, literalType string",
		"Unary : operator token.Token, right Expression",
	}

	funcMap = template.FuncMap{
		"defineTypeClass":      defineTypeClass,
		"typeClassName":        typeClassName,
		"typeClassFields":      typeClassFields,
		"typeClassFieldString": typeClassFieldString,
		"typeClassFieldArgs":   typeClassFieldArgs,
		"trimString":           trimString,
		"capitolCaseString":    capitolCaseString,
	}

	defineAST(baseClassName, typeAnnotations)
}

// defineAST defines the annotations to generate the language AST.
func defineAST(baseClassName string, classTypes []string) {
	tmpl := template.Must(template.New("expression.tmpl").Funcs(funcMap).ParseFiles("expression.tmpl"))
	err := tmpl.Execute(os.Stdout, TemplateData{classTypes})
	utils.CheckError(err)
}

// typeClassName returns the name of type class from annotations
func typeClassName(annotation string) string {
	return trimString(strings.Split(annotation, ":")[0])
}

/*
typeClassFields returns the string slices containing type fields
Each element in the slice has the following format
["NAME1 TYPE1" "NAME2 TYPE2" "NAME3 TYPE3"...]
*/
func typeClassFields(annotation string) []string {
	return strings.Split(typeClassFieldString(annotation), ",")
}

// typeClassFieldString returns a string representation of arguments
// NAME1 TYPE1, NAME2 TYPE2, NAME3 TYPE3 ...
func typeClassFieldString(annotation string) string {
	fieldString := trimString(strings.Split(annotation, ":")[1])
	return fieldString
}

// typeClassFieldArgs returns the struct constructor arguments as a string
func typeClassFieldArgs(fields []string) string {
	var argStrBuffer bytes.Buffer

	// {arg1:Arg1, arg2:Arg2 ...}
	argStrBuffer.WriteString("{")

	for i, field := range fields {
		arg := strings.Split(trimString(field), " ")[0]

		argStrBuffer.WriteString(strings.Title(arg))
		argStrBuffer.WriteString(":")
		argStrBuffer.WriteString(arg)

		// prevent from a trailing comma
		if i+1 < len(fields) {
			argStrBuffer.WriteString(", ")
		}
	}

	// closing brace
	argStrBuffer.WriteString("}")

	return argStrBuffer.String()
}

// trimString trims whitespace from a string using strings.TrimSpace
func trimString(str string) string {
	return strings.TrimSpace(str)
}

// capitolCaseString changes the very first letter of a string to capitol case
func capitolCaseString(str string) string {
	splitStrings := strings.Split(str, " ")
	return fmt.Sprintf("%v %v", strings.Title(splitStrings[0]), splitStrings[1])
}

// defineTypeClass
func defineTypeClass(typeAnnotation string) string {
	return typeClassName(typeAnnotation)
}
