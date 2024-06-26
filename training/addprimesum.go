package training

func AddPrimeSum(n *int) {
	if *n > 0 {
		primeSum := 0

		for i := *n; i >= 2; i-- {
			if IsPrime(i) {
				primeSum += i
			}
		}

		*n = primeSum

	} else {
		*n = 0
	}
}
