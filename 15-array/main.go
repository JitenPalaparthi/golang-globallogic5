package main

import (
	"fmt"
)

func main() {
	var arr1 [3]int // int array with the length 3
	fmt.Println(arr1)
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	println()

	// range loop
	// for array, slice range loop returns index and value
	// for maps range loop returns key and value
	// for channels range loop returns value from the channel
	//sum := 0
	for i, v := range arr1 {
		fmt.Println("index:", i, "Value:", v)
		//sum += v
	}
	//fmt.Println("Sum:", sum)

	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println("Sum:", sum)

	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d)

	for i1, v1 := range arr2d {
		println("row--", i1)
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		println()
	}

}

// collection of elements of same type is array
// arrays are fixed length. The length of the array is given at the compiletime.
// an array cannot be extend or shrink at runtime
// index starts from 0
// type inference works for array as well
