package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/scanner"
)

// Token struct represents a lexical token
type Token struct {
	Typ   string
	Value string
}

// Lex function tokenizes the input string
func Lex(input string) []Token {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	var tokens []Token

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, Token{Typ: fmt.Sprintf("%s", scanner.TokenString(tok)), Value: s.TokenText()})
	}

	return tokens
}

// Node interface for AST nodes
type Node interface{}

// ClassDef struct represents a class definition
type ClassDef struct {
	Name    string
	Methods []MethodDef
}

// MethodDef struct represents a method definition
type MethodDef struct {
	Name   string
	Params []Param
	Body   []Node
}

// Param struct represents a parameter in a method
type Param struct {
	Name string
	Type string
}

// String methods for pretty-printing
func (cd ClassDef) String() string {
	methodsStr := ""
	for _, method := range cd.Methods {
		methodsStr += method.String() + "\n"
	}
	return fmt.Sprintf("Class Name: %s\nMethods:\n%s", cd.Name, methodsStr)
}

func (md MethodDef) String() string {
	paramsStr := ""
	for _, param := range md.Params {
		paramsStr += param.String() + ", "
	}
	paramsStr = strings.TrimSuffix(paramsStr, ", ")
	return fmt.Sprintf("  Method Name: %s\n  Parameters: [%s]", md.Name, paramsStr)
}

func (p Param) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.Type)
}

// Parse function to create an AST from tokens
func Parse(tokens []Token) ClassDef {
	var classDef ClassDef
	var currentMethod *MethodDef
	var currentParam *Param
	var expectMethodName, expectParam, expectParamType, expectBody bool

	for _, token := range tokens {
		switch {
		case token.Typ == "Ident" && classDef.Name == "":
			// Expecting a class name
			classDef.Name = token.Value

		case token.Typ == "Ident" && expectMethodName:
			// Expecting a method name
			currentMethod = &MethodDef{Name: token.Value}
			expectMethodName = false
			expectParam = true

		case token.Typ == "Ident" && expectParam:
			// Expecting a parameter name
			currentParam = &Param{Name: token.Value}
			expectParam = false
			expectParamType = true

		case token.Typ == "Ident" && expectParamType:
			// Expecting a parameter type
			currentParam.Type = token.Value
			currentMethod.Params = append(currentMethod.Params, *currentParam)
			currentParam = nil
			expectParamType = false
			expectParam = true

		case token.Value == "->":
			// End of method signature, start of method body
			expectBody = true
			expectParam = false

		case token.Typ == "STRING" && expectBody:
			// Dummy handling of method body
			// In a real scenario, you would parse the body properly
			currentMethod.Body = append(currentMethod.Body, token.Value)

		case token.Value == "\n":
			// End of method definition
			if currentMethod != nil {
				classDef.Methods = append(classDef.Methods, *currentMethod)
				currentMethod = nil
				expectMethodName = true
				expectBody = false
			}
		}

		if token.Value == ":" && classDef.Name != "" {
			// Start of method definitions
			expectMethodName = true
		}
	}

	return classDef
}

// Main function for testing the lexer and parser
func main() {
	file, err := os.Open("advanced/Controller7.aur")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	tokens := Lex(input)
	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Printf("  %s: '%s'\n", token.Typ, token.Value)
	}

	ast := Parse(tokens)
	fmt.Printf("AST:\n%s\n", ast)

	// Serialize the AST to JSON
	astJson, err := json.Marshal(ast)
	if err != nil {
		fmt.Println("Error serializing AST:", err)
		return
	}

	// Save the JSON to a file
	err = ioutil.WriteFile("ast.json", astJson, 0644)
	if err != nil {
		fmt.Println("Error writing AST to file:", err)
		return
	}
}
