package main

import (
	"sort"
)

func main() {
	// []int{-1, 0, 1, 2, -1, -4}
	// -4,-1,-1,0,1,2
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 2 {
		return nil //[][]int{}
	}
	var result [][]int
	// 先将3数求和数组排序,按照大小
	sort.Sort(sort.IntSlice(nums))
	//第一层循环，找这个数字的指定和的另外两个数 a+b+c, a索引
	for index, _ := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		// 预定两个数的位置,eg: b,c索引
		// 一个指针从头走（a的下一个预定为eg:b的索引），一个指针从尾部倒(预定为eg:c的索引)
		i := index + 1
		j := length - 1
		for i < j {
			sum := nums[index] + nums[i] + nums[j]
			if sum == 0 {
				res := []int{nums[index], nums[i], nums[j]}
				result = append(result, res)
				i++
				j--
				// 可能有多种组合存在，但是相同的组合 只能出现一次
				//查看a这个数的其他组合可能
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
				//sum > 0,因为数组从小到大排序，所以说明大数大了，减小大数索引
			} else if sum > 0 {
				j--
				//sum < 0,因为数组从小到大排序，所以说明小数小了，增大小数索引
			} else if sum < 0 {
				i++
			}
		}
	}
	return result
}
