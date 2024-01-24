package training

import "os"

func MultTab(a_number_strictly_positive int) {
	var result int
	for i := 1; i <= 9; i++ {
		result = a_number_strictly_positive * i
		os.Stdout.WriteString(Itoa(i) + "*" + os.Args[1] + "=" + Itoa(result) + "\n")
	}
}
