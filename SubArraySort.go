package main

import (
	"fmt"
	"math"
)

var array = []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}

func SubArraytSort(array []int) []int {
	minOutOforder, maxOutOfOrder := math.MaxInt32, math.MinInt32
	for i, num := range array {
		if isOutOfOrder(i, num, array) {
			minOutOforder = min(minOutOforder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}
	if minOutOforder == math.MaxInt32 {
		return []int{-1, -1}
	}
	subArrayLeft := 0
	for minOutOforder >= array[subArrayLeft] {
		subArrayLeft += 1
	}
	subArrayRight := len(array) - 1
	for maxOutOfOrder <= array[subArrayRight] {
		subArrayRight -= 1
	}
	return []int{subArrayLeft, subArrayRight}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isOutOfOrder(i int, num int, array []int) bool {
	if i == 0 {
		return num > array[i+1]
	}
	if i == len(array)-1 {
		return num < array[i-1]
	}
	return num > array[i+1] || num < array[i-1]
}

func main() {
	res := SubArraytSort(array)
	fmt.Println(res)
}
