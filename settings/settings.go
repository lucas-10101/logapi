package settings

import (
	"fmt"
	"time"

	"github.com/lucas-10101/logapi/data/clients"
)

var (
	APPLICATION_PROPERTIES = ApplicationProperties{}
)

// Configure any dependencies in api
func Configure() {
	configureApplicationProperties() // must be first call to gather all required configuration
	configureTimeZone()
	configureDefaultConnection()
}

func configureTimeZone() {
	location, err := time.LoadLocation(APPLICATION_PROPERTIES.GetDefaultTimeZone())

	if err != nil {
		panic(fmt.Sprintf("Cant load location: %s\n", APPLICATION_PROPERTIES.GetDefaultTimeZone()))
	}
	time.Local = location
}

// load application required properties
func configureApplicationProperties() {
	APPLICATION_PROPERTIES.defaultNoSQLProvider = "mongodb"
	APPLICATION_PROPERTIES.defaultIanaTimeZone = "Etc/GMT+0"
	APPLICATION_PROPERTIES.defaultDatabase = "teste"
	APPLICATION_PROPERTIES.defaultCollection = "teste"
}

func configureDefaultConnection() {
	clients.GetClient(APPLICATION_PROPERTIES.GetDefaultNoSQLProvider())
}
