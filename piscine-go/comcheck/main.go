package main

import (
	"fmt"
	"os"
)

func main() {
	var alert bool
	argument := os.Args
	for i := 0; i < len(argument); i++ {
		if argument[i] == "01" || argument[i] == "galaxy" || argument[i] == "galaxy 01" {
			alert = true
			break
		} else {
			alert = false
		}
	}
	if alert == true {
		fmt.Println("Alert!!!")
	} else {
		return
	}
	// fmt.Println(argument)
}
