package training

func FindPair(arr []int, x int) string {

	seen := make(map[int]bool)

	for _, element := range arr {
		if seen[element+x] || seen[element-x] {
			return "Yes"
		}

		seen[element] = true
	}
	return "No"
}
