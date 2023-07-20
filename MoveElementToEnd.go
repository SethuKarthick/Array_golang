package main

import "fmt"

var (
	array      = []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove int = 2
)

func MoveElementToEnd(array []int, toMove int) []int {
	i, j := 0, len(array)-1
	for i < j {
		for i < j && array[j] == toMove {
			j--
		}
		if array[i] == toMove {
			array[i], array[j] = array[j], array[i]
		}
		i++
	}
	return array
}

func main() {
	res := MoveElementToEnd(array, toMove)
	fmt.Println(res)
}
