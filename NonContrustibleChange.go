package main

import (
	"fmt"
	"sort"
)

var coins = []int{5, 7, 1, 1, 2, 3, 22}

func NonConstructibleChagne(coins []int) int {
	sort.Ints(coins)

	var currentChange = 0
	for _, coin := range coins {
		if coin > currentChange+1 {
			return currentChange + 1
		}
		currentChange += coin
	}
	return currentChange + 1
}

func main() {
	res := NonConstructibleChagne(coins)
	fmt.Println(res)
}
