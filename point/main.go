package main

import "fmt"

type point struct {
	abs int
	ord int
}

func setPoint(ptr *point, x1 int, y1 int) {
	ptr.abs = x1
	ptr.ord = y1
}

func main() {
	points := point{}

	setPoint(&points, 42, 21)

	fmt.Printf("x = %d, y = %d\n", points.abs, points.ord)
}
