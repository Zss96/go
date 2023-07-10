package main

import "fmt"

func main() {
	//声明和初始化切片
	var a []string
	var b []string
	fmt.Println(a)        //[]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //true
	//fmt.Println(a == b) //切片是引用类型，不支持直接比较，只能和nil比

	//初始化切片
	a = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(a) //[北京 上海 广州 深圳]

	//切片存在自己的长度len()和容量cap()
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(c), cap(c)) // 5 ,5
}
