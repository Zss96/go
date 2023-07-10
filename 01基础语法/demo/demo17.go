package main

import "fmt"

func main() {
	var arr [3]int   //数组定义，并默认初始化为0值
	fmt.Println(arr) //[0,0,0]

	var arr2 = [3]int{1, 2}
	fmt.Println(arr2) //[1 2 0]
	var arr3 = [3]string{"zss", "18", "男"}
	fmt.Println(arr3) //[zss 18 男]
	var arr4 = [3]string{"zss", "18"}
	fmt.Println(arr4) //[zss 18 ] 注意，这里为空字符" "

	//让编译器根据初始值的个数自行推断数组的长度
	arr5 := [...]int{1, 2}
	fmt.Println(arr5)

	//取值:数组下标从0开始算
	fmt.Println(arr5[1]) //5

	//修改
	arr5[1] = 3
	fmt.Println(arr5) //3
}
