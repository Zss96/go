## fmt

fmt包是一个用于格式化输入和输出的包，它提供了一组函数和方法，用于格式化文本，打印到标准输出、读取用户输入以及其他输入输出相关的操作。

### fmt常用函数和方法

#### 打印输出函数

1.`Print`、`Printf`、`Println`用于输出到标准输出

2.`Sprint`、`Sprintf`、`Sprintln`用于将本地格式化为字符串

```go
//demo01
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
```

#### 格式化输入函数

`Scan`、`Scanf`、`Scanln`：用于从标准输入读取用户输入的文本

```go
//demo02
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
```

#### 格式化占位符：

1.`%v`：默认格式化，可以用于大多数类型。

2.`%d`、`%f`、`%s`：用于格式化整数、浮点数和字符串。

3.`%t`：用于格式化布尔值。

4.`%p`：用于格式化指针。

```go
//demo03
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
```

