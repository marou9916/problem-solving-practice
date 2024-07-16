package main

import (
	"fmt"
	"os"
	"piscine-go/training"
	"strconv"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		return
	}

	number, err := strconv.Atoi(arg[0])

	if err != nil || number <= 1 {
		return
	}

	primeFactors := training.GetPrimeFactors(number)

	if len(primeFactors) == 0 {
		fmt.Println(number)
		return
	} else {
		training.Fprime(primeFactors)
		return
	}

}
