package main

import (
	"fmt"
	"net/http"
)

func main() {
	serverConfig, err := parseDSLConfig("app/Server.aur")
	if err != nil {
		fmt.Println("Error parsing server configuration:", err)
		return
	}

	routeConfigs, err := parseRouteConfig("app/Routes.aur")
	if err != nil {
		fmt.Println("Error parsing route configuration:", err)
		return
	}

	// Parse middleware configurations from the DSL file
	middlewareConfigs, err := parseMiddlewareConfig("app/Middlewares.aur")
	if err != nil {
		fmt.Println("Error parsing middleware configuration:", err)
		return
	}

	for _, route := range routeConfigs {
		handler := func(w http.ResponseWriter, r *http.Request) {
			// This is a simplified handler. You might want to add more complex handling here.
			fmt.Fprint(w, route.Response)
		}

		// Wrap the handler with middleware
		handler = applyMiddleware(handler, middlewareConfigs)

		fmt.Printf("Adding route %s %s\n", route.Method, route.Path)
		http.HandleFunc(route.Path, handler)
	}

	fmt.Printf("Starting server at http://%s:%d\n", serverConfig.Host, serverConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
