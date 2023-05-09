package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var hasPrefix bool = strings.HasPrefix("123", "1")
	var hasSuffix bool = strings.HasSuffix("456", "123")
	var s = strings.TrimLeft(" 123", " ")

	var files = strings.Fields("123 23")
	var splitss = strings.Split("123 23", "2")

	var joins = strings.Join(files, ",")

	var itoas = strconv.Itoa(123)
	var atoi, _ = strconv.Atoi("1233")

	fmt.Printf("hasPrefix: %v\n", hasPrefix)
	fmt.Printf("hasSuffix: %v\n", hasSuffix)

	fmt.Printf("s: %v\n", s)

	fmt.Printf("files: %v\n", files)
	fmt.Printf("splitss: %v\n", splitss)

	fmt.Printf("joins: %v\n", joins)

	fmt.Printf("itoas: %v\n", itoas)

	fmt.Printf("atoi: %v\n", atoi)
}
