package piscine

func Compare(a string, b string) int {
	minimaLength := len(a)
	if len(b) < minimaLength {
		minimaLength = len(b)
	}

	for i := 0; i < minimaLength; i++ {
		if a[i] < b[i] {
			return -1
		} else {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
