package main

import "fmt"

func main() {
	//函数调用通过函数名()方法调用
	sayHello()      //无参调用
	a := Sum(1, 2)  //有参函数,
	b := Sum1(1, 2) //有参函数,
	c := Sum2(1, 2)
	d := Sum2(1, 2, 3)
	e := Sum4(1, 2)
	x, y := Point(1, 9)

	fmt.Println(a)    //3
	fmt.Println(b)    //3
	fmt.Println(c)    //3
	fmt.Println(d)    //6
	fmt.Println(e)    //3
	fmt.Println(x, y) //7 10
}
func sayHello() {
	fmt.Println("Hello world!")
}

//函数传参
func Sum(x int, y int) int {
	return x + y
}
func Sum1(x, y int) int {
	return x + y
}

//可变参数
func Sum2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 返回值
//存在返回变量名和返回类型，上面忽略了返回变量名
func Sum4(x, y int) (sum int) {
	return x + y
}

//多返回值

func Point(x, y int) (int, int) {
	return x + 6, y + 1
}
