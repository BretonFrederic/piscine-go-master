package piscine

//import (
//	"strings"
//)

func ToUpper(s string) string {
	// aString := strings.ToUpper(s)
	aString := []rune(s)
	for i := range aString {
		if aString[i] >= 97 && aString[i] <= 122 {
			aString[i] = aString[i] - 32
		}
	}
	newString := string(aString)
	return newString
}
