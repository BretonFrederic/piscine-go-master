package piscine

func UltimateDivMod(a *int, b *int) {
	divide := *a
	modulo := *b
	*a = divide / modulo
	*b = divide % modulo
}
