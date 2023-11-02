package piscine

func Compare(a string, b string) int {
	if len(a) < len(b) {
		for i := 0; i <= len(a); i++ {
			if a[i] == b[i] {
				return -1
			}
		}
	} else if len(a) == len(b) {
		for i := 0; i <= len(a); i++ {
			if a[i] == b[i] {
				return 0
			}
		}
	} else if len(a) > len(b) {
		return -1
	}
	return 0
}
