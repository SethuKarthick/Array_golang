package main

import "fmt"

var array = []int{2, 1, 5, 2, 3, 3, 4}

// func FirstDuplicateValue(array []int) int {
// 	for _, value := range array {
// 		absValue := abs(value)
// 		if array[absValue-1] < 0 {
// 			return absValue
// 		}
// 		array[absValue-1] *= -1
// 	}
// 	return -1
// }

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func FirstDuplicateValue(array []int) int {
	seen := map[int]bool{}
	for _, value := range array {
		if seen[value] {
			return value
		}
		seen[value] = true
	}
	return -1
}

func main() {
	res := FirstDuplicateValue(array)
	fmt.Println(res)
}
