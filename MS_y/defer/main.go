package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("前")
	defer fmt.Println("中")
	defer fmt.Println("后")

	panic("错误")
}
