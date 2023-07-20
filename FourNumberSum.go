package main

import "fmt"

var (
	array  = []int{7, 6, 4, -1, 1, 2}
	target = 16
)

func FourNumberSum(array []int, target int) [][]int {
	allPairs := map[int][][]int{}
	quadraplets := [][]int{}
	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, found := allPairs[difference]; found {
				for _, pair := range pairs {
					newquad := append(pair, array[i], array[j])
					quadraplets = append(quadraplets, newquad)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			allPairs[currentSum] = append(allPairs[currentSum], []int{array[k], array[i]})
		}

	}
	return quadraplets
}

func main() {
	res := FourNumberSum(array, target)
	fmt.Println(res)
}
