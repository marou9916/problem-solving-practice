package training

//Now, i'm not gonna calculate the leftMax and the rightMin for each element at each time. Instead, i calculate all of them once
//in order to reduce complexity and improve efficiency
func PartitionPoint(arr []int) int {
	n := len(arr)

	if n < 3 {
		return -1
	}
	
	leftMax := make([]int, n) 
	leftMax[0] = arr[0]

	minRight := make([]int, n)
	minRight[n-1] = arr[n-1]

	for i := 1; i <= n - 2; i++ {
		if arr[i] > leftMax[i-1] {
			leftMax = append(leftMax, arr[i])
		} else {
			leftMax = append(leftMax, leftMax[i-1])
		}
	}

	for i := n-2; i >= 1; i-- {
		if arr[i] < minRight[i+1] {
			minRight = append(minRight, arr[i])
		} else {
			minRight = append(minRight, minRight[i+1])
		}
	}

	for i := 1; i <= n-2; i++ {
		if arr[i] > leftMax[i-1] && arr[i] < minRight[i+1] {
			return i
		}
	}

	return -1
}