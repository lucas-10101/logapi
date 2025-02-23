package clients

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/lucas-10101/logapi/data/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type mongoDBClient struct {
	rawClient *mongo.Client
}

// Insert document on deafult database/collection
func (client mongoDBClient) InsertOne(document model.Document) {

	client.rawClient.Database("teste").Collection("teste").InsertOne(context.Background(), convertToBsonD(document))
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
func convertToBsonD(document model.Document) bson.D {
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
		asserted, ok := value.(model.Document)
		if ok {
			value = convertToBsonD(asserted)
		}
	}

	return value
}
