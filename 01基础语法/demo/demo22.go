package main

import "fmt"

func main() {
	a := make([]int, 5) //cap不填，默认==size
	//这里b内部存储空间分配了10个，但实际上用了5个。容量不印象当前元素的格式
	b := make([]int, 5, 10)
	fmt.Println(a) //[0 0 0 0 0]
	fmt.Println(b) //[0 0 0 0 0]
	//使用内置函数len获取切片长度，cap获取切片容量
	//注意：判断切片是否为空，使用len(b)==0来判断，而不是b==nil
	fmt.Println(len(b), cap(b)) //5 10

	c := make([]int, 5) //make创建切片
	d := make([]string, 5, 10)
	c[0] = 1
	c[1] = 5
	d[0] = "北京"
	fmt.Println(c) //[1 5 0 0 0]
	fmt.Println(d) //[北京    ]
}
