package main

import "fmt"

var array = []int{5, 1, 4, 2}

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))
	for i := range array {
		products[i] = 1
	}
	leftRunningIdx := 1
	for i := range array {
		products[i] = leftRunningIdx
		leftRunningIdx *= array[i]
	}
	rightRunningIdx := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightRunningIdx
		rightRunningIdx *= array[i]
	}
	return products
}

func main() {
	res := ArrayOfProducts(array)
	fmt.Println(res)
}
