package training

func BinarySearch(input []int, target int) int {
	var middleIndex int
	left, right := 0, len(input)-1

	for left < right {
		middleIndex = (left + right) / 2
		
		if input[middleIndex] == target {
			return target
		} else if input[middleIndex] < target {
			left = middleIndex
		} else {
			right = middleIndex
		}

	}

	return -1
}
