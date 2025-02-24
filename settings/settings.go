package settings

import (
	"fmt"
	"os"
	"time"
)

// Configure any dependencies in api
func Configure() {
	configureApplicationProperties() // must be first call to gather all required configuration
	configureTimeZone()
	configureStdInputOutputMode()
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
	properties.stdInputOutputMode = OUTPUT_NONE
	properties.databaseProperties.defaultDriver = "mongodb"
	properties.databaseProperties.defaultDatabase = "teste"
	properties.databaseProperties.defaultCollection = "teste"
	properties.serverProperties.serverHost = "127.0.0.1"
	properties.serverProperties.serverPort = 2525
	properties.serverProperties.defaultTimeZone = "Etc/GMT+0"
}

// Disable stdInput and stdOutput from whole application, not touching stdError
func configureStdInputOutputMode() {
	switch GetApplicationProperties().GetStdInputOutputMode() {
	case OUTPUT_FILE:
		panic("not implemented yet")
	case OUTPUT_NONE:
		os.Stdout = nil
	case OUTPUT_CONSOLE:
		fallthrough
	default:
	}
}
