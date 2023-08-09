package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s := Student{
		ID:     1,
		Gender: "男",
		name:   "zss",
	}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failde")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
	s1 := string(data)
	s2 := &Student{}
	//josn反序列化
	err = json.Unmarshal([]byte(s1), s2)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Println(s2) //&{1 男 }
}
