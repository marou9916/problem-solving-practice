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

	// Diviser par 2 jusqu'à ce que ce ne soit plus divisible
	for number%2 == 0 {
		primeFactors = append(primeFactors, 2)
		number /= 2
	}

	// Diviser par les nombres impairs de 3 à sqrt(number)
	for i := 3; i*i <= number; i += 2 {
		for number%i == 0 {
			primeFactors = append(primeFactors, i)
			number /= i
		}
	}

	// Si le nombre restant est un nombre premier supérieur à 2
	if number > 2 {
		primeFactors = append(primeFactors, number)
	}

	return primeFactors
}
