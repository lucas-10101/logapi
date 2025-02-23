package clients

import (
	"context"
	"fmt"
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

func (client *mongoDBClient) InsertOne(document model.Document) {

	bsonData := make([]bson.E, len(document))
	for i, e := range document {
		bsonData[i] = bson.E{
			Key:   e.Field,
			Value: e.Value,
		}
	}

	client.rawClient.Database("teste").Collection("teste").InsertOne(context.Background(), bsonData)
}

func newMongoClient() *mongo.Client {

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://lucas:lucas@mongodb:27017"))
	if err != nil {
		fmt.Println("Cant connect, reason: " + err.Error())
	}

	timeoutContext, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Ping(timeoutContext, readpref.Nearest())
	if err != nil {
		fmt.Println("No server response, reason: " + err.Error())
	}

	return client
}
