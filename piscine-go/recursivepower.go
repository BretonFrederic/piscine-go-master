package piscine

func RecursivePower(nb int, power int) int {
	result := nb
	if nb > 5000 {
		return 0
	}
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power > 1 {
		// result = result * nb
		return result * RecursivePower(nb, (power-1))
	}
	return result
}
