package training

import (
	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, table []int, n int) {
	result := n

	for _, value := range table {
		result = f(value, result)
	}

	PrintResult(result)
	z01.PrintRune('\n')
}

func PrintResult(result int) {
	if result < 0 {
		z01.PrintRune('-')
		result *= -1
	}
	PrintDigits(result)
}

func PrintDigits(number int) {
	if number > 9 {
		PrintDigits(number/10)
	}
	z01.PrintRune(rune('0'+(number%10)))
}