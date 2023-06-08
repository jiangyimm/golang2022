package main

import (
	"fmt"
	"time"
)

//channel只支持 for--range 的方式进行遍历：
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i <= 3; i++ {
			ch <- i
			time.Sleep(time.Second * 3)
		}
	}()

	for data := range ch {
		fmt.Println("data==", data)
		if data == 3 {
			break
		}
	}
}
