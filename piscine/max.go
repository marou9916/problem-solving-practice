package piscine

func Max(a []int) int {
	var max int
	max = a[0]
	for _, x := range a {
		if max < x {
			max = x
		}
	}
	return max
}
