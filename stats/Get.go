package stats

import (
	"github.com/Tech-Arch1tect/access-log-stats/configuration"
	"github.com/Tech-Arch1tect/access-log-stats/logs"
)

var Stats Stat

func GetStats() Stat {
	Stats = Stat{}
	Stats.SourceIPsUnsorted = make(map[string]int)
	Stats.MethodsUnsorted = make(map[string]int)
	Stats.URIUnsorted = make(map[string]int)
	Stats.ProtocolsUnsorted = make(map[string]int)
	Stats.StatusCodesUnsorted = make(map[string]int)
	Stats.Date = make(map[int64]int)
	Stats.ByHour = make(map[string]map[int64]int)
	logs.ReadLog(configuration.Config.LogFile, getStatsProcess)
	Stats = sortMaps(Stats)
	return Stats
}
