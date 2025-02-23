package clients

import (
	"github.com/lucas-10101/logapi/data/model"
)

var (
	defaultDriver = "mongodb"
	clients       map[string]*databaseClient
)

// Compatible client definition for use in api
type databaseClient interface {
	InsertOne(data *model.Document)
}

func GetClient(driverName string) *databaseClient {

	client, loaded := clients[defaultDriver]

	if !loaded {
		loadClient(defaultDriver)
	}

	return client
}

func GetDefaultClient() *databaseClient {
	return GetClient(defaultDriver)
}

func loadClient(driverName string) {

	switch driverName {
	case "mongodb":
		fallthrough
	default:
	}
}
