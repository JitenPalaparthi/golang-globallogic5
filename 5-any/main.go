package main

import (
	"fmt"
	"reflect"
)

type integer = int // alias

// type any = interface{}

func main() {

	//var iface1 interface{} // emoty interface
	var iface1 any

	var ok = true
	iface1 = ok
	fmt.Println(iface1, reflect.TypeOf(iface1))
	//var ok1 bool = bool(iface1) // does not work
	// when the value inside any/interface{} has to be converted to a respective type then
	// type casting does not work. Here have to use a type assertion
	var ok1 bool = iface1.(bool) // V.(T)
	// when the right hand side is an interface{} then do type assertion;Dont do type casting
	fmt.Println(ok1)

	str1 := "Hello World"
	iface1 = str1
	//var ok3 bool = iface1.(bool) // does not work
	//fmt.Println(ok3)

	var str2 string = iface1.(string)
	fmt.Println(str2)

	fmt.Println(iface1, reflect.TypeOf(iface1))
	num1 := 12321
	iface1 = num1
	fmt.Println(iface1, reflect.TypeOf(iface1))

	float1 := 12313.123123123123
	iface1 = float1
	fmt.Println(iface1, reflect.TypeOf(iface1))

	var num2 integer = 1231212
	iface1 = num2
	fmt.Println(iface1, reflect.TypeOf(iface1))
	iface1 = nil
	fmt.Println(iface1, reflect.TypeOf(iface1))

	var iface2 any
	fmt.Println(iface2, reflect.TypeOf(iface2))

}

// empty interface / interface{} / any can store any kind of data
// When declared and no data is assigned then interface{} contains nil
// there is no null in go; there is nil
