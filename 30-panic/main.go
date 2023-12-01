package main

import "fmt"

func main() {
	var num = 10 // to check panic change value from 10 to 0
	fmt.Println(100 / num)

	arr := [3]int{10, 12, 34}
	index := 2 // to check panic change value from 2 to 3 or above
	fmt.Println(arr[index])

	var ptr *int
	*ptr = 100
}
