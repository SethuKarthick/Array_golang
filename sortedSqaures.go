package main

import "fmt"

var _array = []int{1, 2, 3, 5, 6, 8, 9}

func SortedSquares(_array []int) []int {
	SortedSquares_ := make([]int, len(_array))
	smallIdx := 0
	largerIdx := len(_array) - 1

	for idx := len(_array) - 1; idx >= 0; idx-- {
		smarllerValue := _array[smallIdx]
		largeerValue := _array[largerIdx]
		if abs(smarllerValue) > abs(largeerValue) {
			SortedSquares_[idx] = smarllerValue * smarllerValue
			smallIdx += 1
		} else {
			SortedSquares_[idx] = largeerValue * largeerValue
			largerIdx -= 1
		}
	}
	return SortedSquares_
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	res := SortedSquares(_array)
	fmt.Print(res)
}
