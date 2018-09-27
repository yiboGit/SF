package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "loveleetcode"
	fmt.Printf("%v", firstUniqCharTwo(s))
}

func firstUniqCharTwo(s string) int {
	array := make([]int, 26)
	for _, c := range s {
		array[c-'a']++
	}

	for index, c := range s {
		if array[c-'a'] == 1 {
			return index
		}
	}
	return -1
}

// 感觉这次自己的解法 很不好  效率渣到炸
func firstUniqCharOne(s string) int {
	if s == "" {
		return -1
	}
	indexMap := make(map[string]int)
	existMap := make(map[string]bool)

	runes := []rune(s)
	for index, value := range runes {
		if _, ok := indexMap[string(value)]; ok {
			existMap[string(value)] = true
		} else {
			indexMap[string(value)] = index
			existMap[string(value)] = false
		}
	}

	var ints []int
	for index, value := range existMap {
		if !value {
			ints = append(ints, indexMap[index])
		}
	}
	if len(ints) == 0 {
		return 0
	}
	sort.Sort(sort.IntSlice(ints))
	return ints[0]
}
