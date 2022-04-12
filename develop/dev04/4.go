package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramm(mass *[]string) *map[string]*[]string {
	helpMass := make([]string, len(*mass))
	copy(helpMass, *mass) // копируем основной массив

	for i, elem := range helpMass {
		helpMass[i] = strings.ToLower(elem)
	}
	var runeMass [][]rune

	for i := 0; i < len(helpMass); i++ {
		helpRune := []rune(helpMass[i])
		sort.Slice(helpRune, func(k, j int) bool {
			return helpRune[k] < helpRune[j]
		})
		runeMass = append(runeMass, helpRune)
	}

	mapI := make(map[string][]int)
	for i := 0; i < len(helpMass); i++ {
		mapI[string(runeMass[i])] = append(mapI[string(runeMass[i])], i)
	}

	resMap := make(map[string]*[]string)
	for key := range mapI {
		helpSet := make(map[string]bool)
		for _, item := range mapI[key] {
			helpSet[helpMass[item]] = true
		}
		if len(helpSet) == 1 {
			continue
		}
		mass1 := new([]string)
		for item := range helpSet {
			*mass1 = append(*mass1, item)
		}
		sort.Strings(*mass1)
		resIndex := (*mass1)[0]
		resMap[resIndex] = mass1
	}

	return &resMap
}

func main() {
	mass := []string{"Леса", "сЕла", "села", "олег", "лего", "иван", "Иван"}
	mapa := findAnagramm(&mass)
	for key, item := range *mapa {
		fmt.Print(key)
		fmt.Println(*item)
	}
}
