package main

import (
	"fmt"
	"strings"
)

func main() {
	v1 := "1.2.3a"
	v2 := "1.2.3b"
	v1s := strings.Split(v1, ".")
	v2s := strings.Split(v2, ".")
	fmt.Println(compareVers(v1s, v2s))
}

func compareVers(v1, v2 []string) int {
	if len(v1) != len(v2) {
		return -1
	}

	for index, str := range v1 {
		rune1, rune2 := []rune(str), []rune(v2[index])
		oneResult := compareRunes(rune1, rune2)
		if oneResult != 0 {
			return oneResult
		}
	}
	return 0
}

func compareRunes(rune1, rune2 []rune) int {
	if len(rune1) > len(rune2) {
		return 1
	}
	if len(rune1) < len(rune2) {
		return 2
	}

	for index, _ := range rune1 {
		if rune1[index] > rune2[index] {
			return 1
		}
		if rune1[index] < rune2[index] {
			return 2
		}

	}
	return 0
}
