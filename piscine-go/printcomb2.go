package piscine

import "github.com/01-edu/z01"

/*
	func PrintComb2() {
		var tableau string = "00 00"
		digit1 := 0
		digit2 := 0
		digit3 := 0
		digit4 := 0
		addDigit1 := 0

		for i := 0; i < 10000; i++ {
			// if digit1 != digit3 && digit2 != digit4 {
			total := digit1 + digit2 + digit3 + digit4
			total1 := digit1 + digit2
			z01.PrintRune(rune(tableau[4]) + rune(digit4))
			z01.PrintRune(rune(tableau[3]) + rune(digit3))
			z01.PrintRune(rune(tableau[2]))
			z01.PrintRune(rune(tableau[1]) + rune(digit2))
			z01.PrintRune(rune(tableau[0]) + rune(digit1) + rune(addDigit1))
			// Print de virgule espace tant que le total != 36
			if total < 36 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			//}
			digit1 = digit1 + 1

			if digit1 > 9 {
				digit1 = 0
				digit2++
			}
			if digit2 > 9 {
				digit2 = 0
				digit3++
			}
			if digit3 > 9 {
				digit3 = 0
				digit4++
			}
			if total1 >= 18 {
				addDigit1++
			}

		}
	}
*/
func PrintComb2() {
	for number1 := '0'; number1 <= '9'; number1++ {
		for number2 := '0'; number2 <= '9'; number2++ {
			for number3 := '0'; number3 <= '9'; number3++ {
				for number4 := '0'; number4 <= '9'; number4++ {
					if number1 != number3 || number4 > number2 {
						if number1 <= number3 {

							z01.PrintRune(number1)
							z01.PrintRune(number2)
							z01.PrintRune(' ')
							z01.PrintRune(number3)
							z01.PrintRune(number4)
							if number1 == '9' && number2 == '8' {
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}

/*
func PrintComb2() {
	var number1 string = "00"
	digit1 := 0
	digit2 := 0

	var number2 string = "00"
	digit3 := 0
	digit4 := 0

	for i := 0; i < 9000; i++ {
		total1 := digit1 + digit2
		total2 := digit1 + digit2 + digit3 + digit4
		if total2 > 0 {
			z01.PrintRune(rune(number2[1]) + rune(digit4))
			z01.PrintRune(rune(number2[0]) + rune(digit3))
			z01.PrintRune(' ')
			z01.PrintRune(rune(number1[1]) + rune(digit2))
			z01.PrintRune(rune(number1[0]) + rune(digit1))
			// Print de virgule espace tant que le total != 36
			if total2 < 35 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		digit1 = digit1 + 1
		if digit3 >= 9 {
			digit4++
			digit3 = 0
		}
		if total1 >= 18 {
			digit3++
		}

		if digit1 > 9 {
			digit1 = 0
			digit2++
		}
		if digit2 > 9 {
			digit2 = 0
		}

	}
}
*/ /*
func PrintComb2() {
	var tableau string = "00"
	digit1 := 0
	digit2 := 0
	addDigit1 := 0

	for i := 0; i < 102; i++ {
		z01.PrintRune(rune(tableau[1]) + rune(digit2))
		z01.PrintRune(rune(tableau[0]) + rune(digit1) + rune(addDigit1))
		z01.PrintRune(',')
		z01.PrintRune(' ')

		digit1++

		if digit1 > 9 {
			digit1 = 0
			digit2++
		}
		if digit2 >= 9 {
			digit2 = 0
		}

	}
}*/
