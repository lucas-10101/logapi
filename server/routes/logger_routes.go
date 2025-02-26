package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas-10101/logapi/data/models"
	"github.com/lucas-10101/logapi/server/routes/parser"
	"github.com/lucas-10101/logapi/services"
)

// create logger functionality routes
func loggerRoutes(router *mux.Router) {
	group := router.PathPrefix("/logger")

	createNewLogEndpoint(group)
	createReadLogsEndpoint(group)
}

// Add route for creating new logs in application
func createNewLogEndpoint(group *mux.Route) *mux.Route {
	group.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		payload, err := parser.ReadLogObjectFromRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())

			fmt.Println("")
		}

		err = services.SaveLog(*payload)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}).Methods(http.MethodPost)

	return group
}

// Route for reading logs
func createReadLogsEndpoint(group *mux.Route) *mux.Route {
	group.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageRequest := models.PageRequest{}
		pageRequest.LoadFromUrlQuery(r.URL.RawQuery)

	}).Methods(http.MethodGet)

	return group
}
