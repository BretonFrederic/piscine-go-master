package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	// Init variables
	number1 := ('0')
	number2 := ('0')
	number3 := ('0')
	addNumber2 := 0
	addNumber1 := 0
	// Iteration des centaines
	for hundred := 0; hundred < 10; hundred++ {

		// Iteration des dizaines
		for tens := 0; tens < 10; tens++ {
			// boucle compteur incrementation unité
			for u := 0; u <= 9; u++ {
				number3 = ('0' + rune(u))
				number2 = ('0' + rune(addNumber2))
				number1 = ('0' + rune(addNumber1))
				// Affichage nombres valides
				if number1 != number2 && number2 != number3 {
					if number2 <= number3 && number1 <= number2 {
						z01.PrintRune(number1)
						z01.PrintRune(number2)
						z01.PrintRune(number3)
						if number1 < ('7') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune('\n')
						}
					}
				}
			}

			// Réinitialisation number3

			if number3 >= 9 {
				number3 = ('0')
				addNumber2++
			}
		} // Réinitialisation number2

		if number3 >= 9 && number2 >= 9 {
			number2 = ('0')
			addNumber2 = 0
			addNumber1++
		}
	}
}
