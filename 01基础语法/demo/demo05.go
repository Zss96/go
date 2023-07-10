package main

import "fmt"

func main() {
	var c1 complex64
	c1 = 1 + 2i
	c2 := complex(2, 3)

	fmt.Println(c1)
	fmt.Println(c2)
}
