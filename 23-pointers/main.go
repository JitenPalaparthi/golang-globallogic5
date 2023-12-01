package main

import "fmt"

func main() {

	var num1 int = 100
	var ptr1 *int
	if ptr1 == nil {
		println("nil pointer; assigning num1")
		ptr1 = &num1
	}

	Incr(num1)
	fmt.Println("Incr", num1)

	IncrP(&num1)
	fmt.Println("IncrP", num1)

	IncrP(ptr1)
	fmt.Println("IncrP", num1)

	*ptr1 = 200 // dereference. The value at the address gets changed
	fmt.Println("num1", num1)

	var ptr2 **int = &ptr1
	var ptr3 ***int = &ptr2

	fmt.Println("Address of num1 , ptr1", ptr1, "Address of ptr1:", ptr2, "Address of ptr2:", ptr3)
	fmt.Println("Value of num1 , ptr1", *ptr1, "Value num1 throu ptr2:", **ptr2, "Value num1 throu ptr3:", ***ptr3)
	fmt.Println("num1 only", *(&num1))

	// var ptr4 *int // what is the value? nil
	// IncrP(ptr4)   // What will happen? Segmentation error , trying to dereference a pointer which is nil that means does not have any address

	ptr4 := new(int) // it allocates some memory based on the type passes as an argument
	// it returns that address to a pointer
	IncrP(ptr4)
	fmt.Println(*ptr4)
	fmt.Println(*sum(10, 20))
	fmt.Println(*sub(10, 20))
	fmt.Println(*mul(10, 20))

}

func Incr(num int) { // num is assigned to a new variable which is pass or call by value
	num++
}

func IncrP(num *int) { // num uses the address of the passing variable so it is pass or call by reference
	*num++
}

func sum(a, b int) *int {
	c := new(int)
	*c = a + b
	return c
}

func sub(a, b int) *int {
	c := a - b
	return &c
}

func mul(a, b int) *int {
	var ptr *int
	c := a * b
	ptr = &c
	return ptr
}

// what is a pointer? Pointer is a variable that holds the address
// How to create a pointer?
// type inference work for pointer as well!
// What is the default value? if no address is assigned , it is nil
