package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize server configuration
	serverConfig, err := initServerConfig()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v\n", err)
	}

	// Initialize routes
	err = initRoutes()
	if err != nil {
		log.Fatalf("Failed to initialize routes: %v\n", err)
	}

	// Start the server
	startServer(serverConfig)
}

// initServerConfig parses the server configuration from the DSL file
func initServerConfig() (ServerConfig, error) {
	serverConfig, err := parseDSLConfig("app/Server.aur")
	if err != nil {
		return ServerConfig{}, fmt.Errorf("error parsing server configuration: %w", err)
	}
	return serverConfig, nil
}

// initRoutes parses route and middleware configurations and sets up routes
func initRoutes() error {
	routeConfigs, err := parseRouteConfig("app/Routes.aur")
	if err != nil {
		return fmt.Errorf("error parsing route configuration: %w", err)
	}

	middlewareConfigs, err := parseMiddlewareConfig("app/Middlewares.aur")
	if err != nil {
		return fmt.Errorf("error parsing middleware configuration: %w", err)
	}

	for _, route := range routeConfigs {
		setupRoute(route, middlewareConfigs)
	}

	return nil
}

// setupRoute sets up a single route with middleware
func setupRoute(route RouteConfig, middlewareConfigs []MiddlewareConfig) {
	handler := createHandler(route)
	handler = applyMiddleware(handler, middlewareConfigs)

	log.Printf("Adding route %s %s\n", route.Method, route.Path)
	http.HandleFunc(route.Path, handler)
}

// createHandler creates a basic HTTP handler for a route
func createHandler(route RouteConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, route.Response)
	}
}

// startServer starts the HTTP server with the provided configuration
func startServer(config ServerConfig) {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Starting server at http://%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
