package piscine

func StrLen(s string) int {
	aRune := []rune(s)
	count := 1
	for i := 1; i < len(aRune); i++ {
		count++
	}
	return count
}
