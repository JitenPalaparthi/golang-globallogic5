package main

import "fmt"

func main() {

	var iface1 any

	iface1 = "Hello World" // type ? string

	iface1 = 101.12 // float64

	iface1 = true // bool

	iface1 = 'A'

	switch iface1.(type) { // special notation for type switch

	case int:
		fmt.Println("integer value:", iface1)
	case string:
		fmt.Println("string value:", iface1)

	case bool:
		fmt.Println("bool value:", iface1)

	case float64:
		fmt.Println("float64 value:", iface1)

	default:
		fmt.Println("unidentified type")
	}

	println("another example")

	var iface2 any

	var num1 uint8 = 14
	iface2 = num1
	iface2 = true

	switch v := iface2.(type) {
	case uint8:
		fmt.Println("uint8", v*v)
	case int:
		fmt.Println("int", v*v)
	case float64:
		fmt.Println("float64", v*v)
	default:
		fmt.Println("cannot perform the square operation")
	}

	char := '䨸' // 19000
	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		fmt.Println(string(char), "is vovel")
	default:
		fmt.Println(string(char), "is consonent or special char")
	}

}

// 1- expression switch
// 2- type switch --> specially used for any
