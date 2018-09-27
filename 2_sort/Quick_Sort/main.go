package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 5, 2, 1, 6, 9, 8}
	QuickSort(nums, 0, len(nums)-1)

	fmt.Printf("%v \n", nums)

}

// 快速排序 递归的方法
func QuickSort(nums []int, start, end int) {
	if start > end {
		return
	}

	i := start
	midVal := nums[end]

	for j := start; j < end; j++ {
		if nums[j] < midVal {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]
	QuickSort(nums, start, i-1)
	QuickSort(nums, i+1, end)

}
