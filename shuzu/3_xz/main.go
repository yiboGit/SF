package main

import (
	"fmt"
	"log"
)

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v", xz(ints, 3))
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
