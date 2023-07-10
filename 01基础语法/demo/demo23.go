package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	//使用[:]截取切元素
	s1 := a[1:3]
	s2 := a[:]
	fmt.Println(a)  //[1 2 3 4 5]
	fmt.Println(s1) //[2 3]
	fmt.Println(s2) //[1 2 3 4 5]

	//追加元素
	s3 := append(a, 6)
	s4 := append(a, 6, 7, 8)
	s5 := append(s1, s2...)
	fmt.Println(s3) //[1 2 3 4 5 6 7 8]
	fmt.Println(s4) //[1 2 3 4 5 6 7 8]
	fmt.Println(s5) //[2 3 1 2 3 4 5]

	//删除元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[1 2 4 5]

	//切片遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	for index, value := range a {
		fmt.Println(index, value)
	}

	//切片复制copy
	b := a
	b[0] = 9
	fmt.Println(a) //[9 2 4 5]
	fmt.Println(b) //[9 2 4 5]
	c := make([]int, 5)
	copy(c, a)
	c[0] = 8
	fmt.Println(a) //[9 2 4 5]
	fmt.Println(c) //[8 2 4 5 0]
}
