package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 10)
	//load the slice
	for i := 0; i < cap(slice1); i++ {
		//扩容
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("slice1: %v\n", slice1)
	}

	//arr
	var arr = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice2 = arr[5:7]
	fmt.Printf("slice2: %v cap(slice2)：%d\n", slice2, cap(slice2))
	slice2 = slice2[0:4]
	fmt.Printf("slice2: %v cap(slice2)：%d\n", slice2, cap(slice2))

	//slice
	var s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice3 = s[5:7]
	fmt.Printf("slice3: %v cap(slice3)：%d\n", slice3, cap(slice3))
	slice3 = slice3[0:4]
	fmt.Printf("slice3: %v cap(slice3)：%d\n", slice3, cap(slice3))
}
