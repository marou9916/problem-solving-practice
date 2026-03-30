package training

func GiveTheTableOfMultiplication(n int) [10]int {
	var multiplesOfTen [10]int

	for i := 1; i <= 10; i++ {

		multiplesOfTen[i-1] = n * i
	}
	return multiplesOfTen
}
