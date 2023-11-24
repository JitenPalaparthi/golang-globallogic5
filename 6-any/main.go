package main

import "fmt"

func main() {

	var num1 uint8 = 255
	var num2 int8 = 127
	var num3 int = 12312312
	var num4 int32 = 12312312
	var num5 int64 = 1231232323231
	var float1 float32 = 1231.21
	var float2 float64 = 12311.12312312
	var iface1 any = 12312.12312 // type ?
	var iface2 any = 123         // type ?
	var iface3 any = float1      // type ?
	var sum float64 = float64(num1) + float64(num2) + float64(num3) + float64(num4) + float64(num5) + float64(float1) + float2 + iface1.(float64) + float64(iface2.(int)) + float64(iface3.(float32))
	fmt.Println(sum)
	fmt.Printf("Value: %.4f Type: %T", sum, sum) // formater

}
