package piscine

func CheckTheLine(s string) string {
	// var sumLine int
	line := []rune(s)
	// for c := 0; c <= 8; c++ {
	// 	if line[c] >= '1' && line[c] <= '9' {
	// 		sumLine += int(line[c])
	// 	}
	// }
	for dotCase := 0; dotCase <= 8; dotCase++ {
		if line[dotCase] == '.' {
			line[dotCase] = 49
			for index := 0; index <= 8; index++ {
				if index != dotCase {
					if line[dotCase] == line[index] {
						line[dotCase] += 1
					}
				}
			}
		}
	}
	// fmt.Printf("somme de la ligne = %v, valeur de line index: %v", sumLine, string(line))
	return string(line)
}
