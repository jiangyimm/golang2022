package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(time.Second * 3)
}

func write(ch chan int) {
	ch <- 100
	fmt.Printf("ch: %v\n", ch)
	ch <- 200
	fmt.Printf("ch: %v\n", ch)
	ch <- 300                  // 写入第三个，造成阻塞
	fmt.Printf("ch: %v\n", ch) // 没有输出
}
