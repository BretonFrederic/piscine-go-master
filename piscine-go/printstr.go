package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	aRune := []rune(s)
	for i := range aRune {
		z01.PrintRune(aRune[i])
	}
}
