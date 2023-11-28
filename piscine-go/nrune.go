package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return 0
	}
	NRune := []rune(s)
	var runeFinal rune
	for i := 0; i < n; i++ {
		runeFinal = NRune[i]
	}
	return runeFinal
}
