package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// 全部打印11：因为开启协程也会耗时，协程没有准备好，循环已经走完
	// ps：开启task打印内容，但是打印的i用的是主线程的，i一直在递增，
	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)

	fmt.Println(`------------------`)

	// 打印无规律数字
	// ps：开启task打印内容，有入参x，所以1~10都传入了，但是由于task执行没有顺序，所以输出不是按顺序
	for i := 1; i <= 10; i++ {
		go func(x int) {
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second)

	fmt.Println(`------------------`)

	//ps：`runtime.Gosched()`：用于让出CPU时间片，即让出当前协程的执行权限，调度器可以去安排其他等待的任务运行。
	for i := 1; i <= 10; i++ {
		go func(x int) {
			if x == 5 {
				runtime.Gosched()
			}
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second)

	fmt.Println(`------------------`)

	//ps：`runtime.Goexit()`：用于立即终止当前协程运行，调度器会确保所有已注册defer延迟调用被执行。
	for i := 1; i <= 5; i++ {
		defer fmt.Println("defer ", i)
		go func(x int) {
			if x == 3 {
				runtime.Goexit()
			}
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second)
}
