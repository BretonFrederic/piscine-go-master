package piscine

import "fmt"

/*func CheckTheSudoku(array []string) {
	var aRune []rune
	for s := 0; s <= 8; s++ {
	}

	// boucle des 9 arguments
	for i := 1; i <= 9; i++ {
		// Initialisation d'une string avec l'argument courant
		aString := array[i]
		// Initialisation d'un tableau de runes avec la string
		aRune = []rune(aString)
		// Recherche d'un point dans le tableau de runes
		for dotCase := 0; dotCase <= 8; dotCase++ {
			if aRune[dotCase] == '.' {
				// Change la valeur du point en '1'
				aRune[dotCase] = 49
				// Parcourir l'index 0 à 8 du tableau de runes
				for index := 0; index <= 8; index++ {
					// Vérifie qu'on ne compare pas le même index
					if index != dotCase {
						// Compare la nouvelle valeur de dotCase avec les valeurs existantes
						if aRune[dotCase] == aRune[index] {
							aRune[dotCase] += 1
						}
					}
				}
			}
		}
		aString = string(aRune)
		fmt.Println(aString)
	}
}*/

func CheckTheSudoku(array []string) {
	var aRune []rune

	// boucle des 9 arguments
	for i := 1; i <= 9; i++ {
		// Initialisation d'une string avec l'argument courant
		aString := array[i]
		// Initialisation d'un tableau de runes avec la string
		aRune = []rune(aString)
		// Recherche d'un point dans le tableau de runes
		for dotCase := 0; dotCase <= 8; dotCase++ {
			if aRune[dotCase] == '.' {
				// Change la valeur du point en '1'
				aRune[dotCase] = 49
				// Parcourir l'index 0 à 8 du tableau de runes
				for index := 0; index <= 8; index++ {
					// Vérifie qu'on ne compare pas le même index
					if index != dotCase {
						// Compare la nouvelle valeur de dotCase avec les valeurs existantes
						if aRune[dotCase] == aRune[index] {
							aRune[dotCase] += 1
						}
					}
				}
			}
		}
		aString = string(aRune)
		fmt.Println(aString)
	}
}

// for c := 0; c <= 8; c++ {
// 	if line[c] >= '1' && line[c] <= '9' {
// 		sumLine += int(line[c])
// 	}
// }
