package piscine

func Unmatch(a []int) int {
	var pair int
	var counter int
	for i := range a {
		for j := range a {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for u := 1; u < len(a)-1; u++ {
		if (a[u-1]+a[u])%2 == 0 {
			counter++
		}
		if a[u-1] != a[u] {
			return a[u-1]
		}
	}
	if len(a)/counter == 2 {
		pair = -1
	}
	return pair
}
