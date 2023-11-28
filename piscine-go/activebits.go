package piscine

func ActiveBits(n int) int {
	statusByte := 0
	for n != 0 {
		if n%2 == 1 {
			statusByte++
		}
		n = n / 2
	}
	return statusByte
}
