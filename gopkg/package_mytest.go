package main

import (
	"fmt"

	"./pack1"
	//. "./pack1"
	//_ "./pack1"
)

func main() {
	str := pack1.ReturnStr()
	fmt.Printf("str: %v\n", str)
}
