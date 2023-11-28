package piscine

func IsNumeric(s string) bool {
	aString := []rune(s)
	sizeString := len(s)
	var isNumeric bool
	var counter int
	for i := range aString {
		if aString[i] >= 48 && aString[i] <= 57 {
			counter++
		}
		if counter == sizeString {
			isNumeric = true
		} else {
			isNumeric = false
		}
	}
	return isNumeric
}
