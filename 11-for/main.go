package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	println()

	i := 1        // init
	for i <= 10 { //condition
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
		i++ //post
	}
	println()
	j := 0 //init
	for {
		j++ // incrementor
		if j == 100 {
			break // breaks the whole loop
		}
		if j%2 == 1 { // if it is an odd number
			continue // continue to the next iteration
		}
		fmt.Print(j, " ")
	}
}

// There are only two kids of loops in go.
// 1. for
// 2. for range loop
