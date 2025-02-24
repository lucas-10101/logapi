package routes

import (
	"github.com/gorilla/mux"
)

func MakeRoutes() *mux.Router {
	router := mux.NewRouter()
	loggerRoutes(router)

	return router
}
