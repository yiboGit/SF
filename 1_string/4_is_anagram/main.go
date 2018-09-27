package main

import (
	"fmt"
)

func main() {
	s := "anagramjdh"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

// 借用上一个思想 所以这是最优解
func isAnagram(s string, t string) bool {
	sa := make([]int, 26)
	ta := make([]int, 26)

	for _, c := range s {
		sa[c-'a']++
	}

	for _, c := range t {
		ta[c-'a']++
	}

	for i, v := range sa {
		fmt.Printf("sa %v ; ta %v\n", v, ta[i])
		if ta[i] != v {
			return false
		}
	}
	return true

}
