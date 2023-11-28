package piscine

func IsUpper(s string) bool {
	aString := []rune(s)
	sizeString := len(s)
	var isUpper bool
	var counter int
	for i := range aString {
		if aString[i] >= 65 && aString[i] <= 90 {
			counter++
		}
	}
	if counter == sizeString {
		isUpper = true
	} else {
		isUpper = false
	}
	return isUpper
}
