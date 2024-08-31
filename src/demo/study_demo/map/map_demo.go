package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		map类型默认值为nil 必须要使用make函数初始化才可以使用
		map的定义 map[keyType]valueType

	*/
	//方式1
	var userInfo map[string]string
	userInfo = make(map[string]string, 10)
	userInfo["name"] = "张三"
	fmt.Println(userInfo)
	fmt.Println(userInfo["name"])
	var map_1 = make(map[string]int, 2)
	map_1["age"] = 18
	fmt.Println(map_1)
	//方式2
	var map_2 = make(map[string]string)
	map_2["key1"] = "张三"
	fmt.Printf("%d", len(map_2))

	//方式3
	map_3 := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	// 修改
	map_3["username"] = "lisi"
	fmt.Println(map_3)
	//查询
	fmt.Println(map_3["age"])
	//判断是非存在指定的key
	v, ok := map_3["age"] //v是值 ok 是否存在
	fmt.Println(v, ok)
	v1, ok1 := map_1["height"]
	fmt.Println(v1, ok1)

	//删除
	delete(map_3, "username")
	fmt.Println(map_3)

	//统计字符串
	var str = "hey gays how do you do"
	strSlice := strings.Split(str, " ")
	//map
	countMap := make(map[string]int)
	for _, v := range strSlice {
		countMap[v]++
	}
	fmt.Println(countMap)
}
