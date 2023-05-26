package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	conn.Write([]byte("hello"))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("client read err:", err)
		return
	}
	fmt.Println("client receive: ", string(buf[:n]))
}
