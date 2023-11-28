package piscine

func IsSorted(comparison func(first, second int) int, array []int) bool {
	var sorted bool
	counter1 := 1
	counter2 := 1

	for i := 1; i < len(array); i++ {
		if comparison(array[i-1], array[i]) > 0 || comparison(array[i-1], array[i]) == 0 {
			counter1++
		}
	}
	for i := len(array) - 1; i > 0; i-- {
		if comparison(array[i], array[i-1]) > 0 {
			counter2++
		}
	}
	if counter1 == len(array) || counter2 == len(array) {
		sorted = true
	} else {
		sorted = false
	}
	return sorted
}
