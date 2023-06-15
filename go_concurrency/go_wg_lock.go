package main

import (
	"fmt"
	"sync"
)

//等待组与互斥锁（同步锁）配合解决钱数问题示例
func main() {
	var mt sync.Mutex
	var wg sync.WaitGroup
	var money = 10000

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			mt.Lock()
			fmt.Printf("协程 %d 抢到锁\n", index)
			for j := 0; j < 100; j++ {
				money += 10
			}
			fmt.Printf("协程 %d 准备解锁\n", index)
			mt.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("最终的 money = ", money)
}
