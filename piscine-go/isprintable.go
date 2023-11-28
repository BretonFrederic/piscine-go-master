package piscine

func IsPrintable(s string) bool {
	aString := []rune(s)
	sizeString := len(s)
	var isPrintable bool
	var counter int
	for i := range aString {
		if aString[i] >= 32 && aString[i] <= 126 {
			counter++
		}
		if counter == sizeString {
			isPrintable = true
		} else {
			isPrintable = false
		}
	}
	return isPrintable
}
