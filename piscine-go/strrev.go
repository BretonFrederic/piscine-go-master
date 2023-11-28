package piscine

func StrRev(s string) string {
	aRune := []rune(s)
	var newRune []rune
	var reverseString string
	for i := len(aRune); i > 0; i-- {
		newRune = append(newRune, aRune[i-1])
	}
	reverseString = string(newRune)
	return reverseString
}
