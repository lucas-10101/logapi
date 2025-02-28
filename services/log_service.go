package services

import (
	"context"
	"net/http"

	"github.com/lucas-10101/logapi/db"
	"github.com/lucas-10101/logapi/models"
	"github.com/lucas-10101/logapi/web/webutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogService struct{}

func (service *LogService) Save(model *models.LogModel) webutils.HttpError {
	collection := db.GetDefaultCollectionClient()
	_, err := collection.InsertOne(context.TODO(), model.ToBsonD())
	if err != nil {
		return webutils.NewHttpError(http.StatusInternalServerError, "cant save log data")
	}
	return nil
}

func (service *LogService) ReadPaginated(pagination models.Pagination) ([]*models.LogModel, webutils.HttpError) {
	collection := db.GetDefaultCollectionClient()

	findOptions := options.Find()
	findOptions.SetSkip(pagination.PageSize * pagination.PageNumber)
	findOptions.SetLimit(pagination.PageSize)
	filter := bson.D{}
	context := context.TODO()

	cursor, err := collection.Find(context, filter, findOptions)

	if err != nil {
		return nil, webutils.NewHttpError(http.StatusInternalServerError, "cant read from logs")
	}
	defer cursor.Close(context)
	results := []*models.LogModel{}

	for cursor.Next(context) {
		element := &models.LogModel{}
		err = cursor.Decode(element)

		if err != nil {
			return nil, webutils.NewHttpError(http.StatusInternalServerError, "cant read log pages")
		}
		results = append(results, element)
	}
	return results, nil
}

func (service *LogService) SetupDatabaseCollection() ([]bson.D, webutils.HttpError) {
	panic("TODO")
}
