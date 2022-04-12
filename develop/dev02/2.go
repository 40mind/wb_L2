package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "///230/a10"
	fmt.Println(unpack(str))
}

func unpack(str string) string {
	outMass := make([]rune, 0)
	strMass := []rune(str)
	strLen := len(strMass)
	if strLen == 0 {
		return ""
	}
	if '0' <= strMass[0] && strMass[0] <= '9' {
		return ""
	}
	var lastRune rune

	for i := 0; i < strLen; i++ {
		if ('a' <= strMass[i] && strMass[i] <= 'z') || ('A' <= strMass[i] && strMass[i] <= 'Z') {
			if ('a' <= lastRune && lastRune <= 'z') || ('A' <= lastRune && lastRune <= 'Z') || ('0' <= lastRune && lastRune <= '9') {
				outMass = append(outMass, lastRune)
				lastRune = strMass[i]
			} else {
				lastRune = strMass[i]
			}
		} else if '0' <= strMass[i] && strMass[i] <= '9' {
			var num []rune
			for ; i < strLen; i++ {
				if '0' <= strMass[i] && strMass[i] <= '9' {
					num = append(num, strMass[i])
				} else {
					i--
					break
				}
			}
			if lastRune != ' ' {
				num, _ := strconv.Atoi(string(num))
				for j := 0; j < num; j++ {
					outMass = append(outMass, lastRune)
				}
				lastRune = ' '
			} else {
				return ""
			}
		} else if strMass[i] == '/' {
			if lastRune != ' ' && i != 0 {
				outMass = append(outMass, lastRune)
			}
			if i == strLen-1 {
				return ""
			}
			i++
			lastRune = strMass[i]
		}
	}
	if lastRune != ' ' {
		outMass = append(outMass, lastRune)
	}
	return string(outMass)
}
