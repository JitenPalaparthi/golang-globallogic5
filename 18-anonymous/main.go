package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello World")
	}()

	// this is assigning an anonymous func to a variable
	hw := func() {
		fmt.Println("Hello World")
	}
	hw()

	// this function to be executed
	func(msg string) {
		fmt.Println(msg)
	}("Hello GlobalLogic")

	hwm := func(msg string) {
		fmt.Println(msg)
	}

	hwm("Hello Global Logic Folks!")

	area1 := func(l, b float32) float32 {
		return l * b
	}(12.45, 14.45)

	area2 := func(l, b float32) float32 {
		return l * b
	}

	a2 := area2(12.45, 14.45)

	fmt.Println(area1)
	fmt.Println(a2)

	r1 := Caluclate(12.34, 13.43, func(f1, f2 float32) float32 {
		return f1 + f2
	})
	fmt.Println("Sum:", r1)

	var sub func(float32, float32) float32

	sub = func(f1, f2 float32) float32 {
		return f1 - f2
	}

	r2 := Caluclate(12.34, 13.43, sub)
	fmt.Println("Sub:", r2)

	r3 := Caluclate(12.34, 13.43, Multiplication)
	fmt.Println("Multiplication:", r3)

}

func Caluclate(a, b float32, fn func(float32, float32) float32) float32 {
	return fn(a, b)
}

func Multiplication(f1, f2 float32) float32 {
	return f1 * f2
}
