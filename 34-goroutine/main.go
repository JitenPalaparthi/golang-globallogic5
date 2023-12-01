package main

import (
	"fmt"
	"runtime"
)

func main() {
	// os.Exit(1) exit without zero
	fmt.Println("Hello World")
	go Print100Odd()
	go Print100Even()
	runtime.Goexit() // on main(this is only on main) this call always return exit 2(non zero). if a process exit without non zero it is considerd as failure execution
} // exit(0)

func Print100Even() {
	i := 1
	for {
		if i > 100 {
			runtime.Goexit() // exit the goroutine in which it is called
		}
		if i%2 == 0 {
			println("Even:", i)
		}
		i++
	}
}

func Print100Odd() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			println("Odd:", i)
		}
	}
}

// runtime.Goexit has two kids behaviours
