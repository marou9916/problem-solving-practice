package training

import (
	"fmt"
	"strings"
)

func Fprime(tab []int) {
	var combination []string

	for _, element := range tab {
		combination = append(combination, fmt.Sprintf("%d", element))
	}

	fmt.Println(strings.Join(combination, "*"))

}

func GetPrimeFactors(number int) []int {
	primeFactors := []int{}

	for i := 2; i <= number; i++ {
		if IsPrime2(i) && number%i == 0 {
			primeFactors = append(primeFactors, i)
		}
	}

	return primeFactors
}

func IsPrime2(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
