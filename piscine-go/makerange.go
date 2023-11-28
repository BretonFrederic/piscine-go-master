package piscine

func MakeRange(min, max int) []int {
	var arrayListNumbers []int
	var counter int
	if min >= max {
		return arrayListNumbers
	} else {
		for i := min; i <= max; i++ {
			counter++
		}
		arrayListNumbers = make([]int, counter-1)
		for u := range arrayListNumbers {
			for t := min; t < max; t++ {
				arrayListNumbers[u] = min + u
			}
		}
	}
	return arrayListNumbers
}
