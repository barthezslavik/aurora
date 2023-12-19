package lib

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// MiddlewareConfig represents the configuration of a middleware
type MiddlewareConfig struct {
	Name    string
	Action  string
	Options map[string]string
}

// parseMiddlewareConfig parses the middleware configuration from a DSL file
func ParseMiddlewareConfig(filePath string) ([]MiddlewareConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var middlewares []MiddlewareConfig
	scanner := bufio.NewScanner(file)

	var currentMiddleware *MiddlewareConfig
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Middleware") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				fmt.Println("Warning: skipping middleware due to insufficient arguments")
				continue
			}
			currentMiddleware = &MiddlewareConfig{
				Name:    parts[1],
				Options: make(map[string]string),
			}
		} else if strings.HasPrefix(line, "Action:") {
			if currentMiddleware != nil {
				action := strings.TrimPrefix(line, "Action:")
				currentMiddleware.Action = strings.TrimSpace(action)
			}
		} else {
			keyVal := strings.SplitN(line, ":", 2)
			if len(keyVal) == 2 && currentMiddleware != nil {
				key := strings.TrimSpace(keyVal[0])
				value := strings.TrimSpace(keyVal[1])
				currentMiddleware.Options[key] = value
			}
		}

		if line == "}" && currentMiddleware != nil {
			middlewares = append(middlewares, *currentMiddleware)
			currentMiddleware = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return middlewares, nil
}

// applyMiddleware wraps an HTTP handler with the specified middlewares
func ApplyMiddleware(handler http.HandlerFunc, middlewares []MiddlewareConfig) http.HandlerFunc {
	for _, mw := range middlewares {
		switch mw.Action {
		case "LogRequest":
			handler = LogRequestMiddleware(handler)
		case "Authenticate":
			handler = AuthenticateMiddleware(handler, mw.Options["Key"])
		}
	}
	return handler
}

// logRequestMiddleware is a middleware that logs each request
func LogRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

// authenticateMiddleware is a middleware that handles authentication
func AuthenticateMiddleware(next http.HandlerFunc, key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authentication logic here...
		// For example, check some header against the 'key'
		authHeader := r.Header.Get("Authorization")
		if authHeader != key {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
