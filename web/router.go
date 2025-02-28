package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lucas-10101/logapi/web/handlers"
)

// Make router with route matching
func Router() *http.ServeMux {
	router := http.NewServeMux()
	routes := makeRouteMappings()

	router.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		identifier := getRouteIdentifier(request.Method, request.URL.Path)
		if handler, exists := routes[identifier]; exists {
			handler(writter, request)
		} else {
			writter.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	return router
}

func addRouteMapping(httpMethod string, requestPath string, routes map[string]http.HandlerFunc, handler http.HandlerFunc) {
	identifier := getRouteIdentifier(httpMethod, requestPath)
	if _, exists := routes[identifier]; exists {
		panic(fmt.Sprintf("route %s with method %s already registerd on router", requestPath, httpMethod))
	}
	routes[identifier] = handler
}

// Create an route identifier like GET:/some/path
func getRouteIdentifier(httpMethod string, requestPath string) string {
	return strings.ToUpper(httpMethod) + ":" + strings.ToUpper(requestPath)
}

// Actual route building function
func makeRouteMappings() map[string]http.HandlerFunc {
	routes := map[string]http.HandlerFunc{}

	addRouteMapping(http.MethodGet, "/logs", routes, handlers.HandlerCreateNewLog)

	return routes
}
