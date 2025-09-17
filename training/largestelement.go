package training

func LargestElement(elements []int) int {
	if len(elements) < 1 {
		return -1
	}

	max := elements[0]

	for i := 1; i < len(elements); i++ {

		if elements[i] > max {
			max = elements[i]
		}
	}
	return max
}
