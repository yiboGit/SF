package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}

//可以用三路快排 但是也有一个更简单的思路
// w  r  b 三个索引 永远代表 0 1 2 三种数的最后一个数的索引
// 初始化 为 -1  为了方便 索引的自增
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	w, r, b := -1, -1, -1

	for _, v := range nums {
		if v == 0 {
			b++
			r++
			w++
			nums[b] = 2
			nums[r] = 1
			nums[w] = 0
			fmt.Printf("%v", nums)
		} else if v == 1 {
			b++
			r++
			nums[b] = 2
			nums[r] = 1
			fmt.Printf("%v", nums)
		} else if v == 2 {
			b++
			fmt.Printf("%v", nums)
			nums[b] = 2
		}
	}
	fmt.Printf("%v", nums)
}
