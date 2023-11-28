package piscine

func Fibonacci(index int) int {
	// entrée : un nombre entier n
	// sortie : le terme de rang n de la suite de Fibonacci
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	if index > 1 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
	return index
}
