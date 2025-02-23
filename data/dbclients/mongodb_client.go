package dbclients

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/lucas-10101/logapi/data/models"
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
		InsertOne(context.Background(), convertToBsonD(document))
}

// starts the managed connection
func (client *mongoDBClient) Connect() {
	if client.rawClient != nil {
		return
	}

	connection, err := mongo.Connect(options.Client().ApplyURI("mongodb://lucas:lucas@mongodb:27017"))
	if err != nil {
		fmt.Println("Cant connect, reason: " + err.Error())
	}

	timeoutContext, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = connection.Ping(timeoutContext, readpref.Nearest())
	if err != nil {
		fmt.Println("No server response, reason: " + err.Error())
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
