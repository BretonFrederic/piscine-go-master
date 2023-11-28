package piscine

func Any(f func(string) bool, a []string) bool {
	var numeric bool
	for i := range a {
		if f(a[i]) {
			numeric = true
			break
		} else {
			numeric = false
		}
	}
	return numeric
}
