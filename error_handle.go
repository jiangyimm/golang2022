package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("ww")
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(1)
	}
	fmt.Println(i)

	if _, err1 := OpenFile("time_fun.go1"); err1 != nil {
		fmt.Println(err1.Error())
	}
}

func OpenFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return file, err
	}
	return file, nil
}
