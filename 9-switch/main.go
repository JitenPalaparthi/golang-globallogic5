package main

func main() {
	num := 42
	// println("reminder", num%2)
	// fmt.Println("coif", float32(num/2))
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough // if you avoid break in other lang
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println("not divisible by 8,4 or 2")
	}
}

// where or when ever ypu avoid break in other programming switch you add fallthrough in go
