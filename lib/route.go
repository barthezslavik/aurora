package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RouteConfig struct {
	Method   string
	Path     string
	Response string
}

func ParseRouteConfig(filePath string) ([]RouteConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var routes []RouteConfig
	scanner := bufio.NewScanner(file)

	var currentRoute *RouteConfig
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		dslLine, err := ParseDSLLine(line)
		if err != nil {
			fmt.Println("Warning: skipping line due to parse error:", err)
			continue
		}

		switch dslLine.Directive {
		case "Route":
			currentRoute, err = InterpretRoute(dslLine)
			if err != nil {
				fmt.Println("Warning: skipping route due to interpretation error:", err)
				continue
			}
		case "Response:":
			if currentRoute != nil {
				currentRoute.Response = InterpretResponse(dslLine)
				routes = append(routes, *currentRoute)
				currentRoute = nil
			}
		}
	}

	return routes, scanner.Err()
}

func InterpretRoute(dslLine DSLLine) (*RouteConfig, error) {
	if len(dslLine.Arguments) < 2 {
		return nil, fmt.Errorf("insufficient arguments for Route")
	}

	method := strings.ToUpper(dslLine.Arguments[0])
	path := dslLine.Arguments[1]

	return &RouteConfig{Method: method, Path: path}, nil
}

func InterpretResponse(dslLine DSLLine) string {
	if len(dslLine.Arguments) > 0 {
		// Removes quotes and semicolon
		response := strings.Join(dslLine.Arguments, " ")
		response = strings.Trim(response, "\";")
		return response
	}
	return ""
}
