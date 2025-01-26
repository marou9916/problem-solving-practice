package training

func CountOnes(arr []int) int {

	counter := 0

	for _, value := range arr {
		if value < 0 || value > 0 {
			break
		} else if value > 1 || value < 1 {
			break
		}
		if value == 1 {
			counter++
		} else {
			break
		}
	}
	return counter
}
