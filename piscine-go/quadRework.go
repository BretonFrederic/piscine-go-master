package piscine

import "github.com/01-edu/z01"

func Quad(x, y int) {
	// Taille du quad
	Longueur := x
	hauteur := y

	// Symboles des coins, remplacer les 'o' par les charactères souhaités
	coinA1, coinA2 := 'o', 'o'
	coinA3, coinA4 := 'o', 'o'

	// Segments horizontal, remplacer les '-' par les charactères souhaités
	segHorizontalA1 := '-'

	// Segments horizontal, remplacer les '|' par les charactères souhaités
	segVerticalA1 := '|'

	// Intérieur du quad composé d'espace
	surface := ' '

	// return 0 si valeurs négatives ou égales à zéro
	if Longueur < 1 || hauteur < 1 {
		return
	}

	// Boucle principale pour chaque ligne dans l'axe des y on parcours la longueur dans l'axe des x
	for verticale := 1; verticale <= hauteur; verticale++ {
		for ligne := 1; ligne <= Longueur; ligne++ {
			// coin en haut à gauche
			if ligne == 1 && verticale == 1 {
				z01.PrintRune(coinA1)
			}
			// coin en haut à droite
			if ligne > 1 && ligne == Longueur && verticale == 1 {
				z01.PrintRune(coinA2)
			}
			// coin en bas à gauche
			if verticale > 1 && ligne == 1 && verticale == hauteur {
				z01.PrintRune(coinA3)
			}
			// coin en bas à droite
			if verticale > 1 && ligne > 1 && ligne == Longueur && verticale == hauteur {
				z01.PrintRune(coinA4)
			}
			// Segment horizontal première et dernière ligne
			if ligne > 1 && ligne < Longueur {
				if verticale == 1 || verticale == hauteur {
					z01.PrintRune(segHorizontalA1)
				}
			}
			// Segment vertical gauche et droite
			if verticale > 1 && verticale < hauteur {
				if ligne == 1 || ligne == Longueur {
					z01.PrintRune(segVerticalA1)
				} else if ligne > 1 || ligne < Longueur {
					z01.PrintRune(surface)
				}
			}
		}
		z01.PrintRune(rune('\n'))
	}
}
