package main

import "fmt"

func main() {
	//map基本使用
	m := make(map[string]int, 8)
	m["张三"] = 90
	m["李华"] = 100
	fmt.Println(m)                 //map[张三:90 李华:100]
	fmt.Println(m["张三"])           //90
	fmt.Printf("typeof a %T\n", m) //typeof a map[string]int

	//在什么时，填充元素
	userInfo := map[string]string{
		"username": "zss",
		"password": "123456", //注意，在go中，最有以为也要有,
	}
	fmt.Println(userInfo) //map[password:123456 username:zss]

	//判断某个键是否存在
	v, ok := userInfo["username"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("该用户不存在")
	}

	//map遍历
	//遍历的顺序与添加键值对的循环无关
	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	//删除键值对 delete(map,key)
	delete(userInfo, "username")

	for k, v := range userInfo {
		fmt.Println(k, v)
	}

}
