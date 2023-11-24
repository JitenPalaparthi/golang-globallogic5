package main

import "fmt"

var num1 int // global with respect to main package

const pi = 3.14

func main() {
	var num2 int = 100
	str := "Hello GlobalLogic"
	{ // since it is a separate scope a new stack frame is created
		Greet() // a separate stack frame is created and reffered
		num3 := 101
		fmt.Println(num3)
		{ //since it is a separate scope a new stack frame is created
			num4 := 102
			fmt.Println(num4)
		}
	}
	num5 := 200
	fmt.Println(num2, str, num5)
	fmt.Println("Global variable num1:", num1)
}
func Greet() {
	msg := "Hello GlobalLogic"
	fmt.Println(msg)
}

// What is a process ?
// What is a thread?
// How memory is managed?
