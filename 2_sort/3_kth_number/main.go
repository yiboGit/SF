package main

import "fmt"

func main() {
	nums := []int{4, 3, 5, 2, 1, 6, 9, 8}
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("%v \n", nums)

}

func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quickSort(nums []int, start, end int) {
	if start > end {
		return
	}
	midVal := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < midVal {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[end], nums[i] = nums[i], nums[end]
	quickSort(nums, start, i-1)
	quickSort(nums, i+1, end)

}
