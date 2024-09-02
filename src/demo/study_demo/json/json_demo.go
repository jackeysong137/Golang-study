package main

import (
	"encoding/json"
	"fmt"
)

/*
json包无法 访问 私有字段 反序列化的时候也是一样的
*/
type Student struct {
	Username string `json:"name"` // 可以指定序列化和反序列化的时候的别名
	Age      uint
}

func main() {
	s := Student{
		Username: "hello",
		Age:      18,
	}
	strByte, _ := json.Marshal(s)
	jsonStr := string(strByte)
	fmt.Println(jsonStr)              //{"Username":"hello","Age":18}
	s1 := `{"name":"hello","Age":18}` //可以使用反引号来定义字符串
	//解析
	var stu Student
	json.Unmarshal([]byte(s1), &stu)
	fmt.Println(stu)
}
