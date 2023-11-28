package piscine

func AppendRange(min, max int) []int {
	var arrayListNumbers []int
	if min >= max {
		return arrayListNumbers
	} else {
		for i := min; i < max; i++ {
			if i >= min && i < max {
				arrayListNumbers = append(arrayListNumbers, i)
			}
		}
	}
	return arrayListNumbers
}
