package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 6, 4, 3, 1}
	fmt.Println(buy(nums))
}

//优解  思路是分情况 具体分析
func buyBest(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	var (
		buy   int
		isBuy = false
		price int
	)
	len := len(prices)

	for i := 0; i < len; i++ {
		if i < len-1 && prices[i] < prices[i+1] {
			if !isBuy {
				isBuy = true
				buy = prices[i]
			}
		} else if i > 0 && prices[i] > prices[i-1] {
			if isBuy && prices[i] >= prices[i+1] {
				isBuy = false
				price += prices[i] - buy
			}
			if isBuy && i == len-1 {
				isBuy = false
				price += prices[i] - buy
				return price
			}
		}
	}
	return price
}

//非最优解
func buy(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	i := 0
	var price int
	for j := 1; j < len(prices); j++ {
		if prices[j] > prices[i] {
			price += prices[j] - prices[i]
		}
		i++
	}
	return price
}
