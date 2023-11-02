package piscine

import {
	"fmt"
	"os"
}

func PrintProgramme() {
	myProgramName = os.Args[0]
	fmt.Println(myProgramName)
}