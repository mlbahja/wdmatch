package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	args1 := os.Args[1] //string 1 / 2
	s1 := []rune(args1)
	args2 := os.Args[2]
	s2 := []rune(args2)

	j := 0
	result := ""
	for i := 0; i < len(s2); i++ {
		if j < len(s1) && s1[j] == s2[i] {
			result += string(s1[j])
			j++
		}
	}
	if result == string(s1) {
		printstr(result)
	}
}
