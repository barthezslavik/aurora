package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	// Example: Convert a simple custom DSL construct to Go AST
	// Let's say our DSL has a construct like "create function hello"
	// which we want to translate to a Go function that prints "Hello, World!"

	// Create a new function declaration
	f := &ast.FuncDecl{
		Name: ast.NewIdent("hello"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: ast.NewIdent("fmt.Println"),
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"Hello, World!\"",
							},
						},
					},
				},
			},
		},
	}

	// Use the printer package to output the AST as Go code
	fs := token.NewFileSet()
	printer.Fprint(os.Stdout, fs, f)
}
