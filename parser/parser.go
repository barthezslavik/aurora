package parser

import (
	"strings"
)

func ParseControllerMethod(dsl string) (ControllerMethod, error) {
	var method ControllerMethod
	lines := strings.Split(dsl, "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Name":
			method.Name = value
		case "Path":
			method.Path = value
		case "Method":
			method.Method = value
		case "Action":
			method.Action = value
		}
	}

	return method, nil
}
