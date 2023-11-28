package piscine

func Abort(a, b, c, d, e int) int {
	nbr := []int{a, b, c, d, e}
	for i := range nbr {
		for j := range nbr {
			if nbr[i] > nbr[j] {
				nbr[i], nbr[j] = nbr[j], nbr[i]
			}
		}
	}
	return nbr[2]
}
