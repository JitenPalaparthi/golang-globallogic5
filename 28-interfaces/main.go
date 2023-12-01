package main

import (
	"fmt"

	"github.com/JitenPalaparthi/allshapes/shapes/rect"
	"github.com/JitenPalaparthi/allshapes/shapes/square"
)

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IShape interface {
	IArea
	IPerimeter
}

// class A:IArea

// class A extends IArea

func Area(iarea IArea) {
	fmt.Println("Area is ", iarea.Area())
}

func Perimeter(iPerimeter IPerimeter) {
	fmt.Println("Perimeter is ", iPerimeter.Perimeter())
}

func Shape(ishape IShape) {
	Area(ishape)
	Perimeter(ishape)
}

func main() {

	r1 := &rect.Rect{L: 12.45, B: 23.56}
	var s1 square.Square = 12.34

	c1 := new(Cuboid)
	c1.L = 12.34
	c1.B = 12.55
	c1.H = 14.56

	var iShape IShape
	// Area(r1)
	// Perimeter(r1)
	iShape = r1
	Shape(iShape)

	iShape = s1
	Shape(iShape)

	iShape = c1
	Shape(iShape)
}

// What is an interface
//

type Cuboid struct {
	L, B, H float32
}

func (c *Cuboid) Area() float32 {
	return c.B * c.H * c.L
}

func (c *Cuboid) Perimeter() float32 {
	return 4 * (c.B + c.H + c.L)
}
