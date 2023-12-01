package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [3]int // int array with the length 3
	fmt.Println(arr1)
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	fmt.Println(arr1)

	arr2 := [5]int{10, 11, 12, 13, 14}
	fmt.Println(arr2)
	arr3 := [...]int{110, 111, 112, 2, 3232, 43, 43, 25, 45}
	fmt.Println(arr3)

	// what is the type or array?
	// the type of an array includes its length as well

	fmt.Println("Type of arr1:", reflect.TypeOf(arr1))
	fmt.Println("Type of arr2:", reflect.TypeOf(arr2))
	fmt.Println("Type of arr3:", reflect.TypeOf(arr3))

	//var arr4 [4]int = arr3 //cannot assign; both of them are two different types

	var arr4 [9]int = arr3 // this is okay becase both of them are same type.
	fmt.Println(arr4)

	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d)
	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
	fmt.Println(arr3d)

}

// collection of elements of same type is array
// arrays are fixed length. The length of the array is given at the compiletime.
// an array cannot be extend or shrink at runtime
// index starts from 0
// type inference works for array as well
