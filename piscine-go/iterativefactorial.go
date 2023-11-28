package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int = 1
	for i := 2; i <= nb; i++ {
		if result*i <= result {
			return 0
		}
		result *= i
	}
	return result
}
