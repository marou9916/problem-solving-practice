package main

import (
	"fmt"
	"os"
	"piscine-go/training"
)

//import "fmt"

func main() {
	fmt.Println(training.Inter(os.Args[1], os.Args[2]))
}
