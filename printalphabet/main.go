package main

import "github.com/01-edu/z01"

func main() {

  char:='a'
   for i:=1; i<=26; i++{
    z01.PrintRune(char)
    char++
   }
}

