package piscine

func Max(a []int) int {
	for i := range a {
		for j := range a {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a[len(a)-1]
}
