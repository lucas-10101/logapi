package settings

//////////////////////////////////////////////////////////////////////////////////////////////

var (
	properties = ApplicationProperties{
		databaseProperties: DatabaseProperties{
			driverProperties: map[string]string{},
		},
		serverProperties: ServerProperties{},
	}
)

func GetApplicationProperties() ApplicationProperties {
	return properties
}

//////////////////////////////////////////////////////////////////////////////////////////////

type ApplicationProperties struct {
	databaseProperties DatabaseProperties
	serverProperties   ServerProperties
}

func (properties ApplicationProperties) GetDatabaseProperties() DatabaseProperties {
	return properties.databaseProperties
}

func (properties ApplicationProperties) GetServerProperties() ServerProperties {
	return properties.serverProperties

}

//////////////////////////////////////////////////////////////////////////////////////////////

type DatabaseProperties struct {
	defaultDatabase   string
	defaultCollection string
	defaultDriver     string
	driverProperties  map[string]string
}

func (propeties DatabaseProperties) GetDefaultDatabase() string {
	return propeties.defaultDatabase
}

func (propeties DatabaseProperties) GetDefaultCollection() string {
	return propeties.defaultCollection
}

func (propeties DatabaseProperties) GetDefaultDriver() string {
	return propeties.defaultDriver
}

func (propeties DatabaseProperties) GetDriverProperties() map[string]string {
	return propeties.driverProperties
}

//////////////////////////////////////////////////////////////////////////////////////////////

type ServerProperties struct {
	serverHost      string
	serverPort      int
	defaultTimeZone string
}

func (properties ServerProperties) GetServerHost() string {
	return properties.serverHost
}

func (properties ServerProperties) GetServerPort() int {
	return properties.serverPort
}

func (properties ServerProperties) GetDefaultTimeZone() string {
	return properties.defaultTimeZone
}

//////////////////////////////////////////////////////////////////////////////////////////////
