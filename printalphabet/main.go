package main

import "github.com/01-edu/z01"

func main() {
    char := 'a'
    for i := 0; i < 26; i++ {
        z01.PrintRune(char)
        char++
    }
    z01.PrintRune('\n')
}