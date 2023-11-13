package main

// json 序列化与反序列化
// 只需要结构体第一个字母是大写
import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"` // 控制AGE字段的输出
	Hobby []string
}

func main() {
	a := userInfo{"wang", 18, []string{"Golang", "TypeScipt"}}
	buf, _ := json.Marshal(a)

	fmt.Println(string(buf))

	var b userInfo
	_ = json.Unmarshal(buf, &b)
	fmt.Printf("%+v", b)
}
