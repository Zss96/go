package main

import "fmt"

//结构体定义
type person struct {
	name string
	city string
	age  int8
}

func main() {
	//实例化
	var p1 person
	p1.name = "zss"
	p1.city = "广州"
	p1.age = 25

	fmt.Printf("%v\n", p1)  //{zss 广州 25}
	fmt.Printf("%#v\n", p1) //main.person{name:"zss", city:"广州", age:25}
	fmt.Println(p1.name)    //zss

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "zss"
	user.Age = 25
	fmt.Printf("%#v\n", user) //struct { Name string; Age int }{Name:"zss", Age:25}

	//使用new实例化,得到的是结构体的地址
	p2 := new(person)
	fmt.Printf("p2=%T\n", p2)  //p2=*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	p2.name = "zss"
	p2.age = 27
	p2.city = "广州"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"zss", city:"广州", age:27}

	//取结构体地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "zss"
	p3.age = 23
	p3.city = "广州"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"zss", city:"广州", age:23}
}
