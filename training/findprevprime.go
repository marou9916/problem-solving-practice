package training

func FindPrevPrime(e1 int) int {
	for i := e1; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(e1 int) bool {
	for i := 2; i*i <= e1; i++ {
		if e1%i == 0 {
			return false
		}
	}
	return true
}
