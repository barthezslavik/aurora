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

	for _, route := range routeConfigs {
		fmt.Printf("Adding route %s %s\n", route.Method, route.Path)
		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, route.Response)
		})
	}

	fmt.Printf("Starting server at http://%s:%d\n", serverConfig.Host, serverConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}