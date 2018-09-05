package main

import (
	"fmt"
)

func main() {
	nums := []int{9}
	fmt.Printf("%v \n", addOne(nums))
}

func addOne(num []int) []int {
	len := len(num)
	// 分清楚情况 非首位为9 与 首位为9, 例如 9  /  8999
	for j := len - 1; j >= 0; j-- {
		if num[j] < 9 {
			num[j]++
			return num
		} else if j == 0 {
			result := make([]int, len+1)
			result[0] = 1
			return result
			// result = append(result, 1)
			// result = append(result, num...)
		} else {
			num[j] = 0
		}
	}
	return []int{}
}
