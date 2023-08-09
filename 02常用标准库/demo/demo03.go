package main

import (
	"fmt"
)

func main() {
	// 格式化整数
	age := 25
	fmt.Printf("My age is %d\n", age) //My age is 25
	// 格式化浮点数
	pi := 3.14159
	fmt.Printf("The value of pi is %.2f\n", pi) //The value of pi is 3.14

	// 格式化字符串
	name := "Alice"
	fmt.Printf("Hello, %s\n", name) //Hello, Alice

	// 格式化布尔值
	isTrue := true
	fmt.Printf("The value is %t\n", isTrue) //The value is true

	// 格式化指针
	address := &age
	fmt.Printf("Address of age variable: %p\n", address) //Address of age variable: 0xc0000a6058

}
