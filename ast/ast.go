package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	filename := "../pure/controller8.go"

	fs := token.NewFileSet() // positions are relative to fs
	node, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Print(fs, node)
}
