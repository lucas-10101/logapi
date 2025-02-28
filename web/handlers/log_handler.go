package handlers

import (
	"net/http"

	"github.com/lucas-10101/logapi/models"
	"github.com/lucas-10101/logapi/services"
)

func HandlerCreateNewLog(writter http.ResponseWriter, request *http.Request) {

	var logModel models.LogModel
	if err := services.BodyParser(request, &logModel); err != nil {
		err.Write(writter)
		return
	}

	if !logModel.Validate() {
		writter.WriteHeader(http.StatusBadRequest)
		writter.Write([]byte("body is not valid"))
		return
	}
	service := services.LogService{}
	if err := service.Save(&logModel); err != nil {
		err.Write(writter)
		return
	}

	writter.WriteHeader(http.StatusAccepted)
}

func HandlerReadLogs(writter http.ResponseWriter, request *http.Request) {
	paginationData, err := models.GetPaginationFromUrl(request.URL)
	if err != nil {
		err.Write(writter)
		return
	}

	service := services.LogService{}
	results, err := service.ReadPaginated(*paginationData)

	if err != nil {
		err.Write(writter)
		return
	}

	if err = services.BodyWritter(writter, request, results); err != nil {
		err.Write(writter)
	}
}
