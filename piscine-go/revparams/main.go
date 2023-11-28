package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for u := len(os.Args) - 1; u >= 1; u-- {
		stringArgument := []rune(os.Args[u])

		for i := 0; i < len(os.Args[u]); i++ {
			z01.PrintRune(stringArgument[i])
		}
		z01.PrintRune(rune('\n'))
	}
}
