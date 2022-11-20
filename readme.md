This is an extremely basic go module which aims to collect some statistics from a web server access log.

This was written for learning and as such there are currently a number of problems with the module:

- Error handling either nonexistent or bad
- The regex's are hardcoded to match the first occurrence of the capture group only (stats/GetFromLine.go#L15)
- Lack of extensibility 
- Not built for performance
- etc etc


If you're feeling risky, here's a basic example:

```
// set log location
configuration.Config.LogFile = "/path/to/access.log"
// Set regex values from our config
configuration.Config.SourceIPRegex = "REGEX"
configuration.Config.MethodRegex = "REGEX"
configuration.Config.URIRegex = "REGEX"
configuration.Config.ProtocolRegex = "REGEX"
configuration.Config.StatusCodeRegex = "REGEX"
configuration.Config.DateRegex = "REGEX"
configuration.Config.TimeRegex = "REGEX"
configuration.Config.DateFormat = "2006-01-02"
// get stats
stats := stats.GetStats()
```