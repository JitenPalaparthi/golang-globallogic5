package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num1 myint = 100

	var num2 int = 200

	fmt.Println("Value of num1:", num1, "To string:", num1.ToString(), "Type of:", reflect.TypeOf(num1.ToString()))

	str2 := myint(num2).ToString()
	fmt.Println(str2, "Type:", reflect.TypeOf(str2))

	var mm1 mymap

	mm1 = make(mymap)
	mm1["bangalore-1"] = 560086
	mm1["bangalore-2"] = 560096
	mm1["guntur-1"] = 522001
	fmt.Println(mm1.GetKeys())
	fmt.Println(mm1.GetValues())

	map1 := make(map[string]any)
	map1["bangalore-1"] = 560086
	map1["bangalore-2"] = 560096
	map1["guntur-1"] = 522001

	keys := mymap(map1).GetKeys()
	values := mymap(map1).GetValues()
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)

	var v1 Empty
	v1.Greet()
	v1.GreetP()

	var v2 struct{}

	Empty(v2).Greet()

}

type myint int

type mymap map[string]any

func (mi myint) ToString() string {
	return fmt.Sprint(mi)
}

func (mm mymap) GetKeys() []string {
	keys := make([]string, len(mm))
	index := 0
	for k := range mm {
		keys[index] = k
		index++
	}
	return keys
}

func (mm mymap) GetValues() []any {
	values := make([]any, len(mm))
	index := 0
	for _, v := range mm {
		values[index] = v
		index++
	}
	return values
}

type Empty struct{}

func (e Empty) Greet() {
	fmt.Println("hello world!")
}

func (e *Empty) GreetP() {
	fmt.Println("hello world!")
}
