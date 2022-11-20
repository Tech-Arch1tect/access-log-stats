package stats

import (
	"log"
	"time"

	"github.com/Tech-Arch1tect/access-log-stats/configuration"
)

func getStatsProcess(line string) {
	statsToGet := make(map[string]string)
	statsToGet["SourceIP"] = configuration.Config.SourceIPRegex
	statsToGet["Method"] = configuration.Config.MethodRegex
	statsToGet["Protocol"] = configuration.Config.ProtocolRegex
	statsToGet["StatusCode"] = configuration.Config.StatusCodeRegex
	statsToGet["URI"] = configuration.Config.URIRegex
	statsToGet["Date"] = configuration.Config.DateRegex
	statsToGet["Time"] = configuration.Config.TimeRegex

	for k, v := range statsToGet {
		if v != "" {
			val := getValFromLine(line, v)
			if val != "" {
				switch k {
				case "SourceIP":
					Stats.SourceIPsUnsorted[val]++
				case "Method":
					Stats.MethodsUnsorted[val]++
				case "Protocol":
					Stats.ProtocolsUnsorted[val]++
				case "StatusCode":
					Stats.StatusCodesUnsorted[val]++
				case "URI":
					Stats.URIUnsorted[val]++
				case "Date":
					dateObj, err := time.Parse(configuration.Config.DateFormat, val)
					if err != nil {
						log.Println(err)
					} else {
						Stats.Date[dateObj.Unix()]++
					}
				case "Time":
					if configuration.Config.DateRegex != "" {
						// get date
						date := getValFromLine(line, configuration.Config.DateRegex)
						dateTimeObj, err := time.Parse(configuration.Config.DateFormat+" 15:04", date+" "+val[0:2]+":00")
						if err != nil {
							log.Println(err)
						} else {
							// initialise map if not already
							if Stats.ByHour[date] == nil {
								Stats.ByHour[date] = make(map[int64]int)
							}
							// handle date by hour
							Stats.ByHour[date][dateTimeObj.Unix()]++
						}
					}
				}
			}
		}
	}
}
