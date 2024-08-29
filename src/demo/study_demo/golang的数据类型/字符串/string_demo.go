package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//1 、 定义字符串类型
	name := "张三"
	var str1 string = "你好golang"
	var str2 = "hello"
	fmt.Printf("%v--%T\n", name, name)
	fmt.Printf("%v--%T\n", str1, str1)
	fmt.Printf("%v--%T\n", str2, str2)

	//2 多行字符串 使用反引号来实现 ``

	text := `hello boy
	 you guass will happen
	 yes lock will coming!
	`
	fmt.Println(text)
	fmt.Println("-------------------")

	//3、 字符串的常用方法 长度：len（str）
	text = "Golang"
	fmt.Printf("长度: %v\n", len(text))

	// 4、拼接字符串 使用+号 或者 fmt.Sprintf()
	// +号拼接
	username := "张三"
	age := "18"
	str3 := username + age
	fmt.Println(str3)
	//fmt.Sprintf() 拼接

	s1 := "账号"
	s2 := "密码"
	// fmt.Sprintf( 通过占位来拼接 和java的formt类似
	s3 := fmt.Sprintf("%v%v", s1, s2)
	fmt.Println(s3)
	fmt.Println("-----------")
	//通过 	strBuf := bytes.NewBufferString("") 这个方式速度比较快 这个类似于java的stringbuilder 运用在字符串多次拼接的时候优化
	strBuf := bytes.NewBufferString("")
	strBuf.WriteString(s1)
	strBuf.WriteString(s2)
	fmt.Println(strBuf)

	strLine := "第一行" +
		"第二行"

	fmt.Println(strLine)

	//字符串的方法 很多常用的方法在strings包里面
	//切割
	content := "123-456-789"
	arr := strings.Split(content, "-") //切割得到的事切片 这里go的切片和数组有区别 后续会讲
	//[123 456 789]
	fmt.Println(arr)

	//拼接数组 strings.Join()
	content = strings.Join(arr, "---")
	fmt.Println(content)

	//包含 string.Contains()
	//第一次出现的位置 Index
	//最后一次出现的位置LastIndex  这里的索引和java一样 查询不到返回-1

	content = "this is go!"
	subStr := "is"
	flag := strings.Contains(content, subStr)
	fmt.Println(flag) //true

	index := strings.Index(content, subStr) //2
	fmt.Println("is 第一次出现的位置", index)

	// 还有很多方法可以使用

}
