package piscine

func Sqrt(number int) int {
	root := 0
	if number < 0 {
		return 0
	}
	for i := 0; i*i <= number; i++ {
		if i*i == number {
			root = i
			break
		}
	}
	return root
}
