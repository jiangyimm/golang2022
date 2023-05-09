package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Student struct {
	Name    string
	Age     uint8
	Address string
	Hobby   []string
}

func main() {
	//1序列化
	// stu := Student{
	// 	Name:    "张三",
	// 	Age:     12,
	// 	Address: "hefei",
	// 	Hobby:   []string{"哈哈", "14A80249.png"},
	// }
	// b, _ := jsoniter.Marshal(stu)
	// fmt.Println(string(b))

	// str, _ := jsoniter.MarshalToString(stu)
	// fmt.Println(str)

	// b, _ = jsoniter.MarshalIndent(stu, "", " ")
	// fmt.Println(string(b))

	//2反序列化
	//var stu2 Student
	stu2 := new(Student)
	jsonBlob := []byte(`{"Name":"张三","Age":12,"Address":"hefei","Hobby":["哈哈","14A80249.png"]}`)
	jsonStr := `{"Name":"张三","Age":12,"Address":"hefei","Hobby":["哈哈","14A80249.png"]}`
	jsoniter.Unmarshal(jsonBlob, stu2)
	fmt.Printf("%+v\n", *stu2)
	jsoniter.UnmarshalFromString(jsonStr, stu2)
	fmt.Printf("%+v\n", *stu2)

}
