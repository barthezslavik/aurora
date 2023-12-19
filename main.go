package main

import (
	"aurora/lib"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize server configuration
	serverConfig, err := InitServerConfig()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v\n", err)
	}

	// Initialize routes
	err = InitRoutes()
	if err != nil {
		log.Fatalf("Failed to initialize routes: %v\n", err)
	}

	// Start the server
	StartServer(serverConfig)
}

// initServerConfig parses the server configuration from the DSL file
func InitServerConfig() (lib.ServerConfig, error) {
	serverConfig, err := lib.ParseDSLConfig("app/Server.aur")
	if err != nil {
		return lib.ServerConfig{}, fmt.Errorf("error parsing server configuration: %w", err)
	}
	return serverConfig, nil
}

// initRoutes parses route and middleware configurations and sets up routes
func InitRoutes() error {
	routeConfigs, err := lib.ParseRouteConfig("app/Routes.aur")
	if err != nil {
		return fmt.Errorf("error parsing route configuration: %w", err)
	}

	middlewareConfigs, err := lib.ParseMiddlewareConfig("app/Middlewares.aur")
	if err != nil {
		return fmt.Errorf("error parsing middleware configuration: %w", err)
	}

	for _, route := range routeConfigs {
		SetupRoute(route, middlewareConfigs)
	}

	return nil
}

// setupRoute sets up a single route with middleware
func SetupRoute(route lib.RouteConfig, middlewareConfigs []lib.MiddlewareConfig) {
	handler := CreateHandler(route)
	handler = lib.ApplyMiddleware(handler, middlewareConfigs)

	log.Printf("Adding route %s %s\n", route.Method, route.Path)
	http.HandleFunc(route.Path, handler)
}

// createHandler creates a basic HTTP handler for a route
func CreateHandler(route lib.RouteConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, route.Response)
	}
}

// startServer starts the HTTP server with the provided configuration
func StartServer(config lib.ServerConfig) {
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Starting server at http://%s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
