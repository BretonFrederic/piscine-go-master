package piscine

func ToLower(s string) string {
	// aString := strings.toLower(s)
	aString := []rune(s)
	for i := range aString {
		if aString[i] >= 65 && aString[i] <= 90 {
			aString[i] = aString[i] + 32
		}
	}
	newString := string(aString)
	return newString
}
