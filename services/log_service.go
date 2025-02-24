package services

import (
	"github.com/lucas-10101/logapi/data/dbclients"
	"github.com/lucas-10101/logapi/data/models"
)

func SaveLog(log models.LogDocument) error {
	_, error := dbclients.GetDefaultClient().InsertOne(log.ToDocument())
	return error
}

func ReadLogs()
