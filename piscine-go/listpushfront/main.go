package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{} // link pointe vers l'adresse de la liste de type List

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
