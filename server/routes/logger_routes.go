package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas-10101/logapi/data/models"
	"github.com/lucas-10101/logapi/server/http_utils"
	"github.com/lucas-10101/logapi/services"
)

// create logger functionality routes
func loggerRoutes(router *mux.Router) {
	group := router.PathPrefix("/logger").Subrouter()

	createReadLogsEndpoint(group)
	createNewLogEndpoint(group)
}

// Add route for creating new logs in application
func createNewLogEndpoint(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

		var data models.LogDocument
		err := http_utils.RequestBodyParser(r, &data)

		if err != nil {
			err.Send(w)
			return
		}

		err = services.SaveLog(data)
		if err != nil {
			err.Send(w)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}).Methods(http.MethodPost)
}

// Route for reading logs
func createReadLogsEndpoint(router *mux.Router) {
	router.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		pageRequest := models.PageRequest{}
		pageRequest.LoadFromUrlQuery(r.URL.RawQuery)

		data, err := services.ReadLogs(pageRequest)
		if err != nil {
			err.Send(w)
			return
		}

		err = http_utils.ResponseBodyWritter(http_utils.MimeTypeApplicationJson, w, data)
		if err != nil {
			err.Send(w)
		}
	}).Methods(http.MethodGet)
}
