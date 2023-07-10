package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	//求字符串长度
	fmt.Println(len(str)) //11

	//字符串拼接 +或fmt.Sprintf

	str2 := "start " + str

	str3 := fmt.Sprintf("%s %s", str2, "end")
	name := "zss"
	age := 25
	fmtStr := "Name: %s,age:%d"
	str4 := fmt.Sprintf(fmtStr, name, age)

	fmt.Println(str2) //start Hello World
	fmt.Println(str3) //start Hello World en
	fmt.Println(str4) //Name: zss,age:25

	//字符串分割
	fmt.Println(strings.Split(str, " ")) //[Hello World]

	//是否包含
	fmt.Println(strings.Contains(str, "H")) //true

	//前后缀判断 strings.HasPrefix 、strings.HasSuffix
	fmt.Println(strings.HasPrefix(str, "Hello")) //true
	fmt.Println(strings.HasSuffix(str, "Hello")) //false

	//字符串出现的位置
	fmt.Println(strings.Index(str, "o"))     //4  第一个出现位置
	fmt.Println(strings.LastIndex(str, "o")) //7  最后一个出现位置

	//join操作 与分割相反
	slice := []string{"apple", "banana", "cherry"}
	result := strings.Join(slice, ", ")
	fmt.Println(result) // apple, banana, cherry

}
