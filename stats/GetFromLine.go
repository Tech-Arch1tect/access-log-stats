package stats

import (
	"log"
	"regexp"
)

func getValFromLine(line string, regex string) string {
	r, err := regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}
	if r.MatchString(line) {
		val := r.FindAllStringSubmatch(line, -1)
		return val[0][1]
	} else {
		log.Println("Missing value in log line. Regex: " + regex + ".")
		return ""
	}
}
