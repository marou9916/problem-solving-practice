package main

import (
	"fmt"
	"piscine-go/training"
)

func main() {
	arr1 := []int{1, 1, 1, 3, 1, 1}
	arr2 := []int{1, 1, 1, 0, 0, 0}
	arr3 := []int{0, 0, 0, 0}
	fmt.Println(training.CountOnes(arr1))
	fmt.Println(training.CountOnes(arr2))
	fmt.Println(training.CountOnes(arr3))
}
