package main

import (
	"fmt"     // stadard package
	"reflect" // it is a standard package
	"strconv"
	_ "time" // _ blank identifer can be used here as well
)

func main() {

	var num1 int8 = 123
	//var num2 int = num1 // this is not allowed
	var num2 int = int(num1) // is it not (int)num1;
	fmt.Println(num2)

	num3 := 12312312             // there are 64 bits for int on 64 bit arch
	var num4 uint8 = uint8(num3) // there are 8 bits for int8
	fmt.Println(num4)            // 248
	float1 := 12312.12321
	var num5 int = int(float1)
	fmt.Println(num5)

	// strings

	var num6 int = 19000

	str1 := string(num6)

	fmt.Println(str1, reflect.TypeOf(str1))
	//"19000"

	str2 := fmt.Sprint(num6)
	fmt.Println(str2, reflect.TypeOf(str2))

	str3 := strconv.Itoa(num6) // this converts int as it is to string
	fmt.Println(str3, reflect.TypeOf(str3))
	str4 := "A" //array of chars
	var num7 int = int(str4[0])
	fmt.Println(num7) // what is the value of num7?

	str5 := "19000"
	var num8 int

	num8, _ = strconv.Atoi(str5) // Funcs or methods in go can return more than one value like touple
	// _ blank identifier. If a func or method is returning more than one return type , if you dont want all the return values
	// you can place a _ blank identifer
	// for any identifer you can use blank identifer
	fmt.Println(num8)
}

// A type can never be converted to another type automatically/implicitly
// Go does not support/ allow implicit type cast
// Converting one type another type is only thru type casting
