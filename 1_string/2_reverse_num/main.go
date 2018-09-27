package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%v", reverse(-192))
}

// 我写的这个竟然是速度最快的
func reverse(x int) int {
	var result int
	symbol := 1

	if x < 0 {
		symbol = -1
		x = x * -1
	}
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > math.MaxInt32 {
		return 0
	}
	return result * symbol
}
