package main

import (
	"fmt"
	"piscine-go/training"
)

func main() {
	arr1 := []int{5, 20, 3, 2, 50, 80}
	target := 78
	fmt.Println(training.FindPair(arr1, target))

}
