package main

import "fmt"

func main() {

	fmt.Println("Hello 世界!")
	var num1 int = 100
	var num2 uint8 = 255
	var float1 float32 = 123.2342
	var str1 string = "Hello GlobalLogic"
	var ok1 bool = true
	var char1 rune = '世'
	fmt.Println(num1, num2, float1, str1, ok1, char1)

	// type inference : go compiler inferes a value(default) based on the type
	var (
		num3   int32
		float2 float64
		ok2    bool
		str2   string
	)
	fmt.Println(num3, float2, ok2, str2)

	// automatically the compiler assign database based on the value given
	var num4 = 300 // int

	var ok3 = true // bool

	var str3 = "Hello World" //string

	fmt.Println(num4, ok3, str3)

	// short hand declaration
	num5 := 12123123123

	float3 := 123.12321

	ok4 := false

	str5 := "Hello World"
	fmt.Println(num5, float3, ok4, str5)

	age := 39      // 255
	marks := 89.23 //
	fmt.Println(age, marks)

	const pi float32 = 3.14 //declare a constant with proper datatype

	const (
		max  = 10 // short hand declartion
		min  = 1
		days = 7

		ok = true
	)

}

// notes:
// numbers
// int, uint, int8,int16,int32,int64,uint8,uint16,uint32,uint64
// float32, float64
// if you compile application for 32 bit arch then int,uint is 4 bytes if on 64bit arch int,uint is 8 bytes
// boolean
// bool
// string
// string
// special
// rune, byte : rune is nothing but an alias for int32. But in practice we use rune for charcters
// technically you can also use int32.
// byte: alias for uint8. technically you can use uint8 intead of byte

// utf vs ascii ?
// ascii are 256 chars.
