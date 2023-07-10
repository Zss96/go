package main

import "fmt"

func main() {
	//元素为map类型的切片

	mapSlice := make([]map[string]string, 3)
	fmt.Println(mapSlice)
	for i, v := range mapSlice {
		fmt.Println(i, v)
	}
	//对切片进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "zss"
	mapSlice[0]["password"] = "123456"
	for i, v := range mapSlice {
		fmt.Println(i, v)
	}
}
