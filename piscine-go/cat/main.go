package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argument := os.Args
	for i := 1; i < len(argument); i++ {
		file, err := ioutil.ReadFile(argument[i])
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(file))
	}
	fmt.Printf("\n")
}
