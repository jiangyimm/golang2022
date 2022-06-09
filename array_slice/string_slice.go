package main

import "fmt"

func main() {
	//1.subtsring function
	//使用 substr := str[start:end] 可以从字符串 str 获取到从索引 start 开始到 end-1 位置的子字符串。
	//同样的，str[start:] 则表示获取从 start 开始到 len(str)-1 位置的子字符串。
	//而 str[:end] 表示获取从 0 开始到 end-1 的子字符串。
	str1 := "123qwe江"
	sub1 := str1[2:]
	fmt.Printf("sub1: %v\n", sub1)

	//2.modify function
	//您必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，最后将字节数组转换回字符串格式。
	str2 := "123qwe江"
	c := []byte(str2)
	c[0] = '9'
	str3 := string(c)
	fmt.Printf("str3: %v\n", str3)
}
