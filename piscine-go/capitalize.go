package piscine

func Capitalize(s string) string {
	aRune := []rune(s)
	var separator bool
	var upperCase bool
	var lowerCase bool
	for i := 0; i < len(aRune)-1; i++ {
		// if aRune[i] == ' ' || aRune[i] == '+' || aRune[i] == '(' {
		if aRune[i] < 'a' && aRune[i] > 'z' || aRune[i] < 'A' && aRune[i] > 'Z' {
			separator = true
		}
		if separator {
			for u := len(aRune) - 1; u > 0; u-- {
				// if aRune[u] == ' ' || aRune[u] == '+' || aRune[i] == ')' {
				if aRune[i] < 'a' && aRune[i] > 'z' || aRune[i] < 'A' && aRune[i] > 'Z' {
					upperCase = true
				}
				if upperCase {
					if aRune[u+1] >= 'a' && aRune[u+1] <= 'z' {
						aRune[u+1] = aRune[u+1] - 32
						lowerCase = true
					}
					upperCase = false
				}
				if lowerCase {
					for l := aRune[u+2]; l < aRune[i-1]; l++ {
						if aRune[l] >= 'A' && aRune[l] <= 'Z' {
							aRune[l] = aRune[l] + 32
						}
					}
				}
				lowerCase = false
			}
			separator = false
		}

	}
	return string(aRune)
}
