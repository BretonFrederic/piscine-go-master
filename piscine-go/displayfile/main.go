/*
package main

import (

	"fmt"
	"os"

)

	func main() {
		if len(os.Args) == 1 {
			fmt.Println("File name missing")
			return
		}
		if len(os.Args) > 2 {
			fmt.Println("Too many arguments")
			return
		}
		file, _ := os.Open(os.Args[1])

		// defer file.Close() // it's important to close the file after reading it

		// create a byte slice to hold the file contents
		data := make([]byte, 15)

		file.Read(data)
		fmt.Printf(" Valeur de data : %v", data)
		fmt.Println(string(data))
	}
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Print(string(file))
		return
	}
}
