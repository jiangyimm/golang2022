package main

import (
	"fmt"
	"unicode"
)

var a = "G"

func main() {
	n()
	m()
	n()
	var isLetter = unicode.IsLetter('a')
	fmt.Printf("isLetter: %v\n", isLetter)
}

func n() {
	fmt.Printf("a: %v\n", a)
}

func m() {
	a = "O"
	fmt.Printf("a: %v\n", a)
}
