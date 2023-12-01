package main

func main() {
	Greet()
	Greet()
	GreetM("Hello World!")
	msg1 := "Hello Gofers!"
	GreetM(msg1) // deep copy of a variable
	num1 := 100
	Incr(num1) // call or pass by value, deep copy of a variable
	println(num1)
}

func Greet() {
	msg := "Hello GlobalLogic folks!" // msg is a string variable
	println(msg)                      // what is the scope of msg? with in the func
}

func GreetM(msg string) {
	println(msg) // what is the scope of msg? with in the func
}

func Incr(num int) { // a new variable is created called num in the Incr func stackframe
	num++
	println("Inside func:", num)
}
