package main

import (
	"fmt"
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	userArgs := os.Args
	var aString string
	if len(os.Args) < 9 {
		fmt.Println("Error")
		return
	}
	// Nombre de string limite 9 (y)
	for i := 1; i <= 9; i++ {
		// longueur de chaque strings limite 9 charactères
		if len(userArgs[i]) <= 9 {
			// os.Args index, vers string
			aString = userArgs[i]
			var arrayString []rune
			// conversion de string vers tableau de rune
			for u := 0; u <= 8; u++ {
				arrayString = []rune(piscine.CheckTheLine(aString))
				// Affichage de la rune courante suivi d'un espace
				z01.PrintRune(arrayString[u])
				z01.PrintRune(' ')
			}
			// Retour à la ligne avant d'afficher la ligne suivante
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}
