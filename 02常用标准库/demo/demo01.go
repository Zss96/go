package main

import "fmt"

func main() {
	name := "zss"
	age := 25
	//print
	//Print直接输出内容
	fmt.Print("Hello") //Hello
	//Printf支持格式化输出字符串
	fmt.Printf("%s 在努力学习Go", name)//zss 在努力学习Go
	//Println输出内容的结尾换行
	fmt.Println("Hello")//zss 在努力学习Go

	//Sprint
	s1 := fmt.Sprint("zss")
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("zss")
	fmt.Println(s1) //zss
	fmt.Println(s2) //name:zss,age:25
	fmt.Println(s3) //zss
}
