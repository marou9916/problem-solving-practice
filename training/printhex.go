package training

import (
	"fmt"
	"strings"
)

func PrintHex(decimalNumber int) {
	hexaDecimalNumerals := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	var hexadecimalNumber strings.Builder
	rests := make([]int, 0, 20)
	n := decimalNumber

	for n > 0 {
		rests = append(rests, n%16)
		n /= 16
	}

	for i := len(rests) - 1; i >= 0; i-- {
		hexadecimalNumber.WriteString(hexaDecimalNumerals[rests[i]])

	}

	fmt.Println(hexadecimalNumber.String())
}
