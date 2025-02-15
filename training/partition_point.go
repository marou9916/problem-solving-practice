package training

//This time, i'm gonna calculate all the right minimum, store them into a table, then change the left maximum dynamically
func PartitionPoint(arr []int) int {
	n := len(arr)
	if n < 3 {
		return -1
	}

	minRight := make([]int, n)
	minRight[n - 1] = arr[n - 1] 
	
	for i := n-2; i >= 0; i-- {
		minRight[i] = min(minRight[i+1], arr[i+1])
	}

	maxLeft := arr[0]

	for j := 1; j <= n - 2; j++ {
		if arr[j] > maxLeft && arr[j] < minRight[j] {
			return j
		}
		maxLeft = max(maxLeft, arr[j]) 
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}