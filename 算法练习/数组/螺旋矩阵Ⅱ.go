// author:zfy  date:2022/7/7

package main

import "fmt"

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1
	target := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= target {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}

func main() {
	n := 4
	fmt.Println(generateMatrix(n))
}
