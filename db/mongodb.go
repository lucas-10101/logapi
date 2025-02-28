package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
)

func GetMongodbClient() *mongo.Client {
	if client != nil {
		return client
	}

	cli, err := mongo.Connect(context.TODO(), GetMongodbClientOptions())
	if err != nil {
		panic("database connection unavailable")
	}

	if err = cli.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic("database connection unreachable")
	}

	client = cli
	return client
}

func GetMongodbClientOptions() *options.ClientOptions {
	clientOptions := &options.ClientOptions{}
	clientOptions.ApplyURI("mongodb://lucas:lucas@mongodb:27017") // TODO use dinamic configuration

	return clientOptions
}

func GetDefaultCollectionClient() *mongo.Collection {
	return GetMongodbClient().Database("logsdb").Collection("applicationlogs")
}
