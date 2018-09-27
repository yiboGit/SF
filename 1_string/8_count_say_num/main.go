package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	result := []byte{'1'}
	var (
		target []byte
		count  int
	)
	for j := 2; j <= n; j++ {
		for i := 0; i < len(result); i++ {
			count++
			if i+1 >= len(result) || result[i+1] != result[i] {
				target = append(target, fmt.Sprintf("%d%c", count, result[i])...)
				count = 0
			}
		}
		result = target
		target = []byte{}
	}
	return string(result)
}
