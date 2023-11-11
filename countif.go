package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, x := range tab {
		if f(x) {
			counter += 1
		}
	}
	return counter
}
