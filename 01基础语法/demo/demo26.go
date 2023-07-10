package main

import "fmt"

func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "广东"
	v, ok := sliceMap[key]
	if !ok {
		v = make([]string, 0, 2)
	}
	v = append(v, "广州", "珠海")
	sliceMap[key] = v
	fmt.Println(sliceMap) //map[广东:[广州 珠海]]
}
