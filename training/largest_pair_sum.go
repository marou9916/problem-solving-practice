package training

func LargestPairSum(arr []int) int {
	max := arr[0]
	largestSum := -1
	if len(arr) < 2 {
		return -1
	}

	for _, element := range arr {
		if element > max {
			max = element
		}
	}

	for i := range arr {
		if arr[i] == max {
			continue
		}
		if max+arr[i] > largestSum {
			largestSum = max + arr[i]
		}
	}

	return largestSum
}
