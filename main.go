package main

import (
	"fmt"
	"net/http"
)

func main() {
	serverConfig, err := parseDSLConfig("server_config.dsl")
	if err != nil {
		fmt.Println("Error parsing server configuration:", err)
		return
	}

	routeConfigs, err := parseRouteConfig("Routes.dsl")
	if err != nil {
		fmt.Println("Error parsing route configuration:", err)
		return
	}

	for _, route := range routeConfigs {
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, route.Response)
		})
	}

	fmt.Printf("Starting server at %s:%d\n", serverConfig.Host, serverConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
