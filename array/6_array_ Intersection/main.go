package main

import (
	"fmt"
	"sort"
)

func main() {
	one := []int{-2147483648, 1, 2, 3}
	two := []int{1, -2147483648, -2147483648}
	fmt.Printf("%v \n", intersect(one, two))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))
	i := 0
	j := 0
	var smaller, largar, result []int
	if len(nums1) > len(nums2) {
		smaller = nums2
		largar = nums1
	} else {
		smaller = nums1
		largar = nums2
	}
	for j < len(largar) {
		if smaller[i] > largar[j] {
			j++
			if j >= len(largar) {
				break
			}
			continue
		}
		if smaller[i] < largar[j] {
			i++
			if i >= len(smaller) {
				break
			}
			continue
		}
		if smaller[i] == largar[j] {
			result = append(result, largar[j])
			i++
			j++
			if i >= len(smaller) || j >= len(largar) {
				break
			}
		}
	}
	return result
}
