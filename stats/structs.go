package stats

type Stat struct {
	SourceIPsUnsorted   map[string]int
	SourceIPs           []genericKeyValuePair
	MethodsUnsorted     map[string]int
	Methods             []genericKeyValuePair
	URIUnsorted         map[string]int
	URI                 []genericKeyValuePair
	ProtocolsUnsorted   map[string]int
	Protocols           []genericKeyValuePair
	StatusCodesUnsorted map[string]int
	StatusCodes         []genericKeyValuePair
	Date                map[int64]int
	ByHour              map[string]map[int64]int
}

type genericKeyValuePair struct {
	Key   string
	Value int
}
