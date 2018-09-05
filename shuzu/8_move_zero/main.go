package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2}
	moveZero(nums)
	fmt.Printf("%v", nums)
}

// 没有0 大家一起走 ， 有0 i就停下来，index继续走
func moveZero(nums []int) {
	i := 0
	for index, value := range nums {
		if value != 0 {
			nums[index], nums[i] = nums[i], nums[index]
			i++
		}
	}
}
