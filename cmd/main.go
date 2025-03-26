package main

import (
	"fmt"
	"piscine-go/training"
)

func main() {
	arr1 := []int{0, -1, 2, -3, 1}
	target := 9
	fmt.Printf("Indice: %v\n", training.LinearSearch(arr1, target))

}
