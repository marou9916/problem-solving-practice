package main

import (
	"fmt"
	"piscine-go/training"
)

func main() {
	arr1 := []int{90, 70, 20, 80, 50}
	target := 45
	fmt.Println(training.FindPair(arr1, uint(target)))

}
