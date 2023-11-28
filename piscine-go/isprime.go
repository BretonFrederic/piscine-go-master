package piscine

func IsPrime(n int) bool {
	// Vérifie si le nombre est inférieur ou égal à 1
	if n <= 1 {
		return false
	}

	// Vérifie si le nombre est égal à 2 ou 3
	if n == 2 || n == 3 {
		return true
	}

	// Vérifie si le nombre est pair
	if n%2 == 0 {
		return false
	}

	// Parcourt tous les nombres impairs jusqu'à la racine carrée de n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
