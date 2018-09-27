package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{4, 1, 2, 1, 2, 4, 7}
	fmt.Printf("%v", onlyOne(i))
}

func onlyOneTwo(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	var res int
	for i, v := range nums {
		k := i % 2
		if k == 0 {
			res += v
		} else {
			res -= v
		}
	}
	return res

}

func onlyOne(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Sort(sort.IntSlice(nums))
	i := 0
	one := nums[i]
	for j := 1; j < len(nums); j += 2 {
		if nums[j] != one {
			return one
		}
		if j+2 < len(nums) {
			i += 2
			one = nums[i]
		} else if j+1 < len(nums) {
			return nums[j+1]
		}
	}
	return 0
}
