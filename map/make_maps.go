package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssiged map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssiged = mapLit
	fmt.Printf("mapLit: %v\n", mapLit)
	fmt.Printf("mapAssiged: %v\n", mapAssiged)
	fmt.Printf("mapCreated: %v\n", mapCreated)

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssiged["two"] = 3 //mapAssigned 也是 mapLit 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。
	fmt.Printf("mapLit: %v\n", mapLit)
	fmt.Printf("mapAssiged: %v\n", mapAssiged)
	fmt.Printf("mapCreated: %v\n", mapCreated)
}
