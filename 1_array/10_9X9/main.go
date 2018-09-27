package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", isValidSudoku([][]string{
		// {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}))
}

// 重点  三部分
// 再看
func isValidSudoku(board [][]string) bool {
	var rows, cols, boxes []map[string]bool
	for i := 0; i < 9; i++ {
		rows = append(rows, make(map[string]bool))
		cols = append(cols, make(map[string]bool))
		boxes = append(boxes, make(map[string]bool))
	}
	for i := range board {
		for j, num := range board[i] {
			if num == "." {
				continue
			}
			if rows[i][num] || cols[j][num] || boxes[(i/3)*3+j/3][num] {
				return false
			}
			rows[i][num] = true
			cols[j][num] = true
			boxes[(i/3)*3+j/3][num] = true
		}
	}
	return true
}
