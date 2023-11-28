package piscine

func TrimAtoi(s string) int {
	var byteArray []byte
	var intFromBytes int

	for i := range s {
		// Ajout des char compris entre 0 et 9 dans byteArray[]
		if s[i] >= '0' && s[i] <= '9' {
			byteArray = append(byteArray, s[i])
		}
	}
	// '1' dans ASCII = 49 en décimal donc 49 - 48 = 1 en décimal
	// On multiplie par 10 à chaque boucle pour obtenir 1, 12, 123, 1234, 12345
	for i := range byteArray {
		intFromBytes = intFromBytes*10 + int(byteArray[i]-48)
	}
	// si signe - (sd-) alors on multiplie par -1 pour obtenir nombre négatif
	for i := range s {
		if s[:i] == "sd-" {
			intFromBytes *= -1
		}
	}
	return intFromBytes
}
