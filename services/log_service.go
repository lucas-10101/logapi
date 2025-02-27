package services

import (
	"net/http"

	"github.com/lucas-10101/logapi/data/dbclients"
	"github.com/lucas-10101/logapi/data/models"
	"github.com/lucas-10101/logapi/server/http_utils"
)

func SaveLog(log models.LogDocument) http_utils.HttpError {
	_, err := dbclients.GetDefaultClient().InsertOne(log.ToDocument())
	if err != nil {
		return http_utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func ReadLogs(pageRequest models.PageRequest) ([]models.Document, http_utils.HttpError) {
	results, err := dbclients.GetDefaultClient().Read(pageRequest, models.Document{})
	if err != nil {
		return nil, http_utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	return results, nil
}
