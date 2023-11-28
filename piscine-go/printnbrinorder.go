package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	number := n
	var rest int
	var newNumber int
	var arrayNbr []int

	if number < 0 {
		number = number * -1
	}
	for number > 0 {
		rest = number % 10
		newNumber = number / 10
		arrayNbr = append(arrayNbr, rest)
		number = newNumber
	}

	for l := range arrayNbr {
		for i := range arrayNbr {
			if arrayNbr[l] < arrayNbr[i] {
				arrayNbr[l], arrayNbr[i] = arrayNbr[i], arrayNbr[l]
			}
		}
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for r := range arrayNbr {
		z01.PrintRune(rune(arrayNbr[r] + '0')) // rune(1+'0') --> (1+ 48)
	}

	// fmt.Println(arrayNbr)
}

/*

if n == 0 {
	z01.PrintRune('0')
}
 var digits [10] int
 for n > 0{
	digits[n%10]++
	n/=10
 }
 for i := 0; i<10;i++{
	for j := 0; j< digits[i];j++{
		z01.PrintRune(rune(i)+'0')
	}
 }*/
