package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 首先要被转的数据 需要是一个json字符串，变成的[]byte
	// 其次转的目标 需要时 一个有标识 "json:"name"" 的结构 供其转译
	b := []byte(`"Name":"gopher","IsAdmin":false,"Followers":8900`)
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	fmt.Println(s)
}
