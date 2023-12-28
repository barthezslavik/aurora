package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Node represents a JSON AST node
type Node struct {
	Type       string   `json:"type"`
	Name       string   `json:"name,omitempty"`
	Parameters []Field  `json:"parameters,omitempty"`
	Returns    []string `json:"returns,omitempty"`
}

// Field represents a field in a struct or a parameter in a method
type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	// Sample DSL input
	dsl := `
        Controller.UserProfile ->
            GetProfile(userId: Integer) -> { GetUser(userId), GetPostStats(userId) }
            UpdateAge(userId: Integer, newAge: Integer) -> ValidateAge(newAge) ? UpdateUserAge(userId, newAge) : Error("Invalid age")
            GetPostStats(userId: Integer) -> { Count(Post.findAllByUserId(userId)), avgLength(Post.findAllByUserId(userId)) }
            getUser(userId: Integer) -> User.find(userId) or Error("Not found")
            validateAge(age: Integer) -> age in range(0, 150)
            updateUserAge(userId: Integer, age: Integer) -> User.find(userId).update(age: age)
            avgLength(posts: Post[]) -> posts.empty() ? 0 : SumLengths(posts) / posts.count()
    `

	nodes, err := parseDSL(dsl)
	if err != nil {
		fmt.Println("Error parsing DSL:", err)
		return
	}

	jsonAst, err := json.MarshalIndent(nodes, "", "    ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}

	fmt.Println(string(jsonAst))
}

// parseDSL parses the given DSL string into a slice of Node structs
func parseDSL(dsl string) ([]Node, error) {
	var nodes []Node
	methodRegex := regexp.MustCompile(`(?m)^\s*(\w+)\(([\w\s:]+)\)\s*->\s*(\{?.*\}?)$`)
	paramRegex := regexp.MustCompile(`(\w+):\s*(\w+)`)

	lines := strings.Split(dsl, "\n")
	for _, line := range lines {
		matches := methodRegex.FindStringSubmatch(line)
		if len(matches) != 4 {
			continue
		}

		methodName := matches[1]
		paramsPart := matches[2]
		returnsPart := matches[3]

		var parameters []Field
		paramsMatches := paramRegex.FindAllStringSubmatch(paramsPart, -1)
		for _, p := range paramsMatches {
			parameters = append(parameters, Field{Name: p[1], Type: p[2]})
		}

		returns := []string{strings.TrimSpace(returnsPart)}

		nodes = append(nodes, Node{
			Type:       "MethodDeclaration",
			Name:       methodName,
			Parameters: parameters,
			Returns:    returns,
		})
	}

	return nodes, nil
}
