package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	printNXN(a)
}

func printNXN(a [][]int) {
	rowLen := len(a)
	colLen := len(a[0])
	if rowLen == 0 && colLen == 0 {
		return
	}
	// 分两部分 记性打印  上班不 从右上角 进行打印
	for colIndex := colLen - 1; colIndex >= 0; colIndex-- {
		i, j := 0, colIndex
		for i < rowLen && j < colLen && i <= j {
			fmt.Print(a[i][j])
			i++
			j++
		}
	}

	// 下半部 从中对角线 以下进行打印
	for rowIndex := 0; rowIndex < rowLen; rowIndex++ {
		i, j := rowIndex, 0
		for i < rowLen && j < colLen && i > j {
			fmt.Print(a[i][j])
			i++
			j++
		}
	}
}
