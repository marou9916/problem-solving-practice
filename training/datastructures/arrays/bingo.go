package arrays

func Bingo(arr [25]int, randomNumber int) bool {
	for i := 0; i < 5; i++ {
		counter := 0

		for j := 0; j < 5; j++ {
			if arr[i*5+j] == randomNumber {
				counter++
			}
		}
		if counter == 5 {
			return true
		}
	}

	for i := 0; i < 5; i++ {
		counter := 0

		for j := 0; j < 5; j++ {
			if arr[j*5+i] == randomNumber {
				counter++
			}
		}
		if counter == 5 {
			return true
		}
	}
	return false
}
