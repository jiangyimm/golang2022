package main

import (
	"fmt"
	"os"
)

func main() {
	// //创建文件 存在则清空文件内容
	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// 	return
	// }
	// fmt.Printf("f: %v\n", f)
	// f.Close()

	//打开文件
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	//写入文件
	b := []byte("12345")
	fmt.Printf("len(b): %v\n", len(b))
	b = []byte("哈哈哈哈哈")
	fmt.Printf("len(b): %v\n", len(b))
	n, _ := f.Write(b)
	fmt.Printf("n: %v\n", n)
	n, _ = f.WriteString("123123123")
	fmt.Printf("n: %v\n", n)
	f.Close()
}
