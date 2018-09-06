package main

import (
	"fmt"
)

// 排序数据去重复元素
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2}
	fmt.Println(remove(nums))
}

func remove(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	// len
	fmt.Printf("%v", nums)
	return i + 1
}
