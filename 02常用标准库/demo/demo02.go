package main

import "fmt"

func main() {

	var name string
	fmt.Print("请输入你的名字:")
	//注意，这里要传地址，因为你要修改其值
	fmt.Scanln(&name) //从标准流输入读取一行文本并存储再name变量中
	fmt.Printf("Hello, %s\n", name)

	fmt.Printf("%s今年多少岁\n", name)
	var age int
	fmt.Scan(&age)
	fmt.Printf("%s今年%d岁\n", name, age)

	fmt.Println("请输入你最好的朋友的名字和年龄")
	/**
	*Scanln 函数会留下一个换行符在输入缓冲中，导致 Scanf 函数在读取朋友的名字时会直接读取这个换行符，而无法等待用户的输入
	 */
	fmt.Scanln()                    // 读取并消费掉之前剩余的换行符
	fmt.Scanf("%s %d", &name, &age) //输入空格隔开
	fmt.Printf("%s今年%d岁\n", name, age)
}
