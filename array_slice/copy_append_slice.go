package main

import "fmt"

func main() {

	//copy
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	fmt.Printf("slFrom: %v\n", slFrom)
	fmt.Printf("slTo: %v\n", slTo)
	fmt.Printf("copy...\n")
	copyNum := copy(slTo, slFrom)
	fmt.Printf("slTo: %v\n", slTo)
	fmt.Printf("copyNum: %v\n", copyNum)

	//append
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Printf("sl3: %v , len(sl3)：%d, cap(sl3)：%d", sl3, len(sl3), cap(sl3))
}
