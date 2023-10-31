package piscine

func RecursivePower(nb1 int, power int) int {
	nb := nb1
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}

}
