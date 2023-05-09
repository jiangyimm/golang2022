package main

import "fmt"

func main() {
	bool1 := true
	if bool1 {
		fmt.Println("the value is true")
	} else {
		fmt.Println("the value is false")
	}

	abs := Abs(-999)
	fmt.Println("abs is", abs)

	isGreater := IsGreater(1, 2)
	fmt.Println("1 > 2 ?", isGreater)

	if val := GetVal(10); val > 1 {
		fmt.Println("10>1")
	}
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func IsGreater(x, y int) bool {
	return x > y
}

func GetVal(x int) int {
	return x
}
