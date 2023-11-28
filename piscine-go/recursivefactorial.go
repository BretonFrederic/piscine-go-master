package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 5000 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		// nb = nb * (nb - 1)
		return nb * RecursiveFactorial(nb-1)
	}
	return result
	// result = 1 * 2 * 3 * 4
}
