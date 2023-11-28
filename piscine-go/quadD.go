package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	// Test Si les entier entrer sont négatif ou égal a 0
	if x > 0 && y > 0 {
		// Test Si les deux valeurs sont égal a 1
		/*
			/
		*/
		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
		} else if x > 1 && y == 1 {
			/*
				/***\
			*/
			// Dans le cas où ils ne sont pas a 1 Et que x vaut +1 et y vaut 1
			// Pour x jusqu'a 1 en retirant 1
			for i := x; i >= 1; i-- {
				// Si i vaut x ou 1 on affiche différente valeur
				if i == x { // Affiche le caractere du Début
					z01.PrintRune('A')
				} else if i == 1 { // Affiche le caractere de fin
					z01.PrintRune('C')
				} else { // Affiche le caractere central
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		} else if x == 1 && y > 1 {
			/*
				/
				*
				*
				*
				\
			*/
			// Dans le cas où ils ne sont pas a 1 Et que x vaut 1 et y vaut +1
			// Pour y jusqu'a 1 en retirant 1
			for j := y; j >= 1; j-- {
				if j == y { // Affiche le caractere du Début
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else if j == 1 { // Affiche le caractere de fin
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else { // Affiche le caractere central
					z01.PrintRune('B')
					z01.PrintRune('\n')
				}
			}
		} else {
			/*
				/***\
				*   *
				\**°/
			*/
			// Dans le cas où tout est faux
			// Pour y jusqu'a 1 en retirant 1
			for j := y; j >= 1; j-- {
				if j == y { // Affiche coin superieur gauche
					z01.PrintRune('A')
				} else if j == 1 { // Affiche coin inferieur gauche
					z01.PrintRune('A')
				} else { // Affiche coin centre gauche
					z01.PrintRune('B')
				}
				// Pour x jusqu'a 1 en retirant 1
				for i := x - 2; i >= 1; i-- {
					if j == y || j == 1 { // Affiche caractere du centre
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
				if j == y { // Affiche coin superieur droit
					z01.PrintRune('C')
				} else if j == 1 { // Affiche coin inferieur droit
					z01.PrintRune('C')
				} else { // Affiche coin centre droit
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
			}
		}
	}
}
