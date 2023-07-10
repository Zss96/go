package main

import "fmt"

func main() {
	//常规写法
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	fmt.Println(score)

	//特殊写法
	if s := 95; s >= 90 {
		fmt.Println("A")
	} else if s > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//fmt.Println(s) //报错，s undefined

}
