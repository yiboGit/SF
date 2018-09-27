package main

func main() {

}

// 输入:
// [
//   [1,1,1],
//   [1,0,1],
//   [1,1,1]
// ]
// 输出:
// [
//   [1,0,1],
//   [0,0,0],
//   [1,0,1]
// ]

func setZeroes(matrix [][]int) {
	// 第一列 是不是0的标记
	col0 := false
	// 行 索引
	rows := len(matrix)
	// 列 索引 （每一行的索引）
	cols := len(matrix[0])

	// 第一次大循环， 把上边界，左边界先变0
	// 内外正循环
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	// 第二次大循环， 根据 上边界，左边界是不是0
	// 来使其他部分变0  内外逆循环
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- { // 第一列有特殊操作
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		// 根据标记 判断是否把第一列变0
		if col0 == true {
			matrix[i][0] = 0
		}
	}
}
