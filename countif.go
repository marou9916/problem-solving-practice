package piscine

func CollatzCountdown(start int) int {
	count := 0
	for start != 1 {
		if start <= 0 {
			return -1
		}
		if start%2 == 0 {
			start = start / 2
			count += 1
		} else {
			start = 3*start + 1
			count += 1
		}
	}
	return count
}
