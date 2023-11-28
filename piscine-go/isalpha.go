package piscine

func IsAlpha(s string) bool {
	aString := []rune(s)
	sizeString := len(s)
	var counter int
	var isAlpha bool
	for i := range aString {
		if aString[i] >= 65 && aString[i] <= 90 || aString[i] >= 97 && aString[i] <= 122 || aString[i] >= 48 && aString[i] <= 57 {
			counter++
		}
		if counter == sizeString {
			isAlpha = true
		} else {
			isAlpha = false
		}
	}
	return isAlpha
}
