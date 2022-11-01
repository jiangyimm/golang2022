package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	//var ms *struct1=new(struct1)
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "chris"

	fmt.Printf("ms.i1: %d\n", ms.i1)
	fmt.Printf("ms.f1: %f\n", ms.f1)
	fmt.Printf("ms.str: %s\n", ms.str)
	fmt.Printf("ms: %v\n", ms)

	//结构体
	ms1 := struct1{10, 15.5, "qwr"}
	fmt.Printf("ms1: %v\n", ms1)
	//指针
	ms2 := &struct1{10, 15.5, "jjjy"}
	fmt.Printf("ms2: %v\n", ms2)
}
