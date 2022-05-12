package main

import "fmt"

func main() {
	ForRange()
	//ForCharactor()
	//for {
	//	fmt.Println("123")
	//}
}

func ForCharactor() {
	for i := 1; i <= 25; i++ {
		for y := 1; y <= i; y++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}

func ForRange() {
	str := "江毅jiangyi"
	for pos, char := range str {
		fmt.Printf("pos:%d char:%c \n", pos, char)
	}
}
