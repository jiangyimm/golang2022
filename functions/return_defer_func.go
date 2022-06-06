package main

import "fmt"

func main() {
	ret := test()
	fmt.Printf("ret is %d", ret)
}
func test() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
