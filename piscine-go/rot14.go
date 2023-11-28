package piscine

func Rot14(s string) string {
	arrayRune := []rune(s)

	for l := range arrayRune {
		if arrayRune[l] >= 'A' && arrayRune[l] <= 'Z' {
			arrayRune[l] += 14
			if arrayRune[l] > 'Z' {
				arrayRune[l] = ('A' - 1) + (arrayRune[l] - 'Z')
			}
		}
		if arrayRune[l] >= 'a' && arrayRune[l] <= 'z' {
			arrayRune[l] += 14
			if arrayRune[l] > 'z' {
				arrayRune[l] = ('a' - 1) + (arrayRune[l] - 'z')
			}
		}
	}
	newString := string(arrayRune)
	return newString
}
