package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	fmt.Println("Hello World")
	wg.Add(1)
	go func(*sync.WaitGroup) {
		Print100Odd()
		wg.Done()
	}(wg)

	wg.Add(1)
	go Print100Even(wg, "even--->1")

	wg.Add(3)
	go Print100Even(wg, "even--->2")

	go Print100Even(wg, "even--->3")

	go Print100Even(wg, "even--->4")

	wg.Wait()
} // exit(0)

func Print100Even(wg *sync.WaitGroup, msg string) {
	i := 1
	for {
		if i > 100 {
			wg.Done()
			runtime.Goexit() // exit the goroutine in which it is called
		}
		if i%2 == 0 {
			println(msg, ":", i)
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

// WaitGroup --> Contain a counter
// For each Goroutine , the counter has to be increased.
// For each completion of Goroutine . the counter has to be decreased(This is done by Wg.Done)
// The main waits untile the Counter becomes zero
