package main

import (
	"fmt"
)

func main() {
	//对值进行判断
	num := 2
	switch num {
	case 1:
		fmt.Println("数字1")
	case 2:
		fmt.Println("数字2")
	case 3:
		fmt.Println("数字3")
	default:
		fmt.Println("无效数字")
	}

	//一个case可以存在多个值，多个case值使用逗号(,)隔开
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//case还可以使用表达式，但是这时候，switch语句不需要再跟判断变量
	score := 80
	switch {
	case score > 85:
		fmt.Println("A")
	case score > 75:
		fmt.Println("B")
	case score > 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

}
