package main

import "fmt"

func main() {
	//算数运算符
	a := 10
	b := 3
	fmt.Println("算数运算符")
	fmt.Println("a+b=", a+b) //13
	fmt.Println("a-b=", a-b) //7
	fmt.Println("a*b=", a*b) //30
	fmt.Println("a/b=", a/b) //3

	//关系运算符
	fmt.Println("关系运算符")
	fmt.Println("a == b:", a == b) //false
	fmt.Println("a != b:", a != b) //true
	fmt.Println("a > b:", a > b)   //true
	fmt.Println("a < b:", a < b)   //false
	fmt.Println("a >= b:", a >= b) //true
	fmt.Println("a <= b:", a <= b) //false

	// 逻辑运算符
	x := true
	y := false
	fmt.Println("逻辑运算符")
	fmt.Println("x && y:", x && y) //false
	fmt.Println("x || y:", x || y) //true
	fmt.Println("!x:", !x)         //false
	fmt.Println("!y:", !y)         //true

	// 位运算符
	fmt.Println("位运算符")
	num1 := 5                               // 二进制：101 =>十进制 5
	num2 := 3                               // 二进制：011 =》3
	fmt.Println("num1 & num2 =", num1&num2) // 位与  001 => 1
	fmt.Println("num1 | num2 =", num1|num2) // 位或  111 =>7
	fmt.Println("num1 ^ num2 =", num1^num2) // 位异或 110=> 6
	fmt.Println("num1 << 1 =", num1<<1)     // 101 左移一位 =1010 =>10
	fmt.Println("num2 >> 1 =", num2>>1)     //  011右移 1一位 =001=>1

	// 赋值运算符
	fmt.Println("赋值运算符")
	c := 5
	c += 2
	fmt.Println("c += 2 => c =", c) //7
	c -= 2
	fmt.Println("c -= 2 => c =", c) //5
	c *= 2
	fmt.Println("c *= 2 => c =", c) //10
	c /= 2
	fmt.Println("c /= 2 => c =", c) //5
	c %= 2
	fmt.Println("c %= 2 => c =", c)  //1
	c <<= 2                          //十进制  二进制=>二进制结果=>十进制
	fmt.Println("c <<= 2 => c =", c) //4   001=>100==>4

	c >>= 1
	fmt.Println("c >>= 1 => c =", c) //2 100 =>10=>2

	d := 3                 //11
	c &= d                 //十进制 二进制运行=>二进制=>十进制
	fmt.Println("c&=d", c) //2  11&10=>10 =>2
	c |= d
	fmt.Println("c|=d", c) //3 10|11 =>11 =>3
	c ^= d
	fmt.Println("c^=d", c) //0 11^11 =>00=>0
}
