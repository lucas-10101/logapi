package services

import (
	"context"
	"net/http"

	"github.com/lucas-10101/logapi/db"
	"github.com/lucas-10101/logapi/models"
	"github.com/lucas-10101/logapi/web/webutils"
	"go.mongodb.org/mongo-driver/bson"
)

type LogService struct{}

func (service *LogService) Save(model *models.LogModel) webutils.HttpError {
	collection := db.GetMongodbClient().Database("logsdb").Collection("applicationlogs") // TODO use application validation rules
	_, err := collection.InsertOne(context.TODO(), model.ToBsonD())
	if err != nil {
		return webutils.NewHttpError(http.StatusInternalServerError, "cant save log data")
	}
	return nil
}

func (service *LogService) ReadPaginated() ([]bson.D, webutils.HttpError) {
	panic("TODO")
}

func (service *LogService) SetupDatabaseCollection() ([]bson.D, webutils.HttpError) {
	panic("TODO")
}
