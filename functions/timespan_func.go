package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	longCalc()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalc took this amount of time:%s\n", delta)
}

func longCalc() {
	timespan := 3 * time.Second
	fmt.Printf("timespan:%s\n", timespan)
	time.Sleep(timespan)
}
