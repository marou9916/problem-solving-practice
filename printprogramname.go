package piscine

import github.com/01-edu/z01

func PrintProgramName() {
	myProgramName = os.Args[0]
	z01.PrintRune(myProgramName)
}