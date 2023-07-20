package main

import (
	"fmt"
)

var matrix = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

func TransposeMatrix(matrix [][]int) [][]int {
	// Write your code here.
	transposedMatrix := make([][]int, len(matrix[0]))
	for col := 0; col < len(matrix[0]); col++ {
		newRow := make([]int, len(matrix))
		fmt.Println(newRow)
		for row := 0; row < len(matrix); row++ {
			newRow[row] = matrix[row][col]
		}
		transposedMatrix[col] = newRow
		fmt.Println(transposedMatrix)
	}
	return transposedMatrix
}

func main() {
	res := TransposeMatrix(matrix)
	fmt.Println(res)
}
