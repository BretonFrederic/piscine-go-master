package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	var resx [2]int
	resx[0] = points.x % 10
	resx[1] = points.x / 10

	var resy [2]int
	resy[0] = points.y % 10
	resy[1] = points.y / 10

	aStringx := "x = "
	aStringy := "y = "

	// fmt.Printf("x = %v, y = %v\n", resx[0], points.y)
	for _, ch := range aStringx {
		z01.PrintRune(ch)
	}

	z01.PrintRune(rune(resx[1] + '0'))
	z01.PrintRune(rune(resx[0] + '0'))
	z01.PrintRune(',')
	z01.PrintRune(' ')

	for _, ch := range aStringy {
		z01.PrintRune(ch)
	}
	z01.PrintRune(rune(resy[1] + '0'))
	z01.PrintRune(rune(resy[0] + '0'))
	z01.PrintRune('\n')
}
