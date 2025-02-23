package settings

import (
	"fmt"
	"time"
)

// Configure any dependencies in api
func Configure() {
	configureApplicationProperties() // must be first call to gather all required configuration
	configureTimeZone()
}

func configureTimeZone() {
	location, err := time.LoadLocation(GetApplicationProperties().GetServerProperties().GetDefaultTimeZone())

	if err != nil {
		panic(fmt.Sprintf("Cant load location: %s\n", GetApplicationProperties().GetServerProperties().GetDefaultTimeZone()))
	}
	time.Local = location
}

// load application required properties
func configureApplicationProperties() {
	properties.databaseProperties.defaultDriver = "mongodb"
	properties.databaseProperties.defaultDatabase = "teste"
	properties.databaseProperties.defaultCollection = "teste"
	properties.serverProperties.serverHost = "127.0.0.1"
	properties.serverProperties.serverPort = 2525
	properties.serverProperties.defaultTimeZone = "Etc/GMT+0"
}
