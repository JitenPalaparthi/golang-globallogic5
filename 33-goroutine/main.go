package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	go Print100Odd()
	go Print100Even()
	time.Sleep(time.Millisecond * 1)
} // exit(0)

func Print100Even() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			println("Even:", i)
		}
	}
}

func Print100Odd() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			println("Odd:", i)
		}
	}
}

// 1. go is keyword that is used to create a goroutine
// 2. go runtime creates threads(p) by default the number threads are equal to the number of processes
// 	runtime.GOMAXPROC
// 3. There is global queue that is maintained by go runtime ..when a request come
// 		to create a new goroutine is placed in that Global Queue
// 4. A thread pics that goroutine and place it in the local queue. Local queue is for each thread.
// 5. So behind the scene all gorotuines run on top of the OS threads.
// 6. Is a gorountine is busy(may be performing any io operation) and blocked,
//    the threads is also blocked.On that thread there may be other goroutines are running.
// 	   Go rountine understands that the thread is blocked and place all the goroutines that are being executed
//     on a blocked thread to the Global Queue
// 	   If all threads are blocked then go runtime creates additional threads.
// 7. if one or some of the threads are busy but some of them are free then there is aconcept
// call work stealing. Free threads steals work from busy threads(generally from local queue) if there is nothing
// from the Global queue.

// main is also a goroutine
// by default no goroutine waits for other gorountine to complete its execution
// we cannot guarantee the order of execution of Goroutines
