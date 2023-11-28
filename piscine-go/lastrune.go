package piscine

func LastRune(s string) rune {
	lastRune := []rune(s)
	var runeFinal rune
	for i := range lastRune {
		runeFinal = lastRune[i]
	}
	return runeFinal
}

// return rune(s[0])
