package main

import (
	"fmt"
	"math"
)

func main() {

	//print slice
	slice1 := make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i + 1
	}

	for ix, val := range slice1 {
		fmt.Printf("slice1 at %d is %d\n", ix, val)
	}

	//fmt.Printf("slice1: %v\n", slice1)

	//modify arr
	items := [...]int{10, 20, 30, 40, 50}
	//invalid
	for _, item := range items {
		item *= 2
	}
	//valid
	for ix := range items {
		items[ix] *= 2
	}
	for ix, item := range items {
		fmt.Printf("items at %d is %d\n", ix, item)
	}

	min := minSlice(slice1)
	fmt.Printf("min: %v\n", min)
	max := maxSlice(slice1)
	fmt.Printf("max: %v\n", max)

	arrF := [...]float32{1, 2, 3.11, 4.222, 0.23}
	sum := sumSlice(arrF[:])
	fmt.Printf("sum: %v\n", sum)

	sum1, avg1 := sumAndAvg(arrF[:])
	fmt.Printf("sum1: %v\n", sum1)
	fmt.Printf("avg1: %v\n", avg1)
}

func minSlice(slice1 []int) (min int) {
	min = math.MaxInt32
	for _, v := range slice1 {
		if v < min {
			min = v
		}
	}
	return
}

func maxSlice(slice1 []int) (max int) {
	for _, v := range slice1 {
		if v > max {
			max = v
		}
	}
	return
}

func sumSlice(arrF []float32) (sum float32) {
	for _, v := range arrF {
		sum += v
	}
	return
}

func sumAndAvg(arrF []float32) (int, float32) {
	var sum int = int(sumSlice(arrF))
	avg := float32(sum / len(arrF))
	return sum, avg
}
