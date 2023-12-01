package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 10)

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println(slice1)

	fmt.Println(SumOf(slice1))

	arr1 := [5]int{10, 34, 22, 4, 45}
	// array can be convereted to slice using arr[:] notation
	fmt.Println(SumOf(arr1[:])) // arr[:] converts an array to a slice

	slice2 := slice1[2:5] // :upto 5 but is not equal to 5
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice3 := slice1[3:8]
	fmt.Println(slice3)

	slice4 := slice1[4:] // 4th to the end
	fmt.Println(slice4)
	slice5 := slice1[:5]
	fmt.Println(slice5)

	slice6 := slice1[:] // all elements
	fmt.Println(slice6)

	fmt.Println("Variadic")
	fmt.Println(Sum())
	fmt.Println(Sum(10, 20))
	fmt.Println(Sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150))
	//fmt.Println(Sum(arr1...)) // array cannot be directly passed as variadic argument
	fmt.Println(Sum(arr1[:]...))
	fmt.Println(Sum(slice1...))

	arr2 := [3]int{10, 12, 13}
	slice7 := append(arr1[:], arr2[:]...)

	fmt.Println(Sum(slice7...))

}

// nums is a variadic parameter
// variadic parameter is identified with eclipse symbol
// to create a param ...int
// to pass as an argument of a slice slice...
// variadic parameter must be the last parameter
// there can only be one variadic parameter
// can use range loop on variadic
// func Sum(val int, nums ...int) (sum int) { // it is fine
// func Sum( nums ...int,val int) (sum int) { // it is not fine. The variadic must be the last param
func Sum(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumOf(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}
	return sum
}
