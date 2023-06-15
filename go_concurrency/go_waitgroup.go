package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.weibo.com",
		"https://www.baidu.com",
		"https://www.163.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			http.Get(u)
			fmt.Printf("url: %v\n", u)
		}(url)
	}

	wg.Wait()
	fmt.Println("over")
}
