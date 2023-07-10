package main

import "fmt"

func main() {
	a := [...]string{"北京", "上海", "广州", "深圳"}
	//for遍历循环
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	//for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}
