package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args[0]
	nameProgram := []rune(x)
	for i := 2; i < len(x); i++ {
		z01.PrintRune(nameProgram[i])
	}
	z01.PrintRune('\n')
}