package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Program struct {
	Body []Node `json:"body"`
}

type Node struct {
	Type       string   `json:"type"`
	Name       string   `json:"name,omitempty"`
	Fields     []Field  `json:"fields,omitempty"`
	Receiver   string   `json:"receiver,omitempty"`
	Parameters []Field  `json:"parameters,omitempty"`
	Returns    []string `json:"returns,omitempty"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	jsonFile, err := os.Open("generated_ast.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var program Program
	json.Unmarshal(byteValue, &program)

	for _, node := range program.Body {
		switch node.Type {
		case "StructDeclaration":
			generateStruct(node)
		case "MethodDeclaration":
			generateMethod(node)
		}
	}
}

func generateStruct(node Node) {
	fmt.Printf("type %s struct {\n", node.Name)
	for _, field := range node.Fields {
		fmt.Printf("    %s %s\n", field.Name, field.Type)
	}
	fmt.Println("}\n")
}

func generateMethod(node Node) {
	receiver := strings.ToLower(node.Receiver[:1])
	fmt.Printf("func (%s *%s) %s(%s) (%s) {\n",
		receiver, node.Receiver, node.Name, formatParameters(node.Parameters), strings.Join(node.Returns, ", "))
	fmt.Println("    // TODO: Implement method logic")
	fmt.Println("}\n")
}

func formatParameters(params []Field) string {
	var result []string
	for _, p := range params {
		result = append(result, fmt.Sprintf("%s %s", p.Name, p.Type))
	}
	return strings.Join(result, ", ")
}
