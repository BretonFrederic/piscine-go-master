package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	tableau := [10]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	if n == 0 {
		z01.PrintRune('0')
	} else {
		res := n
		if n < 0 {
			z01.PrintRune('-')
			res = n * -1
		}
		number1 := res % 10
		number2 := (res % 100) / 10
		number3 := (res % 1000) / 100

		z01.PrintRune(rune(tableau[number3]))
		z01.PrintRune(rune(tableau[number2]))
		z01.PrintRune(rune(tableau[number1]))
	}
}
