package main

import (
	"fmt"
	"sort"
)

func main() {

	ints := []int{1, 2, 3, 4, 5, 1}
	fmt.Printf("%v", onlyOne(ints))

}

func onlyOne(ints []int) bool {
	// 标准库中的可以使用
	sort.Sort(sort.IntSlice(ints))
	last := ints[0]
	for _, item := range ints[1:] {
		if last == item {
			return true
		}
		last = item
	}
	return false
}
