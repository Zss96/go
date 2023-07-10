package main

import "fmt"

type callBack func(int, int) int

func add(x, y int) int {
	return x + y
}

//函数作为参数
func calc(x, y int, callback callBack) int {
	return callback(x, y)
}

//函数作为返回值
func increment(a int) callBack {
	return func(x, y int) int {
		return a + x + y
	}
}

func main() {
	//函数作为参数
	a := calc(10, 20, add)
	fmt.Println(a) //30

	//函数作为返回值
	b := increment(3)
	d := b(4, 5)
	fmt.Println(d) //12

}
