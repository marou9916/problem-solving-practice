package main

import (
	"fmt"
	"os"
	"piscine-go/training"
)

func main() {
	args := os.Args[1:]
	s := args[0]
	if len(args) == 1 {
		fmt.Println(training.RepeatAlpha(s))
	} else {
		return
	}

}
