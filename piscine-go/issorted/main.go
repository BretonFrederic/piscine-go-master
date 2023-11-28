package main

import (
	"fmt"
	"piscine"
)

func main() {
	array1 := []int{0, 1, 2, 3, 4, 5}
	array2 := []int{0, 2, 1, 3}
	array3 := []int{704648, 348155, 154936, 141932, 64803, -305519, -811400, -853633}

	result1 := piscine.IsSorted(comparison, array1)
	result2 := piscine.IsSorted(comparison, array2)
	result3 := piscine.IsSorted(comparison, array3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func comparison(first, second int) int {
	if first > second {
		return 1
	} else if first == second {
		return 0
	} else {
		return -1
	}
}
