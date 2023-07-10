package main

import "fmt"

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}

func main() {
	var c calc
	c = add
	a := c(1, 2)
	fmt.Println(a) //3
}
