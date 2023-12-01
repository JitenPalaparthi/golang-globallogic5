package main

import "fmt"

func main() {

	var age uint8 = 16
	var gender rune = 'F'

	if age >= 18 {
		fmt.Println("eligible for vote")
	} else {
		fmt.Println("not eligible for vote")
	}

	//true && (true) = true
	if age >= 18 && (gender == 'F' || gender == 'f') {
		fmt.Println("She is eligible for marriage")
	} else if age >= 21 && (gender == 'M' || gender == 'm') {
		fmt.Println("He is eligible for marriage")

		//true && true
		// true
		// false && true -> false
		// } else if (gender != 'M' && gender != 'm') || (gender != 'F' && gender != 'f') {
		// 	println("invalid gender-1")
		// 	// } else if gender != 'F' || gender != 'f' {
		// 	// 	println("invalid gender-2")
		// 	//
	} else {
		println("not eligile for marriage")
	}

	if age1, gender1 := 14, 'F'; age1 >= 18 && gender1 == 'F' {
		fmt.Println("Eligible for marriage")
	} else if age1 >= 21 && gender1 == 'M' {
		println("He is eligible for marriage")
	} else {
		println("Not eligible for marriage. age is", age1, "and gender is", string(gender1))
	}
}

// if true {
// 	fmt.Println("It is  true")
// } else {
// 	fmt.Println("It is false")
// }

// println(true && true) // True
// println(true && false)
// println(false && true)
// println(false && false)

// println(true || true)
// println(true || false)
// println(false || true)
// println(false || false)
