/*
Package main implements the toolchain for the Lox language.

- AST (Abstarct Syntax Tree) generator
*/
package main

import (
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
		"Binary : left Expression, operator token.TokenType, right Expression",
		"Group : expr Expression",
		"Literal : value interface{}, literalType string",
		"Unary : operator token.TokenType, right Expression",
	}

	funcMap = template.FuncMap{
		"defineTypeClass": defineTypeClass,
		"typeClassName":   typeClassName,
		"typeClassFields": typeClassFields,
		"trimString":      trimString,
		"capitolCaseString": capitolCaseString,
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
	fieldString := trimString(strings.Split(annotation, ":")[1])
	return strings.Split(fieldString, ",")
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
