package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("I am main, I exit now") // this is the first defer so executed at last due to LIFO
	defer println("hello-1")                   // this is executed before the last
	func() {
		defer greet()
		println("I am inside anonymous function;not a deffered one so executed at firsts in this stack frame")
	}()
	// this is executed at first(from defers perspective)
	fmt.Println("Hello World! Start of main")

	f, err := os.Open("abcd.txt")
	if err != nil {
		panic("user defined panic: " + err.Error())
	}
	defer f.Close()
	defer fmt.Println("Will I be executed becase I am defered one?")

}

// defer defers the execution to the end of the caller
// defer calls are executed even during panics

func greet() {
	println("Hello Global Logic Folks!")
}
