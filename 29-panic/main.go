package main

import (
	"fmt"
	"os"
)

func main() {

	func() { // Anonymous -1
		func() { // Anonymous -2
			file, err := os.Open("abcd.txt")
			if err != nil {
				//fmt.Println("error:", err.Error())

				//return
				panic("user defined panic: file not opened due to the following error. " + err.Error())
			}
			fmt.Println(file.Name())
		}()

		fmt.Println("Anonymous 2 ends here")

	}()
	fmt.Println("Main func ends here")
}

// panic --> panic func panics the whole call stack by default
// there are runtime panics by the runtime or user defined panics
