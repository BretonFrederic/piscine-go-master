package piscine

func SplitWhiteSpaces(s string) []string {
	var arrayString []string
	var sizeArrayString int
	position := 0
	// Taille du tableau en fonction du nombre d'espace dans s
	for i := range s {
		if s[i] == ' ' {
			sizeArrayString++
		}
	}
	// Attribution de la tranche de charactères si ' '
	for i := range s {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if s[position] == ' ' || s[position] == '\t' || s[position] == '\n' {
				position = i + 1
			} else {
				arrayString = append(arrayString, s[position:i])
				position = i + 1
				// fmt.Printf("mot :%v\n", arrayString[0:])
				// fmt.Printf("position :%v\n", position)
			}
		}
	}
	arrayString = append(arrayString, s[position:])
	return arrayString
}
