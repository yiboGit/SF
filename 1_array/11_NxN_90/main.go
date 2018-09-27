package main

func main() {

}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n-n/2; j++ {
			//反转
			//eg 0,1,3,2
			a, b, c, d := i, j, n-i-1, n-j-1
			matrix[a][b], matrix[b][c], matrix[c][d], matrix[d][a] = matrix[d][a], matrix[a][b], matrix[b][c], matrix[c][d]
		}
	}
}
