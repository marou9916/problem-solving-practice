package piscine

func IterativePower(nb1 int, power int) int {
	nb := 1
	i := 0
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		for i <= power {
			nb = nb * nb1
			i++
		}
	}
	return nb
}