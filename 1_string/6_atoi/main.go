package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(str string) int {
	// a := []rune(str)
	index := 0
	result := 0
	sign := 1

	// 去除空的字符
	for index < len(str) && str[index] == ' ' {
		index++
	}

	// 判断空字符后是否跟着符号
	if index < len(str) {
		if str[index] == '-' {
			index++
			sign = -1
		} else if str[index] == '+' {
			index++
		}
	}

	//判断 (有无符号后) 数字
	for index < len(str) {
		if str[index] >= '0' && str[index] <= '9' {
			result = result*10 + int(str[index]-'0')
			// 要每次对结果进行最大 最小值的判断 防止超出 导致计算错误
			if sign*result > math.MaxInt32 {
				return math.MaxInt32
			} else if sign*result < math.MinInt32 {
				return math.MinInt32
			}
			index++
		} else {
			break
		}
	}

	return sign * result
}
