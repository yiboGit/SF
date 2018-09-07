package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome2(str))
}

// 循环 少一次
func isPalindrome2(s string) bool {
	array := []rune(s)
	i := 0
	j := len(array) - 1
	for i < j {
		if !unicode.IsLetter(array[i]) && !unicode.IsNumber(array[i]) {
			i++
			continue
		}
		if !unicode.IsLetter(array[j]) && !unicode.IsNumber(array[j]) {
			j--
			continue
		}
		if unicode.ToLower(array[i]) != unicode.ToLower(array[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string) bool {
	var array []rune
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			array = append(array, unicode.ToLower(c))
		}
	}
	len := len(array)
	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		if array[i] != array[j] {
			fmt.Println(array)
			return false
		}
	}
	return true
}
