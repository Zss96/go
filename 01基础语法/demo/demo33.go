package main

import "fmt"

func main() {
	a := 10
	b := &a                            // 取变量a的地址，将指针保存到b中
	c := *b                            // 指针取值（根据指针去内存取值）
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc0000a6058
	fmt.Printf("b:%p type:%T\n", b, b) //b:0xc0000a6058 type:*int
	fmt.Println(&b)                    //0xc0000ca018
	fmt.Printf("c:%d type:%T\n", c, c) //c:10 type:int
}
