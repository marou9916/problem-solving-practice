package training

func IsPowerOf2(number int) bool {
	const div = 2
	var numberAfterDiv []int
	numberAfterDiv = append(numberAfterDiv, number)
	if number == 1 {
		return true
	} else if number == 0 || number != 1 && number%div != 0 {
		return false
	} else if number%div == 0 {
		for number != 0 {
			number = number / div
			numberAfterDiv = append(numberAfterDiv, number)
		}
		for i := 1; i <= len(numberAfterDiv)-3; i++ {
			if numberAfterDiv[i]%div != 0 {
				return false
			}
		}
	}
	return true
}
