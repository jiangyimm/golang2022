package main

import "fmt"

var a string = "a"

func main() {
	f1()
}

func f1() {
	a := "O" //:= this new a
	fmt.Println(&a)
	f2()
}

func f2() {
	fmt.Println(&a)
}
