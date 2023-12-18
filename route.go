package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type RouteConfig struct {
	Method   string
	Path     string
	Response string
}

func parseRouteConfig(filePath string) ([]RouteConfig, error) {
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

		if strings.HasPrefix(line, "Route") {
			var err error
			currentRoute, err = parseRoute(line)
			if err != nil {
				return nil, err
			}
		} else if currentRoute != nil && strings.HasPrefix(line, "Response:") {
			response := extractResponse(line)
			currentRoute.Response = response
			routes = append(routes, *currentRoute)
			currentRoute = nil // Reset for the next route
		}
	}

	return routes, scanner.Err()
}

func parseRoute(line string) (*RouteConfig, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil, errors.New("invalid Route format")
	}

	method := strings.ToUpper(parts[1])
	path := parts[2]

	return &RouteConfig{Method: method, Path: path}, nil
}

func extractResponse(line string) string {
	// Extracts the response part from the line
	parts := strings.SplitN(line, "\"", 3)
	if len(parts) >= 2 {
		return parts[1] // The response text is the second part
	}
	return ""
}
