package main

import (
	"errors"
	"fmt"
)

func main() {

	var slice1 []int        // only declaration
	slice1 = make([]int, 5) // length= 5, capacity=5

	//slice1 = {10, 11, 12, 13, 14} // once slice is delcared or instantiated , you cannot assign directly like this
	//slice2 := []int{10, 11, 12, 13, 14}
	slice1[0] = 100
	slice1[1] = 101
	slice1[2] = 102
	slice1[3] = 103
	slice1[4] = 104
	fmt.Println("address of first element", &slice1[0], "slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1))
	slice1 = append(slice1, 105) // append a slice
	// if the length(before append) is less than  to capacity , it just appends
	// if the length(before append) is equal to the capacity, then 1. it doubles the capacity
	// and insert the element
	// if it doubles the capacity the slice header is updated with a new pointer
	fmt.Println("address of first element", &slice1[0], "slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1))
	slice1 = append(slice1, 106) // append a slice
	fmt.Println("address of first element", &slice1[0], "slice1:", slice1, "length:", len(slice1), "capacity:", cap(slice1))
	slice2 := slice1 // what is this operation? it is a shallow copy
	// deep copy and shallow copy
	// shallow copy means (in this context) only the slice headers are copied

	slice2[4] = 1004

	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	//slice1 = append(slice1, 107, 108, 109, 110, 111, 112, 113)
	// slice1 and slice2 are divergent. The hearders are entirely different
	slice1, err := add(slice1, 107)
	if err != nil {
		fmt.Println(err)
	}
	slice2[5] = 1005
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	slice3 := make([]int, len(slice1))

	copy(slice3, slice1) // deep copy
}

func add(slice []int, value int) ([]int, error) {
	if slice == nil {
		//return fmt.Errorf("nil slice")
		return slice, errors.New("nil slice")
	}
	slice = append(slice, value)
	return slice, nil
}
