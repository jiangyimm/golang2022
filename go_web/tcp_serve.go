package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			return
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	if conn == nil {
		panic("conn is nil")
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("read EOF")
			break
		}

		if err != nil {
			fmt.Println("conn read err:", err)
			break
		}

		fmt.Println("serve receive msg:", string(buf[:n]))
		conn.Write([]byte("word"))
	}
}
