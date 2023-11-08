package piscine

func Map(f func(int) bool, arr []int) []bool {
	table := make([]bool, len(arr))

	for x, y := range arr {
		table[x] = f(y)
	}
	return table
}
