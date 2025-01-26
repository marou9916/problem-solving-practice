package training


func FindTriplets(arr []int) [][]int {
	var result [][]int

	for i := 0; i <= len(arr)-2; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			for k := j + 1; k <= len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					result = append(result, []int{i, j, k})
				}
			}
		}
	}

	return result
}


