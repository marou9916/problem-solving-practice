package piscine

func IterativeFactorial(nb int) int {
	f := 1
	if nb < 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			f = f * i
		}
	}
	return f
}
