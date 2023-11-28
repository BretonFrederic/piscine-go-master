package piscine

func Compare(a, b string) int {
	// Le résultat sera 0 si a == b, -1 si a < b et +1 si a > b
	var result int
	if a == b {
		result = 0
	} else if a < b {
		result = -1
	} else if a > b {
		result = 1
	}
	return result
}
