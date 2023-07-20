package main

import (
	"fmt"
	"sort"
)

var _array = []int{3, 5, -4, 8, 11, 1, -1, 6}

func TwoNumberSum(_array []int, target int) []int {
	sort.Ints(_array)
	left, right := 0, len(_array)-1
	for left < right {
		currentSum := _array[left] + _array[right]
		if currentSum == target {
			return []int{_array[left], _array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}

	}
	return []int{}
}

func main() {
	res := TwoNumberSum(_array, 10)
	fmt.Println(res)
}
