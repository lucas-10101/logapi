package handlers

import (
	"net/http"

	"github.com/lucas-10101/logapi/models"
	"github.com/lucas-10101/logapi/services"
)

func HandlerCreateNewLog(writter http.ResponseWriter, request *http.Request) {

	var logModel models.LogModel
	if err := services.BodyParser(request, &logModel); err != nil {
		writter.WriteHeader(err.StatusCode())
		writter.Write(err.ErrorBytes())
		return
	}

	if !logModel.Validate() {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("body is not valid"))
		return
	}
	service := services.LogService{}
	if err := service.Save(&logModel); err != nil {
		writter.WriteHeader(err.StatusCode())
		writter.Write(err.ErrorBytes())
		return
	}

	writter.WriteHeader(http.StatusAccepted)
}
