package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Program struct {
	Type string `json:"type"`
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
	dsl := `
Controller.UserProfile ->
    GetProfile(userId: Integer) -> { GetUser(userId), GetPostStats(userId) }
    UpdateAge(userId: Integer, newAge: Integer) -> ValidateAge(newAge) ? UpdateUserAge(userId, newAge) : Error("Invalid age")
    GetPostStats(userId: Integer) -> { Count(Post.findAllByUserId(userId)), avgLength(Post.findAllByUserId(userId)) }
    GetUser(userId: Integer) -> User.find(userId) or Error("Not found")
    ValidateAge(age: Integer) -> age in range(0, 150)
    UpdateUserAge(userId: Integer, age: Integer) -> User.find(userId).update(age: age)
    AvgLength(posts: Post[]) -> posts.empty() ? 0 : SumLengths(posts) / posts.count()
`

	program := parseDSL(dsl)
	jsonString, _ := json.MarshalIndent(program, "", "    ")
	fmt.Println(string(jsonString))
}

func parseDSL(dsl string) Program {
	var program Program
	program.Type = "Program"

	// Predefined structs
	program.Body = append(program.Body, Node{
		Type:   "StructDeclaration",
		Name:   "UserProfileController",
		Fields: []Field{},
	}, Node{
		Type: "StructDeclaration",
		Name: "User",
		Fields: []Field{
			{"ID", "int"},
			{"Age", "int"},
		},
	}, Node{
		Type: "StructDeclaration",
		Name: "Post",
		Fields: []Field{
			{"UserID", "int"},
		},
	})

	lines := strings.Split(dsl, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "Controller.UserProfile ->") {
			continue
		}

		node := parseLine(line)
		program.Body = append(program.Body, node)
	}

	return program
}

func parseLine(line string) Node {
	parts := strings.Split(line, "->")
	methodDef := strings.TrimSpace(parts[0])
	methodNameAndParams := strings.Split(methodDef, "(")
	methodName := strings.TrimSpace(methodNameAndParams[0])
	paramsPart := strings.TrimSuffix(methodNameAndParams[1], ")")
	params := parseParams(paramsPart)

	returns := inferReturnTypes(line)

	return Node{
		Type:       "MethodDeclaration",
		Receiver:   "UserProfileController",
		Name:       methodName,
		Parameters: params,
		Returns:    returns,
	}
}

func parseParams(paramsPart string) []Field {
	var params []Field
	if paramsPart == "" {
		return params
	}

	paramPairs := strings.Split(paramsPart, ",")
	for _, p := range paramPairs {
		parts := strings.Split(strings.TrimSpace(p), ":")
		paramName := strings.TrimSpace(parts[0])
		paramType := mapType(strings.TrimSpace(parts[1]))
		params = append(params, Field{Name: paramName, Type: paramType})
	}

	return params
}

func mapType(paramType string) string {
	switch paramType {
	case "Integer":
		return "int"
	case "Post[]":
		return "[]Post"
	default:
		return paramType
	}
}

func inferReturnTypes(line string) []string {
	var returns []string
	if strings.Contains(line, "Error") {
		returns = append(returns, "error")
	}

	// Add more rules as needed
	// ...

	return returns
}
