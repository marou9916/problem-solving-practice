package piscine

func MakeRange(min, max int) []int {
	var slice []int
	if min >= max {
		return nil
	} else {
		slice = make([]int, max-min)
		for i := 0; i < len(slice); i++ {
			slice[i] = min + i
		}
		return slice
	}
}
