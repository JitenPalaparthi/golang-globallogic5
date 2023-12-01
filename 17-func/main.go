package main

import "fmt"

func main() {
	a1 := AreaOfRect(12.34, 14.45)
	p1 := PerimeterOfRect(12.34, 12.45)
	fmt.Println("Area a1:", a1)
	fmt.Println("Perimeter p1:", p1)
	a2, p2 := Rect(12.34, 14.45)
	fmt.Println("Area a2:", a2)
	fmt.Println("Perimeter p12", p2)

	_, p3 := Rect(12.34, 14.45)
	fmt.Println("Perimeter p12", p3)

	a4, _ := Rect(12.34, 14.45)
	fmt.Println("Area a4:", a4)

	arr1 := [3]int{123, 342, 54}
	fmt.Println("Sum of arr1:", SumOfArr1(arr1))
	arr2 := [4]int{123, 342, 54, 43}
	fmt.Println("Sum of arr1:", SumOfArr2(arr2)) // cannot use SumOfArr1 as input param is [3]int

}

// func with two input parameters and one return value
func AreaOfRect(l float32, b float32) float32 {
	return l * b
}

func PerimeterOfRect(l, b float32) (p float32) {
	p = 2 * (l + b)
	//return 2 * (l + b)
	return p
}

func Rect(l, b float32) (a float32, p float32) {
	return l * b, 2 * (l + b)
}

func SumOfArr1(arr [3]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumOfArr2(arr [4]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return sum
}
