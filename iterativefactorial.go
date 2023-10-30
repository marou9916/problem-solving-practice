package main

import "fmt"

func iterativefactorial(n int) int {
	f := 1
	if n < 0 {
		return 0
	} else {
		for i := 1; i <= n; i++ {
			f = f * i
		}
	}

return f	
}

func main() {
	factorial := 4
	fmt.Printf("%v", iterativefactorial(factorial))
	fmt.Printf("\n")
}
