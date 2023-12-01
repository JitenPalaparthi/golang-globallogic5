package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var slice1 []int // declaration of the slice; not instantiation

	slice1 = make([]int, 5) // instantiation of the slice
	fmt.Println(slice1)

	slice2 := make([]int, 19) // shorthand declaration of the slice
	fmt.Println(slice2)

	for i := range slice2 {
		slice2[i] = rand.Intn(100)
	}
	fmt.Println(slice2)

	slice3 := []int{12, 34, 43, 65, 45, 46, 5, 77, 764, 5, 56}
	fmt.Println(slice3)

	slice4 := make([]int, 3, 10) // shorthand declaration of the slice
	fmt.Println(slice4)

	for i := range slice4 {
		slice4[i] = rand.Intn(100)
	}
	fmt.Println(slice4)
	fmt.Println("Sum of slice1:", SumOf(slice1))
	fmt.Println("Sum of slice2:", SumOf(slice2))
	fmt.Println("Sum of slice3:", SumOf(slice3))
	fmt.Println("Sum of slice4:", SumOf(slice4))
	fmt.Println("Length of slice1", len(slice1), "Capacity of slice1:", cap(slice1))
	fmt.Println("Length of slice2", len(slice2), "Capacity of slice2:", cap(slice2))
	fmt.Println("Length of slice3", len(slice3), "Capacity of slice3:", cap(slice3))
	fmt.Println("Length of slice4", len(slice4), "Capacity of slice4:", cap(slice4))

}

// slice is similar array but can grow or shrink at rumtime
// an array can be converted to a slice with a simple notation
// make is a built in function used to allocate slice
// slice contains len and also capacity
// can check nil on slice variables
// slice can be allocated on stack or on heap based on various other requirements
// type inference for slice as well
// make built in function is used for allocation memory to slice, map and channel
// 	or any underlining type of those three types

// The below is mimic slice header; not a real slice header
//  Slice{
// 	ptr // pointer to the memory region
// 	len // length of the slice
// 	cap // capacity of the slice
// }

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
