package main

import (
	"fmt"
	"unsafe"
)

type test struct {
	a    int8
	b    int8
	c, d int8
}

func main() {
	var n struct{}
	//unsafe.Sizeof
	fmt.Println(unsafe.Sizeof(n)) //0

	p := test{
		1, 2, 3, 4,
	}
	fmt.Printf("p.a %p\n", &p.a) //p.a 0xc0000180a4
	fmt.Printf("p.b %p\n", &p.b) //p.b 0xc0000180a5
	fmt.Printf("p.c %p\n", &p.c) //p.c 0xc0000180a6
	fmt.Printf("p.d %p\n", &p.d) //p.d 0xc0000180a7

}
