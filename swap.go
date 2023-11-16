package piscine

func Swap(x, y *int) {
	tempx := *x
	*x = *y
	*y = tempx
}
