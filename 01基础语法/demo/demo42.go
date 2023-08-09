package main

import "fmt"

type MyInt int

func (m MyInt) sayHello() {
	fmt.Println("Hello World")
}

func main() {
	var m MyInt
	m.sayHello()
	m = 10
	fmt.Printf("%#v,%T\n", m, m) //10,main.MyInt
}
