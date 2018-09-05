package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 3, 7}
	fmt.Printf("%v", twoSum(nums, 9))
}

// 思路很厉害  先初始化一个map  假设 [2...7] 的数组  我输入9， 循环到2的时候 向map添加 key值为（9-2）那么对
// 应的是想要得到9结果我需要和7配对，此时我key: 7（通过9-2得到）对应的是 2的index； 当我走到7的时候，通过7可
// 以得到2的index  此时我得到了所有的结果
func twoSum(nums []int, target int) []int {
	//  A+B = C  A数一定存在  B不一定
	// map  value 是 A数的index  key: 理想中B的值（可以与A数得到正确结果的值）
	maps := make(map[int]int)
	for index, value := range nums {
		_, ok := maps[value]
		if ok {
			return []int{maps[value], index}
		} else {
			maps[target-value] = index
		}
	}
	return nil
}
