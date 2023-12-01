package main

import (
	"fmt"

	"github.com/JitenPalaparthi/allshapes/shapes/rect"
	sqr "github.com/JitenPalaparthi/allshapes/shapes/square"
)

func main() {
	r1 := rect.Rect{L: 10123.123, B: 232.34}
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Printf("Area of r1: %.3f\n", a1)
	fmt.Println("Perimeter of r1:", p1)

	//r1.display() // unexported

	var s1 sqr.Square = 25.34

	a2 := s1.Area()
	p2 := s1.Perimeter()

	fmt.Printf("Area of s1: %.3f\n", a2)
	fmt.Println("Perimeter of s1:", p2)

}
