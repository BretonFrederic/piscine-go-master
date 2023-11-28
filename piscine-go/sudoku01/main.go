package main

import (
	"fmt"
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	userArgs := os.Args
	// Minimum de 9 arguments
	if len(os.Args) < 9 {
		fmt.Println("Error")
		return
	}
	// longueur de chaque strings limite 9 charactères
	piscine.CheckTheSudoku(userArgs)
	z01.PrintRune('\n')
}
