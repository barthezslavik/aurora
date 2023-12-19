// parser/types.go
package parser

// ControllerMethod represents a method definition in your DSL.
// It captures the details of a controller method such as its name, path, HTTP method, and action.
type ControllerMethod struct {
	Name   string // Name of the method
	Path   string // URL path the method responds to
	Method string // HTTP method (GET, POST, etc.)
	Action string // The action code to be executed
}

// You can extend this file with more types as needed for your DSL.
// For example, if you have a part of your DSL that defines models or database interactions,
// you might define corresponding types here.
