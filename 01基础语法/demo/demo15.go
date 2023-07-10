package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	for i, v := range s {
		fmt.Println(i, string(v))
	}
}
