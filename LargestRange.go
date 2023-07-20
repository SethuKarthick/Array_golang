package main

import "fmt"

var array = []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}

func LargestRange(array []int) []int {
	best := []int{}
	largestRange := 0
	nums := map[int]bool{}
	for _, num := range array {
		nums[num] = true
	}
	for _, num := range array {
		if !nums[num] {
			continue
		}
		nums[num] = false
		currentLenght, left, right := 1, num-1, num+1
		for nums[left] {
			nums[left] = false
			currentLenght += 1
			left -= 1
		}
		for nums[right] {
			nums[right] = false
			currentLenght += 1
			right += 1
		}
		if currentLenght > largestRange {
			largestRange = currentLenght
			best = []int{left + 1, right - 1}
		}
	}
	return best
}

func main() {
	res := LargestRange(array)
	fmt.Println(res)
}
