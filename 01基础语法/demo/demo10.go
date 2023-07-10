package main

import "fmt"

func main() {
	s1 := "hello world"
	byteS1 := []byte(s1)
	byteS1[0] = 'F'
	fmt.Println(string(byteS1)) //Fello world

	s2 := "你好"
	runeS2 := []rune(s2)
	runeS2[0] = '您'
	fmt.Println(string(runeS2)) //您好

}
