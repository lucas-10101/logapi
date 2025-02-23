package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRoutes() *mux.Router {
	router := mux.NewRouter()

	statusCheckRoutes(router)
	loggerRoutes(router)

	return router
}

// Status check route group to check if api is alive
func statusCheckRoutes(router *mux.Router) {
	healthGroup := router.PathPrefix("/status")

	healthGroup.PathPrefix("/hearthbeat").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	}).Methods(http.MethodGet)
}

// Logger api routes
func loggerRoutes(router *mux.Router) {
	loggerGroup := router.PathPrefix("/logger")

	loggerGroup.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	}).Methods(http.MethodPost)
}
