package main

import "fmt"

var array = []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}

func isMonotonic(array []int) bool {
	isNondecreasing := true
	isNonincreaseing := true

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNondecreasing = false
		}
		if array[i] > array[i-1] {
			isNonincreaseing = false
		}
	}
	return isNondecreasing || isNonincreaseing
}

func main() {
	res := isMonotonic(array)
	fmt.Println(res)
}
