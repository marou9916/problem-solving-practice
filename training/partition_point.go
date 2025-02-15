package training

//This time, i'm gonna calculate all the right minimum, store them into a table, then change the left maximum dynamically
func PartitionPoint(arr []int) int {
	n := len(arr)
	if n < 3 {
		return -1
	}

	minRight := make([]int, n)
	minRight[n - 2] = arr[n - 2] 
	
	for i := n-3; i >= 0; i-- {
		if minRight[i+1] < arr[i+1] {
			minRight[i] = minRight[i+1]
		} else {
			minRight[i] = arr[i+1]
		}
	}

	maxLeft := arr[0]

	for j := 1; j <= n - 2; j++ {
		if arr[j] > maxLeft && arr[j] < minRight[j] {
			return j
		}
		if arr[j] > maxLeft {
			maxLeft = arr[j]
		} 
	}

	return -1
}