package main

import "fmt"

func foo() (int, string) {
	return 25, "zss"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
