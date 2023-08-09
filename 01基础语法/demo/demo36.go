package main

import "fmt"

//类型定义
type MyInt int

//类型别名
type AliasInt = int

func main() {
	var a MyInt
	var b AliasInt

	fmt.Printf("a:%T\n", a) //a:main.MyInt
	fmt.Printf("b:%T\n", b) //b:int

}
