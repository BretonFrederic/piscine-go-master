package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb > 5000 || power > 5000 {
		return 0
	}
	if nb <= 0 && power == 0 {
		return 1
	}
	result := nb
	for i := 1; i < power; i++ {
		result = result * nb
		//if result < 0 {
		//	return 0
		//}
	}

	return result
}
