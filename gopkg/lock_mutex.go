package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	info := new(Info)
	info.name = "123"

	for i := 0; i < 10; i++ {
		go Update(info, strconv.Itoa(i))
	}

	fmt.Printf("info.name: %v\n", info.name)

}

func Update(info *Info, name string) {
	info.mu.Lock()
	info.name = name
	info.mu.Unlock()
}

type Info struct {
	mu   sync.Mutex
	name string
}

type InfoVO struct {
	mu   sync.Mutex
	name string
	age  int16
}

type InfoDTO struct {
	mu   sync.Mutex
	name string
	age  string
}
