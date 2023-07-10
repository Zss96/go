package main

import "fmt"

func main() {
	//常规for循环
	for i := 0; i < 10; i++ {
		fmt.Println("常规for循环", i)
	}
	//for循环初始语句可以省略

	i := 0
	for ; i < 10; i++ { //注意：初始语句省略，但是后面的分号(;)一定要写
		fmt.Println("初始语句省略", i)
	}

	//结束语句省略
	for j := 0; j < 10; { //注意：结束语句省略也要写分号(;)一
		fmt.Println("结束语句省略", j)
		j++
	}

	//初始语句和结束语句都省略
	for i < 10 {
		fmt.Println("初始语句和结束语句都省略", i)
		i++
	}
}
