package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argument := os.Args
	if len(argument) != 4 {
		return
	}

	n1 := argument[1]
	n2 := argument[3]

	nbr1, _ := strconv.Atoi(n1)
	nbr2, _ := strconv.Atoi(n2)
	operator := argument[2]
	var result int

	if (n1 >= "0" && n1 <= "9") && (n2 >= "0" && n2 <= "9") {
		switch operator {
		case "+":
			result = nbr1 + nbr2
		case "-":
			result = nbr1 - nbr2
		case "*":
			if n2 == "-1" {
				result = (nbr1 * nbr2)
				result *= -1
			} else {
				result = (nbr1 * nbr2)
			}
		case "/":
			if nbr2 != 0 {
				result = nbr1 / nbr2
			} else {
				fmt.Print("No division by 0")
				return
			}
		case "%":
			if nbr2 != 0 {
				result = nbr1 / nbr2
			} else {
				fmt.Print("No modulo by 0")
				return
			}
		default:
			return
		}
	}
	fmt.Print(result)
}
