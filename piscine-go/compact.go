package piscine

func Compact(ptr *[]string) int {
	var counter int
	var newString []string
	aString := *ptr
	for i, ch := range aString {
		if aString[i] != "" {
			newString = append(newString, ch)
			counter++
		}
	}
	*ptr = newString
	return counter
}
