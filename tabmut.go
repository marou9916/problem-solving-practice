package piscine

import "github.com/01-edu/z01"

func TabMut() {
	var result string
	n := 9
	for i := 0; i <= 10; i++ {
		p := n * i
		result = Itoa(n) + " x " + Itoa(i) + " = " + Itoa(p) + "\n"
		for _, r := range result {
			z01.PrintRune(r)
		}
	}

}

func Itoa(nb int) string {
	if nb == 0 {
		return ""
	}
	nbString := nb % 10
	return Itoa(nb/10) + string(rune(nbString))
}
