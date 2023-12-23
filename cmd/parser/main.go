package main

import (
	"fmt"
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

// Parse function to create an AST from tokens (simplified example)
func Parse(tokens []Token) Node {
	// Your parsing logic goes here
	// For now, returning a dummy AST for illustration purposes
	return ClassDef{
		Name: "UserProfileController",
		Methods: []MethodDef{
			{
				Name:   "GetProfile",
				Params: []Param{{Name: "userId", Type: "Integer"}},
				Body:   []Node{},
			},
		},
	}
}

// Main function for testing the lexer and parser
func main() {
	input := `
UserProfileController:
    GetProfile(userId: Integer) -> { GetUser(userId), GetPostStats(userId) }
    `

	tokens := Lex(input)
	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Printf("  %s: '%s'\n", token.Typ, token.Value)
	}

	ast := Parse(tokens)
	fmt.Printf("AST:\n%s\n", ast)
}
