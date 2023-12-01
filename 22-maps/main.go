package main

import "fmt"

func main() {

	var map1 map[string]int // declaration key is string and value is any
	if map1 == nil {
		fmt.Println("yes nil map1")
	}
	map1 = make(map[string]int) // instantiation or allocation
	// can check whether a map is nil or not

	map1["bangalore-1"] = 560086
	map1["bangalore-2"] = 560096
	map1["bangalore-3"] = 560034
	map1["guntur-1"] = 522001
	map1["abcd"] = 0

	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	val := map1["guntur-1"]
	fmt.Println(val)

	val, ok := map1["guntur-3"]
	if !ok {
		fmt.Println("no key exists")
	} else {
		fmt.Println(val)
	}

	_, ok = map1["guntur-4"]
	if !ok {
		map1["guntur-4"] = 522004
		println("insert a new element in the map")
	} else {
		map1["guntur-4"] = 522004
		println("update  existing element in the map")
	}

	delete(map1, "guntur-3")
	fmt.Println(map1)

}

// map is a keyword
// map contains keys and values
// nil can be checked on map
// map is not thread safe
// map is not ordered
// range loop has to be used to fetch all elements from the map
// range loop returns key and value for the map
// if the key is not unique , the value gets updated
// maps are very good for time complexity.Usually they are O(1)
// can check whether a map contains specific key or not as map yields two values
// if the key does not exist, map returns default value for the value
// to delete a key from the map use delete builtin function
// What is hashkey, hashfunc
