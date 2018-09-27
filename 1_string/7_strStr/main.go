package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("hell", "ll"))
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}
	if haystackLen < needleLen {
		return -1
	}

	for j := 0; j <= haystackLen-needleLen; j++ {
		if haystack[j:j+needleLen] == needle {
			return j
		}
	}
	return -1
}
