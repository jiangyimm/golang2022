package main

import (
	"fmt"
	"time"
)

//注意：无缓冲通道的收发操作必须在不同的两个`goroutine`间进行，
//因为通道的数据在没有接收方处理时，数据发送方会持续阻塞，所以通道的接收必定在另外一个 goroutine 中进行。
//如果不按照该规则使用，则会引起经典的Golang错误`fatal error: all goroutines are asleep - deadlock!`：
func main() {
	ch := make(chan int)
	go write(ch)
	go read(ch)

	time.Sleep(time.Second * 3)
}

func write(ch chan int) {
	ch <- 100
	fmt.Printf("ch addr：%v\n", ch)
	ch <- 200
	fmt.Printf("ch addr：%v\n", ch)
	ch <- 300                      //没有读取，后续操作堵塞
	fmt.Printf("ch addr：%v\n", ch) //没有输出
}

func read(ch chan int) {
	fmt.Printf("取出的数据data1：%v\n", <-ch)
	fmt.Printf("取出的数据data2：%v\n", <-ch)
}
