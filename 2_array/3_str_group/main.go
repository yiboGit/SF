package main

import "sort"

func main() {

}

// 优化之后的方法 可以不是通过实现sort接口进行排序
func groupAnagramsYH(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		// 通过sort 切片 指定数据结构，与比较规则
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		groupMap[string(runes)] = append(groupMap[string(runes)], str)
	}

	var result [][]string
	for _, value := range groupMap {
		result = append(result, value)
	}
	return result
}

// 通过实现sort接口，进行排序
func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, str := range strs {
		sign := Bytes(str)
		sort.Sort(sign)
		groupMap[string(sign)] = append(groupMap[string(sign)], str)
	}

	var result [][]string
	for _, value := range groupMap {
		result = append(result, value)
	}
	return result
}

type Bytes []rune

func (b Bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b Bytes) Len() int {
	return len(b)
}
