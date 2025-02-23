package routes

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucas-10101/logapi/data/clients"
	"github.com/lucas-10101/logapi/data/model"
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

		log := model.LogDocument{
			Timestamp: time.Now(),
			ApplicationSource: model.LogDocumentApplicationSource{
				ApplicationName:  "LogApi",
				ApplicationRoute: request.URL.RawPath,
			},
			RequestInfo: model.LogDocumentRequestInfo{
				RequestMethod:     request.Method,
				RequestStatusCode: 204,
				RequestDuration:   nil,
			},
			ErrorInfo: model.LogDocumentErrorInfo{},
		}

		clients.GetClient("mongodb").InsertOne(log.ToDocument())
	}).Methods(http.MethodPost)
}
