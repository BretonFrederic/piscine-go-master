package piscine

func AlphaCount(s string) int {
	countAlphabet := []rune(s)
	var totalAlphabetLetter int
	var index int
	var counterCapital int
	var counterLowercase int
	// 65 90 // 97 122
	for index = range countAlphabet {
		if countAlphabet[index] >= 'A' && countAlphabet[index] <= 'Z' {
			counterCapital++
		} else if countAlphabet[index] >= 'a' && countAlphabet[index] <= 'z' {
			counterLowercase++
		}
		totalAlphabetLetter = counterCapital + counterLowercase
	}
	return totalAlphabetLetter
}
