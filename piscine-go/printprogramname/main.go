package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	stringArgument := []rune(os.Args[0])
	sizeArguments := len(stringArgument)
	for i := 2; i < sizeArguments; i++ {
		z01.PrintRune(stringArgument[i])
	}
	z01.PrintRune(rune('\n'))
}
