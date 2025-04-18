package training

func LargestPairSum(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	first, second := -1, -1

	for _, value := range arr {
		if value > first {
			second = first
			first = value
		} else if value > second {
			second = value
		}
	}

	return first + second
}
