package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	//1-strcut as value type
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Printf("pers1: %v\n", pers1)

	//2-strcut as a pointer
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" //注意也可以通过解指针的方式来设置值：
	upPerson(pers2)
	fmt.Printf("pers2: %v\n", pers2)

	//3-struct as a literal
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Printf("pers3: %v\n", pers3)
}
