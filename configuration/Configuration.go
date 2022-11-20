package configuration

type config struct {
	LogFile         string
	SourceIPRegex   string
	MethodRegex     string
	URIRegex        string
	ProtocolRegex   string
	StatusCodeRegex string
	DateRegex       string
	TimeRegex       string
	DateFormat      string
}

var Config config
