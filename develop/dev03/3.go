package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type columnSlice [][]string

func (col columnSlice) Len() int {
	return len(col)
}

func (col columnSlice) Less(i, j int) bool {
	return col[i][0] < col[j][0]
}

func (col columnSlice) Swap(i, j int) {
	col[i], col[j] = col[j], col[i]
}

func main() {
	var column = flag.Int("k", -1, "указание колонки для сортировки")
	var number = flag.Bool("n", false, "сортировать по числовому значению")
	var reverse = flag.Bool("r", false, "сортировать в обратном порядке")
	var unique = flag.Bool("u", false, "не выводить повторяющиеся строки")

	log.SetFlags(0)
	flag.Parse()

	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	resLines := make([]string, len(lines))

	if *column != -1 {
		columnMass := columnSlice{}
		for index, line := range lines {
			helpMass := []string{strings.Split(line, " ")[*column], strconv.Itoa(index)}
			columnMass = append(columnMass, helpMass)
		}

		if *number {
			var numString = ""
			for index, elem := range columnMass {
				for _, letter := range elem[0] {
					if letter >= '0' && letter <= '9' {
						numString += string(letter)
					}
				}
				columnMass[index][0] = numString
				numString = ""
			}
		}

		if *reverse {
			sort.Sort(sort.Reverse(columnMass))
		} else {
			sort.Sort(columnMass)
		}

		for index, elem := range columnMass {
			i, _ := strconv.Atoi(elem[1])
			resLines[index] = lines[i]
		}
	} else {
		copy(resLines, lines)
		if *reverse {
			sort.Sort(sort.Reverse(sort.StringSlice(resLines)))
		} else {
			sort.Strings(resLines)
		}
	}

	if *unique {
		for i := 0; i < len(resLines)-1; i++ {
			if resLines[i] == resLines[i+1] {
				resLines = del(resLines, i+1)
				i--
			}
		}
	}

	fmt.Println(strings.Join(resLines, "\n"))
}

func del(mass []string, i int) []string {
	if i != len(mass)-1 {
		copy(mass[i:], mass[i+1:])
	}
	mass = mass[:len(mass)-1]
	return mass
}
