package main

func main() {

}

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 示例:
// 现有矩阵 matrix 如下：
// [
//   [1,   4,  7, 11, 15],   <-从15这个位置来进行循环判断 比位置大 则 行 索引加一，小 则 列 索引减一
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]

//由上图可以看到 由左到右增大;由上到下 增大； 整齐趋势是 左上角到 右下角增大的 趋势

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix) // 行
	if n == 0 {
		return false
	}
	m := len(matrix[0]) // 列

	i, j := 0, m-1 //i 行索引，j 列索引
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}
	}
	return false
}
