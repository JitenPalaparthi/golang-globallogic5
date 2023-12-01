package main

import "fmt"

func main() {
	// nested loop
outer: // label
	for i := 1; i <= 5; i += 1 { // outer loop

		for j := 2; j <= 5; j += 1 { // inner loop

			if i == 3 && j == 4 {
				//break // it breaks only the inner loop but not out loop
				break outer
			}
			fmt.Println("i:", i, ": j:", j)
		}

	}
	println()
	k := 2
	for i, j := 1, 2; (i <= 4 && j <= 6) || k <= 3; i, j = i+1, j+2 {
		fmt.Println("i:", i, ": j:", j, "k:", k)
		k++
	}
}
