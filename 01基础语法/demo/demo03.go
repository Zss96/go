package main

import "fmt"

func main() {
	//十进制
	var a int = 10
	fmt.Printf("a的值十进制表示%d \n", a)  // 10
	fmt.Printf("a的值二进制表示%b \n", a)  // 1010  占位符%b表示二进制
	fmt.Printf("a的值八进制表示%o \n", a)  // 12
	fmt.Printf("a的值十六进制表示%x \n", a) //a
	fmt.Printf("a的值十六进制表示%X \n", a) //A

	//八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

}
