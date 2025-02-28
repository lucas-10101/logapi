package dbclients

import (
	"context"
	"net/http"
	"reflect"

	"github.com/lucas-10101/logapi/data/models"
	"github.com/lucas-10101/logapi/server/http_utils"
	"github.com/lucas-10101/logapi/settings"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type mongoDBClient struct {
	rawClient *mongo.Client
}

// Insert document on deafult database/collection
func (client mongoDBClient) InsertOne(document models.Document) (any, error) {
	return client.rawClient.
		Database(settings.GetApplicationProperties().GetDatabaseProperties().GetDefaultDatabase()).
		Collection(settings.GetApplicationProperties().GetDatabaseProperties().GetDefaultCollection()).
		InsertOne(context.TODO(), convertToBsonD(document))
}

// Read all data paginated and ordered by timestamp (timeseries collection expected with timestamp field)
func (client mongoDBClient) Read(paginationData models.PageRequest, filter interface{}) ([]models.Document, error) {

	ctx := context.TODO()

	findOptions := options.Find()
	findOptions.SetSort(
		models.Document{
			{Field: "timestamp", Value: 1},
		},
	).SetSkip(paginationData.PageNumber).SetLimit(paginationData.PageSize)

	cursor, err := client.rawClient.
		Database(settings.GetApplicationProperties().GetDatabaseProperties().GetDefaultDatabase()).
		Collection(settings.GetApplicationProperties().GetDatabaseProperties().GetDefaultCollection()).
		Find(ctx, filter, findOptions)

	if err != nil {
		return nil, http_utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	dataSet := []models.Document{}
	for cursor.Next(ctx) {

		data := models.Document{}
		err = cursor.Decode(&data)

		if err != nil {
			return nil, http_utils.NewHttpError(http.StatusInternalServerError, err.Error())
		}

		dataSet = append(dataSet, data)
	}

	return dataSet, nil
}

// starts the managed connection
func (client *mongoDBClient) Connect() {
	if client.rawClient != nil {
		return
	}

	connection, err := mongo.Connect(options.Client().ApplyURI("mongodb://lucas:lucas@mongodb:27017"))
	if err != nil {
		panic("Cant create connection, reason: " + err.Error())
	}

	err = connection.Ping(context.TODO(), readpref.Nearest())
	if err != nil {
		panic("no server response, reason: " + err.Error())
	}
	client.rawClient = connection
}

// convert application document to mongo data type bson.D
func convertToBsonD(document models.Document) bson.D {
	bsonData := make([]bson.E, len(document))
	for i, e := range document {

		bsonData[i] = bson.E{
			Key:   e.Field,
			Value: convertToBsonDValueCollector(e.Value),
		}
	}

	return bsonData
}

// convert application document value to mongo data type bson.E if needed
func convertToBsonDValueCollector(actual interface{}) any {
	value := actual
	if reflect.TypeOf(value).Name() == "Document" {
		asserted, ok := value.(models.Document)
		if ok {
			value = convertToBsonD(asserted)
		}
	}

	return value
}
