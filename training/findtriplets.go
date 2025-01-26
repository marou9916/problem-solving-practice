package training

import "sort"

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

func FindTriplet(arr []int) [][]int {
	var result [][]int

	// Trier le tableau d'abord
	sort.Ints(arr)

	for i := 0; i < len(arr)-2; i++ {
		// Éviter les doublons pour i
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				result = append(result, []int{i, left, right})

				// Éviter les doublons pour left et right
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
