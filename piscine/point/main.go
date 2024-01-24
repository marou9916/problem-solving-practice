package main

import "github.com/01-edu/z01"

type point struct {
	abs int
	ord int
}

func setPoint(ptr *point, x1 int, y1 int) {
	ptr.abs = 42
	ptr.ord = 21
}

func main() {
	xStr := "x = "
	yStr := "y = "
	points := &point{}

	setPoint(points, points.abs, points.ord)

	xTab := []rune{}
	yTab := []rune{}

	xVal := points.abs
	yVal := points.ord

	for xVal != 0 {
		xTab = append(xTab, rune(xVal%10))
		xVal /= 10
	}
	for _, val := range xStr {
		z01.PrintRune(rune(val))
	}
	for i := len(xTab) - 1; i >= 0; i-- {
		z01.PrintRune(xTab[i] + 48)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for yVal != 0 {
		yTab = append(yTab, rune(yVal%10))
		yVal /= 10
	}
	for _, val := range yStr {
		z01.PrintRune(rune(val))
	}
	for i := len(yTab) - 1; i >= 0; i-- {
		z01.PrintRune(yTab[i] + 48)
	}
	z01.PrintRune('\n')
}
