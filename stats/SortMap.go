package stats

import "sort"

func sortMaps(stats Stat) Stat {
	stats.SourceIPs = sortMap(stats.SourceIPsUnsorted)
	stats.Methods = sortMap(stats.MethodsUnsorted)
	stats.Protocols = sortMap(stats.ProtocolsUnsorted)
	stats.StatusCodes = sortMap(stats.StatusCodesUnsorted)
	stats.URI = sortMap(stats.URIUnsorted)
	return stats
}

func sortMap(dataMap map[string]int) []genericKeyValuePair {
	var valuePair []genericKeyValuePair
	for k, v := range dataMap {
		var keyVal genericKeyValuePair
		keyVal.Key = k
		keyVal.Value = v
		valuePair = append(valuePair, keyVal)
	}
	sort.Slice(valuePair, func(i, j int) bool {
		return valuePair[i].Value > valuePair[j].Value
	})

	return valuePair
}
