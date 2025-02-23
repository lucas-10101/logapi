package settings

type ApplicationProperties struct {
	defaultNoSQLProvider string
	defaultIanaTimeZone  string
	defaultDatabase      string
	defaultCollection    string
}

func (properties ApplicationProperties) GetDefaultNoSQLProvider() string {
	return properties.defaultNoSQLProvider
}

func (properties ApplicationProperties) GetDefaultTimeZone() string {
	return properties.defaultIanaTimeZone
}

func (properties ApplicationProperties) GetDefaultDatabase() string {
	return properties.defaultDatabase
}

func (properties ApplicationProperties) GetDefaultCollection() string {
	return properties.defaultCollection
}
