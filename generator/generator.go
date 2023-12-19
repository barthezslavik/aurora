package generator

import (
	"aurora/parser"
	"fmt"
)

func GenerateBuffaloCode(method parser.ControllerMethod) string {
	return fmt.Sprintf(`func (c App) %s(ctx buffalo.Context) error {
        %s
    }`, method.Name, method.Action)
}
