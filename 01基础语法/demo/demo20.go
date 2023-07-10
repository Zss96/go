package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	arr1 := arr
	arr1[0] = 2
	fmt.Println(arr)  //[1 2 3]
	fmt.Println(arr1) //[2 2 3]

	changeArr(arr)
	fmt.Println(arr) //[1 2 3]
}

func changeArr(x [3]int) {
	x[0] = 9
	fmt.Println("x", x) //[9 2 3]
}
