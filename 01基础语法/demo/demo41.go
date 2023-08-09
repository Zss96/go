package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

//值类型接收者
func (p person) Study() {
	fmt.Printf("%s要学好GO\n", p.name)
}
func (p person) SetAge1(age int8) {
	p.age = age
}

//指针类型接收者
func (p *person) SetAge(age int8) {
	p.age = age
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
	fmt.Printf("%#v\n", p1)
	p1.Study()
	p1.SetAge1(100)
	fmt.Println(p1.age) //20
	p1.SetAge(27)
	fmt.Println(p1.age) //27
}
