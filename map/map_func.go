package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["key1"] = 1
	map1["key2"] = 2

	//判断是否存在
	_, isExist := map1["key1"]
	fmt.Printf("isExist: %v\n", isExist)

	//删除key
	delete(map1, "key1")
	_, isExist = map1["key1"]
	fmt.Printf("isExist: %v\n", isExist)
}
