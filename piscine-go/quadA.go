package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	// Test Si les entier entrer sont négatif ou égal a 0
	if x > 0 && y > 0 {
		// Test Si les deux valeurs sont égal a 1
		if x == 1 && y == 1 {
			z01.PrintRune('o')
			z01.PrintRune('\n')
		} else if x > 1 && y == 1 {
			// Dans le cas où ils ne sont pas a 1 Et que x vaut +1 et y vaut 1
			// Pour x jusqu'a 1 en retirant 1
			for i := x; i >= 1; i-- {
				// Si i vaut x ou 1 on affiche différente valeur
				if i == x || i == 1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		} else if x == 1 && y > 1 {
			// Dans le cas où ils ne sont pas a 1 Et que x vaut 1 et y vaut +1
			// Pour y jusqu'a 1 en retirant 1
			for j := y; j >= 1; j-- {
				// Si j vaut y ou 1 on affiche différente valeur
				if j == y || j == 1 {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('|')
					z01.PrintRune('\n')
				}
			}
		} else {
			// Dans le cas où tout est faux
			// Pour y jusqu'a 1 en retirant 1
			for j := y; j >= 1; j-- {
				// Affiche le caractere du Début & fin de Ligne
				// Si j vaut y ou 1 on affiche différente valeur
				if j == y || j == 1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('|')
				}
				// Pour x jusqu'a 1 en retirant 1
				for i := x - 2; i >= 1; i-- {
					// Affiche le caractere du centre de Ligne
					// Si j vaut y ou 1 on affiche différente valeur
					if j == y || j == 1 {
						z01.PrintRune('-')
					} else {
						z01.PrintRune(' ')
					}
				}
				// Affiche le caractere du Début & fin de Ligne
				if j == y || j == 1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('|')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
