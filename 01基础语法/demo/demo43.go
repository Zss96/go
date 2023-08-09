package main

import "fmt"

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user User
	user.Name = "zss"
	user.Gender = "男"
	//匿名字段
	//当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	user.Address.Province = "广东" // 匿名字段默认使用类型名作为字段名
	user.City = "广州"             // 匿名字段可以省略
	//嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名
	user.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
	fmt.Printf("%#v", user)          //main.User{Name:"zss", Gender:"男", Address:main.Address{Province:"广东", City:"广州", CreateTime:"2000"}, Email:main.Email{Account:"", CreateTime:"2000"}}
}
