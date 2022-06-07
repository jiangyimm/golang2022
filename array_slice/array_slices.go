package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] //not included index 5

	//load arr1
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	//print arr1
	printArr(arr1)

	//print slice
	printSlice(slice1)

	//re load arr1
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 10 + i
	}

	//print arr1
	printArr(arr1)

	//print slice
	printSlice(slice1)

	//grow slice
	slice1 = slice1[0:4]
	//print slice
	printSlice(slice1)

	slice1 = append(slice1, 111)
	//print slice
	printSlice(slice1)

	//给定切片 b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}，那么 b[1:4]、b[:2]、b[2:] 和 b[:] 分别是什么？
	slice2 := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("slice2[1:4] is %c\n", slice2[1:4])
	fmt.Printf("slice2[:2] is %c\n", slice2[:2])
	fmt.Printf("slice2[0:2] is %c\n", slice2[0:2])
	fmt.Printf("slice2[2:] is %c\n", slice2[2:])
	fmt.Printf("slice2[:] is %c\n", slice2[:])

	//因为字符串是纯粹不可变的字节数组，它们也可以被切分成切片
	str := "123hhhasd"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str at %d is %c \n", i, str[i])
	}

	// person := new(Person)
	// person.name = "123"
	// fmt.Printf("person.name is %s", person.name)
}

func printArr(arr [6]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr1[%d]:%d\n", i, arr[i])
	}
}

func printSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice1[%d]:%d\n", i, slice[i])
	}
	fmt.Printf("slice1 cap is %d\n", cap(slice))
}

type Person struct {
	name string
}
