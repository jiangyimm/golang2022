package main

import (
	"fmt"
	"strconv"
)

func main() {
	map1 := make(map[string]float32)
	for i := 0; i < 10; i++ {
		map1[strconv.Itoa(i)] = float32(i)
	}
	fmt.Printf("map1: %v\n", map1)

	for k, v := range map1 {
		fmt.Printf("key is:%v - value is:%f\n", k, v)
	}
}
