package clients

import (
	"github.com/lucas-10101/logapi/data/model"
)

var (
	defaultDriver = "mongodb"
	clients       = map[string]databaseClient{}
)

// Compatible client definition for use in api
type databaseClient interface {
	InsertOne(model.Document)
}

// Gets or loads an client into managed connections
func GetClient(driverName string) databaseClient {

	client, loaded := clients[defaultDriver]

	if !loaded {
		loadClient(defaultDriver)
	}

	return client
}

// Client loading function
func loadClient(driverName string) databaseClient {

	var commonClient databaseClient = nil
	switch driverName {
	case "mongodb":
		fallthrough
	default:
		client := mongoDBClient{}
		client.Connect()
		commonClient = client
	}

	clients[defaultDriver] = commonClient
	return commonClient
}
