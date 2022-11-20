package logs

import (
	"bufio"
	"log"
	"os"
)

func ReadLog(Log string, callback func(string)) {
	file, err := os.Open(Log)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
