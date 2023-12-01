package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	day := 4
	switch day { // this is for each case it is kind of day==1 or day==2
	case 1:
		fmt.Println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("no day")
	}

	num := -55

	switch { // empty switch. a switch without any variable; if it is empty switch can write conditons in the case
	case num < 0:
		println(num, "is less then zero")
	case num >= 0 && num <= 50:
		println(num, "is between 0 and 50")
	case num > 50 && num <= 100:
		println(num, "is between 51 and 100")
	case num > 100 && num <= 200:
		println(num, "is between 101 and 200")
	default:
		println(num, "is  above 200 ")
	}

	age, gender := 25, 'm'
	switch {
	case age >= 18 && (gender == 'F' || gender == 'f'):
		println("she is eligible for marriage")
	case age >= 21 && (gender == 'M' || gender == 'm'):
		println("he is eligible for marriage")
	default:
		println("not eligible for marriage; may have invalid gender")
	}

	fmt.Println("Another switch ")
	char := 'E'
	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		fmt.Println(string(char), " is vovel")
	default:
		fmt.Println(string(char), " is consonent or special char")
	}

}

// no need to give break to each case. break is automatically invoked
// 1-hello
// 	src/main.go
// 	pkg/
// 	bin/
// export GOPATH=$GOPATH:"~/workspace/examples"
// godeps

// go install github.com/abcd/gopals@latest --> This is going to instlal in your gobin directory
// GOOS, GOHOSTOS
// GOARCH, GOHOSTARCH
