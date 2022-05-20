package main

import "fmt"

func main() {
	x := max(1, 3, 5, 0)
	fmt.Println(x)
	arr := []int{1, 3, 5, 0}
	x = max1(arr)
	fmt.Println(x)
}

func max(num ...int) int {
	return max1(num)
}

func max1(num []int) int {
	if len(num) == 0 {
		return 0
	}
	max := num[0]
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
