package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name,
		city,
		age,
	}
}

func main() {
	p1 := newPerson("zss", "广州", 20)
	fmt.Printf("%#v\n", p1) //&main.person{name:"zss", city:"广州", age:20}
}
