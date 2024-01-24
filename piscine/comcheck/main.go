package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, i := range args {
		if i == "01" || i == "galaxy" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
