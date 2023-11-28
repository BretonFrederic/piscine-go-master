package piscine

func IsLower(s string) bool {
	aString := []rune(s)
	sizeString := len(s)
	var isLower bool
	var counter int
	for i := range aString {
		if aString[i] >= 97 && aString[i] <= 122 {
			counter++
		}
	}
	if counter == sizeString {
		isLower = true
	} else {
		isLower = false
	}
	return isLower
}
