package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas-10101/logapi/data/models"
)

func loggerRoutes(router *mux.Router) {
	group := router.PathPrefix("/logger")

	group.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(400)
			return
		}

		var payload models.LogDocument
		err = json.Unmarshal(data, &payload)

		if err != nil {
			w.WriteHeader(400)
			return
		}

	}).Methods(http.MethodGet)
}
