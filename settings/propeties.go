package settings

type ApplicationProperties struct {
	defaultNoSQLProvider string
	defaultIanaTimeZone  string
}

func (properties ApplicationProperties) GetDefaultNoSQLProvider() string {
	return properties.defaultNoSQLProvider
}

func (properties ApplicationProperties) GetDefaultTimeZone() string {
	return properties.defaultTimeZone
}
