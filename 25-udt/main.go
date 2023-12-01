package main

import "fmt"

func main() {

	r1 := Rect{L: 101.13, B: 103.45}
	fmt.Println(Area(r1))
	fmt.Println("Area of Rect r1:", (&r1).Area())
	fmt.Println("Perimeter of Rect r1:", r1.Perimeter())
	fmt.Println("Area inside Rect r1:", r1.A)
	fmt.Println("Perimeter inside Rect r1:", r1.P)

}

type Rect struct {
	L, B float32
	A, P float32
}

// normal receiver r.A is 0. Because r is a call by value
// so r allocates a ew varaible of type rect and deallocates as soon as the func call ends
//
//	func (r Rect) Area() float32 {
//		r.A = r.L * r.B
//		return r.A
//	}

// r is a pointer receiver so it is pass or call by reference
func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// Area is a function
func Area(r Rect) float32 {
	return r.L * r.B
}

// Methods are associated with types
// a type can be a struct or any user defined type from the underlining type
// a method contains a receiver(only one)
// a normal receiver is call or pass by value
// to maintain the state of an object during func calls , use pointer receivers
