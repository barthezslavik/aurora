package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type RouteConfig struct {
	Method   string
	Path     string
	Response string
}

type DSLLine struct {
	Directive string
	Arguments []string
}

func parseRouteConfig(filePath string) ([]RouteConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var routes []RouteConfig
	var currentRoute *RouteConfig

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		dslLine, err := parseDSLLine(line)
		if err != nil {
			fmt.Println("Warning: skipping line due to parse error:", err)
			continue
		}

		switch dslLine.Directive {
		case "Route":
			currentRoute, err = interpretRoute(dslLine)
			if err != nil {
				fmt.Println("Warning: skipping route due to interpretation error:", err)
				continue
			}
		case "Response:":
			if currentRoute != nil {
				currentRoute.Response = interpretResponse(dslLine)
				routes = append(routes, *currentRoute)
				currentRoute = nil
			}
		}
	}

	return routes, scanner.Err()
}

func parseDSLLine(line string) (DSLLine, error) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return DSLLine{}, errors.New("empty line")
	}

	directive := parts[0]
	arguments := parts[1:]

	return DSLLine{Directive: directive, Arguments: arguments}, nil
}

func interpretRoute(dslLine DSLLine) (*RouteConfig, error) {
	if len(dslLine.Arguments) < 2 {
		return nil, errors.New("insufficient arguments for Route")
	}

	method := strings.ToUpper(dslLine.Arguments[0])
	path := dslLine.Arguments[1]

	return &RouteConfig{Method: method, Path: path}, nil
}

func interpretResponse(dslLine DSLLine) string {
	if len(dslLine.Arguments) > 0 {
		// Join the arguments to form the response string
		response := strings.Join(dslLine.Arguments, " ")

		// Remove surrounding quotes and trailing semicolon
		response = strings.Trim(response, "\";")
		return response
	}
	return ""
}
