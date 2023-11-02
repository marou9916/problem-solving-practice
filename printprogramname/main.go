package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args[0]
	nameProgram := []rune(x)
	for i := 0; i <= len(x)-1; i++ {
		z01.PrintRune(nameProgram[i])
	}
}
