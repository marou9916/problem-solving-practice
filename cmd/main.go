package main

import (
	"fmt"
	"piscine-go/training/datastructures/arrays"
)

func main() {
	arr := []int{1, -2}
	mn, err := arrays.MissingNumber(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mn)
	}
}
