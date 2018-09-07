package main

import (
	"fmt"
)

func main() {

	fmt.Println(reverseString("aaaaa aaaa"))

}

func reverseString(s string) string {
	runes := []rune(s)
	len := len(runes)
	fmt.Println(len)
	if len <= 1 {
		return s
	}
	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}
