package main

import (
	"fmt"
	"os"
	"piscine-go/training"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	input := os.Args[1]
	fmt.Println(training.FirstWord(input))
}
