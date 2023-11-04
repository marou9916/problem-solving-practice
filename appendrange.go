package piscine

func AppendRange(min, max int) []int {
	var array []int
	if min >= max {
		return nil
	} else {
		for i := min; i <= max-1; i++ {
			array = append(array, i)
		}
	}
	return array
}
