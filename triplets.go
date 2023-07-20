package main

import (
	"fmt"
	"sort"
)

var _array = []int{12, 3, 1, 2, -6, 5, -8, 6}
var target = 0

func Triplets(_array []int, target int) [][]int {
	sort.Ints(_array)
	triplets := [][]int{}
	for i := 0; i < len(_array)-2; i++ {
		left, right := i+1, len(_array)-1
		for left < right {
			currentSum := _array[i] + _array[left] + _array[right]
			if currentSum == target {
				triplets = append(triplets, []int{_array[i], _array[left], _array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else if currentSum > target {
				right -= 1
			}
		}
	}
	return triplets
}

func main() {
	res := Triplets(_array, target)
	fmt.Println(res)
}
