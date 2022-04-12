package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")
	var context = flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	var count = flag.Bool("c", false, "количество строк")
	var ignore = flag.Bool("i", false, "игнорировать регистр")
	var invert = flag.Bool("v", false, "вместо совпадения, исключать")
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	var linenum = flag.Bool("n", false, "напечатать номер строки")

	log.SetFlags(0)
	flag.Parse()

	args := flag.Args()
	pattern := args[0]

	if *fixed {
		pattern = `\Q` + pattern + `\E`
	}

	if *ignore {
		pattern = "(?i)" + pattern
	}

	reg := regexp.MustCompile(pattern)

	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	a := max(*after, *context)
	b := max(*before, *context)

	var res []int
	for index, line := range lines {
		if reg.Match([]byte(line)) {
			res = append(res, index)
		}
	}

	if *linenum {
		for index := range lines {
			lines[index] = strconv.Itoa(index) + " " + lines[index]
		}
	}

	var resLines []string

	if *invert {
		for index, elem := range lines {
			if findElem(res, index) {
				continue
			}
			resLines = append(resLines, elem)
		}
		fmt.Println(strings.Join(resLines, "\n"))
		return
	}

	var sum = 0
	if *count {
		fmt.Println(len(res))
	}
	for _, elem := range res {
		sum = min(b, elem)
		for {
			if sum == 0 {
				break
			}
			resLines = append(resLines, lines[elem-sum])
			sum--
		}
		resLines = append(resLines, lines[elem])
		sum = 0
		for {
			if sum == min(a, len(lines)-elem-1) {
				break
			}
			resLines = append(resLines, lines[elem+sum+1])
			sum++
		}
	}
	fmt.Println(strings.Join(resLines, "\n"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findElem(str []int, elem int) bool {
	for _, index := range str {
		if index == elem {
			return true
		}
	}
	return false
}
