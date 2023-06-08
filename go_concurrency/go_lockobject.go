package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	a := &Account{
		0,
		&sync.Mutex{},
	}

	for i := 0; i < 100; i++ {
		go func(num int) {
			a.add(num)
		}(10)
	}

	time.Sleep(time.Second * 2)
	a.query()
}

type Account struct {
	money int
	lock  *sync.Mutex
}

func (a *Account) query() {
	fmt.Println("当前金额为：", a.money)
}

func (a *Account) add(num int) {
	a.lock.Lock()
	a.money += num
	a.lock.Unlock()
}
