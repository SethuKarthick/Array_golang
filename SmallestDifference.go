package main

import (
	"fmt"
	"math"
	"sort"
)

var ary1 = []int{-1, 5, 10, 20, 28, 3}
var ary2 = []int{26, 134, 135, 15, 17}

func SmallestDifference(ary1, ary2 []int) []int {
	sort.Ints(ary1)
	sort.Ints(ary2)
	idxOne, idxTwo := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for idxOne < len(ary1) && idxTwo < len(ary2) {
		first, second := ary1[idxOne], ary2[idxTwo]
		if first < second {
			current = second - first
			idxOne += 1
		} else if second < first {
			current = first - second
			idxTwo += 1
		} else {
			return []int{first, second}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}

	}
	return smallestPair
}

func main() {
	res := SmallestDifference(ary1, ary2)
	fmt.Println(res)
}
