package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("I am main, I exit now") // this is the first defer so executed at last due to LIFO
	defer greet()

	//defer recoverme() // all recovers must be deffered

	func() {
		defer recoverme()
		fmt.Println("It is func1 anoymous function")
		num := 0
		println(100 / num)
		fmt.Println("How ever I will not be executed")
	}()

	fmt.Println("Can I be called when there is a panic?")

}

// defer defers the execution to the end of the caller
// defer calls are executed even during panics

func greet() {
	println("Hello Global Logic Folks!")
}

func recoverme() {
	if p := recover(); p != nil { // when d will not be nil is when there is panic
		// todo all things to make sure to gracefully handle things
		fmt.Println("There is a panic but I am trying to minimize the problem:here is panic info", p)
	}
}
