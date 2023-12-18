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

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Route") {
			route, err := parseRoute(line)
			if err != nil {
				return nil, err
			}
			routes = append(routes, route)
		}
	}

	return routes, scanner.Err()
}

func parseRoute(line string) (RouteConfig, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return RouteConfig{}, errors.New("invalid Route format")
	}

	method := strings.ToUpper(parts[1])
	path := parts[2]
	response := strings.Join(parts[3:], " ")

	return RouteConfig{Method: method, Path: path, Response: response}, nil
}
