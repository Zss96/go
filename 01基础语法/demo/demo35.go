package main

import "fmt"

func main() {
	//new
	a := new(int)
	b := new(bool)
	c := new(string)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Printf("%T\n", c) // *string
	fmt.Println(a, *a)    // 0xc000018088 0
	fmt.Println(b, *b)    //0xc0000180a0 false
	fmt.Println(c, *c)    //0xc000054270
	var d *int
	d = new(int)
	*d = 10
	fmt.Println(d, *d) //0xc000018100  10

	//make
	var e map[string]int
	e = make(map[string]int, 10)
	e["张三"] = 100
	fmt.Println(e) //map[张三:100]
}
