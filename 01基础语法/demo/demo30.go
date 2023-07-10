package main

import "fmt"

func main() {
	//将匿名函数保存在变量中
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) //通过变量调用匿名函数

	//自定义匿名函数
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

}
