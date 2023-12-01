package main

import "fmt"

func main() {

	i := 1
loop:
	if i%2 == 1 {
		print(i, " ")
	}
	i++
	if i <= 10 {
		goto loop
	}
	fmt.Println("Done")

}
