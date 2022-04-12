package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var fields = flag.Int("f", -1, "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	localDelimiter := *delimiter
	if *separated {
		var helpMass []string
		for _, line := range lines {
			if strings.Contains(line, localDelimiter) {
				helpMass = append(helpMass, line)
			}
		}
		lines = helpMass
	}

	if *fields > -1 {
		var res []string
		var helpMass [][]string
		for _, line := range lines {
			helpMass = append(helpMass, strings.Split(line, localDelimiter))
		}
		for _, line := range helpMass {
			if *fields >= len(line) {
				res = append(res, "")
			} else {
				res = append(res, line[*fields])
			}
		}
		lines = res
	}

	fmt.Printf(strings.Join(lines, "\n"))
}
