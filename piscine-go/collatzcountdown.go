package piscine

func CollatzCountdown(start int) int {
	var step int

	for start != 1 {
		if start < 1 {
			return -1
		}
		if start%2 == 0 {
			start /= 2
			step++
		} else if start%2 == 1 {
			start = (3*start + 1)
			step++
		}
	}
	return step
}
