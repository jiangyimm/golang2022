package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	result := Sum(2, 3)
	fmt.Println(result)

	var goos string = runtime.GOOS
	var path string = os.Getenv("PATH")
	fmt.Printf("goos: %v\n", goos)
	fmt.Printf("path: %v\n", path)

	fmt.Println(&result)
}

func Sum(a int, b int) int {
	return a + b
}

//const type
const Pi int = 3
const Ln2 = 0.693147180559945309417232121458 / 176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

//const like enum
const (
	UnKnow      = 1
	Female int8 = 2
	Male        = 3
)
const (
	a = iota //default 0
	b
	c
)

type Color int

const (
	RED Color = iota
	YELLOW
)
