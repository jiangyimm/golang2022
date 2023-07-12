package main

import "fmt"

func main() {
	//complex
	var t complex128
	t = 2.1 + 3.14i
	fmt.Printf("real: %v, imag: %v\n", real(t), imag(t))
	t1 := 2.1 + 3.14i
	fmt.Printf("real: %v, imag: %v\n", real(t1), imag(t1))
	t2 := complex(2.1, 3.14)
	fmt.Printf("real: %v, imag: %v\n", real(t2), imag(t2))
}
