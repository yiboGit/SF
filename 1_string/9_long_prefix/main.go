package main

import (
	"fmt"
)

func main() {
	strs := []string{"aacc,aacfgg,aacjkjk"}
	fmt.Println(longestCommonPrefix(strs))
}

// 思路两层循环 外层负责每个字符串的字符的索引，内层为字符串数组的索引
// 然后比较每次字符串 相同索引的字符，相同放到前缀里，不同则直接返回前缀 注意特殊情况 空字符串
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string
	for j := 0; ; j++ {
		var c string
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == 0 || j > len(strs[i])-1 {
				return prefix
			}
			if c == "" {
				c = string(strs[i][j])
			} else if string(strs[i][j]) != c {
				return prefix
			}
		}
		prefix = prefix + c
	}
}
