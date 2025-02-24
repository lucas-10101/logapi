package settings

//////////////////////////////////////////////////////////////////////////////////////////////

var (
	properties = applicationProperties{
		databaseProperties: databaseProperties{
			driverProperties: map[string]string{},
		},
		serverProperties: serverProperties{},
	}
)

func GetApplicationProperties() applicationProperties {
	return properties
}

//////////////////////////////////////////////////////////////////////////////////////////////

type OutputMode string

const (
	OUTPUT_FILE    = OutputMode("FILE")
	OUTPUT_CONSOLE = OutputMode("CONSOLE")
	OUTPUT_NONE    = OutputMode("NONE")
)

//////////////////////////////////////////////////////////////////////////////////////////////

type applicationProperties struct {
	stdInputOutputMode OutputMode

	databaseProperties databaseProperties
	serverProperties   serverProperties
	requestProperties  requestProperties
}

func (properties applicationProperties) GetStdInputOutputMode() OutputMode {
	return properties.stdInputOutputMode
}

func (properties applicationProperties) GetDatabaseProperties() databaseProperties {
	return properties.databaseProperties
}

func (properties applicationProperties) GetServerProperties() serverProperties {
	return properties.serverProperties
}

func (properties applicationProperties) GetRequestProperties() requestProperties {
	return properties.requestProperties
}

//////////////////////////////////////////////////////////////////////////////////////////////

type databaseProperties struct {
	defaultDatabase   string
	defaultCollection string
	defaultDriver     string
	driverProperties  map[string]string
}

func (propeties databaseProperties) GetDefaultDatabase() string {
	return propeties.defaultDatabase
}

func (propeties databaseProperties) GetDefaultCollection() string {
	return propeties.defaultCollection
}

func (propeties databaseProperties) GetDefaultDriver() string {
	return propeties.defaultDriver
}

func (propeties databaseProperties) GetDriverProperties() map[string]string {
	return propeties.driverProperties
}

//////////////////////////////////////////////////////////////////////////////////////////////

type serverProperties struct {
	serverHost      string
	serverPort      int
	defaultTimeZone string
}

func (properties serverProperties) GetServerHost() string {
	return properties.serverHost
}

func (properties serverProperties) GetServerPort() int {
	return properties.serverPort
}

func (properties serverProperties) GetDefaultTimeZone() string {
	return properties.defaultTimeZone
}

//////////////////////////////////////////////////////////////////////////////////////////////

type requestProperties struct {
	maxPaginationSize int64
}

func (properties requestProperties) GetMaxPaginationSize() int64 {
	return properties.maxPaginationSize
}

//////////////////////////////////////////////////////////////////////////////////////////////
