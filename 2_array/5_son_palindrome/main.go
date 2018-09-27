package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("abcba"))
}

// 定义两个全局参数  lo表示开始的索引，maxLen表示回文的最大长度
var (
	lo     int
	maxLen int
)

// 思路十分巧妙
func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	for i, _ := range s {
		//此处调用两次 为两种回文情况
		// 1.  12321
		extendPalindrome(s, i, i)
		// 2.  abccba
		extendPalindrome(s, i, i+1)
		// if i+2 >= len {
		// 	break
		// }
	}
	return s[lo : lo+maxLen]
}

func extendPalindrome(s string, j, k int) {
	for k < len(s) && j >= 0 && s[j] == s[k] {
		//满足条件 做一次指针的移动
		//以 j,k 为中心向左右，外部移动
		j--
		k++
	}
	// 012321   这里的 -1是为了消除，不满足条件的时候 j--;k++ 指针移动的影响
	if maxLen < (k - j - 1) {
		lo = j + 1 //这里 j+1 是消除 之前循环 j-- 的指针移动影响
		maxLen = k - j - 1
	}
}
