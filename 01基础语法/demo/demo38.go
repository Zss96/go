package main

import "fmt"

//结构体定义
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p person
	//没有初始化结构体的值
	fmt.Printf("p=%#v\n", p) //p=main.person{name:"", city:"", age:0}

	//键值对
	//普通类型初始化
	p1 := person{
		name: "zss",
		city: "广州",
		age:  20,
	}
	//结构体指针类型初始化
	p2 := &person{
		name: "zss",
		city: "广州",
		age:  20,
	}
	//某些字段为空时
	p3 := person{
		name: "zxx",
	}
	p4 := person{
		name: "zxx",
	}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"zss", city:"广州", age:20}
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"zss", city:"广州", age:20}
	fmt.Printf("p3=%#v\n", p3) //p3=main.person{name:"zxx", city:"", age:0}
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"zxx", city:"", age:0}

	//在初始化时候，也可以省略键，不过这样，所有字段都必须要初始化，并且字段顺序要结构体一致。
	p5 := person{
		"zjj",
		"北京",
		18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"zjj", city:"北京", age:18}
}
