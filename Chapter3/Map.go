package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1

	fmt.Println("aMap : ", aMap)
	aMap = nil

	if aMap == nil {
		fmt.Println("nil map!")
		aMap = make(map[string]int)
	}

	aMap["123"] = 1
	aMap["key"] = 456

	fmt.Println(len(aMap))

	v, ok := aMap["123"]
	if ok {
		fmt.Println(v)
	}

	for key, v := range aMap {
		fmt.Println("key: ", key, "value: ", v)
	}

}
