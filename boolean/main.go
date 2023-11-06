package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	var lengthOfArg int
	var EvenMsg string
	var OddMsg string
	lengthOfArg = len(os.Args) - 1
	EvenMsg = "I have an even number of arguments"
	OddMsg = "I have an odd number of arguments"
	if !isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
