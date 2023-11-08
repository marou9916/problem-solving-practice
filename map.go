package piscine

func Map(f func(int) bool, arr []int) []bool {
	var table []bool
	for _, entier := range arr {
		table = append(table, f(entier))
	}
	return table
}
