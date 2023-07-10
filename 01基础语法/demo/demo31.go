package main

import "fmt"

//简单闭包
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(20))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
