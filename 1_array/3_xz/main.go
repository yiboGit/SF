package main

import (
	"fmt"
	"log"
)

func main() {
	// ints := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Printf("%v", xz(ints, 3))
	s := 3 % 3
	fmt.Println(s)
}

// 旋转数组
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题
// 要求使用空间复杂度为 O(1) 的原地算法
func xzTwo(nums []int, k int) {
	k = k % len(nums)
	if k != 0 {
		copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
	}
}

func xz(nums []int, k int) []int {
	if len(nums) <= k {
		return nums
	}
	target := nums[len(nums)-k:]
	log.Printf("%v ok", len(nums))
	// target[i] = 1
	for _, item := range nums[:len(nums)-k] {
		target = append(target, item)
	}
	return target
}
