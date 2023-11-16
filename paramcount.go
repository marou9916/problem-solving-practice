package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func Paramcount() {
	if len(os.Args[1:]) == 0 {
		z01.PrintRune('0')
	} else {
		z01.PrintRune(rune(len(os.Args[1:]) + '0'))
	}
	z01.PrintRune('\n')
}
