package main

import (
	"fmt"
	"time"
)

func main() {
	var nowTime = time.Now()
	fmt.Printf("nowTime: %v\n", nowTime)
	var nowFormat = nowTime.Format(time.RFC3339)
	fmt.Printf("nowFormat: %v\n", nowFormat)
}
