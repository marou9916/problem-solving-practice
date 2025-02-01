package main

import (
	"fmt"
	"piscine-go/training"
)

func main() {
	arr1 := []int{0, -1, 2, -3, 1}
	fmt.Printf("Résultat final: %v\n", training.FindTriplets(arr1))

}
