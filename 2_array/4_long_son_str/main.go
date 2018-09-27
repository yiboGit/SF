package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("aab"))
}

// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 无重复字符的最长子串是 "abc"，其长度为 3。
// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 无重复字符的最长子串是 "b"，其长度为 1。
// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 无重复字符的最长子串是 "wke"，其长度为 3。
//      请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串

// 思路 循环字符一遍，将字符更新到map中
// 如果遇到map中已存在本次循环字符的情况
// 比较存在于map中这个重复字符的索引
// 如果比sonStart，子串开始索引大， 那么将sonStart更新为map中重复字符索引
//-----其余部分就是必要做的 将字符及其索引更新到map中
//-----当前循环字符索引 - sonStart子串开始索引  求出长度，与maxLength进行比较，求出最大长度
func lengthOfLongestSubstring(s string) int {
	subMap := make(map[rune]int, 26)
	sonStart := -1 // -1 要考虑的事 索引为0的时候求长度
	maxLen := 0
	for i, v := range s {
		// 以为有 abcbhah 等情况
		if c, ok := subMap[v-'a']; ok {
			// 能取出 说明重复， 那么和目前子串开始索引进行比较
			// 比现在的子串索引大 成为新子串索引
			if c > sonStart {
				sonStart = c
			}
		}
		// 更新最新重复字符 的 索引
		subMap[v-'a'] = i
		// 查看当前子串长度 是不是 最大的
		if maxLen < i-sonStart {
			fmt.Println(i, sonStart)
			maxLen = i - sonStart
		}
	}
	return maxLen
}
