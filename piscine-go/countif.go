package piscine

func CountIf(f func(string) bool, tab []string) int {
	var counter int
	for i := range tab {
		if f(tab[i]) {
			counter++
		}
	}
	return counter
}
