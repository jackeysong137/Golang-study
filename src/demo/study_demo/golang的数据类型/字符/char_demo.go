package main

import (
	"fmt"
)

func main() {

	// 1、 golang 中 字符属于int(uint8)类型 或者叫byte类型 单引号表示
	// byte类型处理的是ASCII编码的  rune类型处理的是UTF-8编码的 在Golang中  汉字默认UTF-8的编码 也就是rune类型
	var a uint8 = 'a'
	fmt.Printf("值: %v 类型: %T\n", a, a)

	//2、还有一种是rune类型 表示的是UTF-8编码 默认是int32
	str := "你好！Gonglang"

	//3、使用这个方式遍历的值就是byte类型 编码是ASCII 中文会出现乱码  %c 是原样输出  %v获取的是编码值
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c(%T)", str[i], str[i])
	}
	fmt.Println("--------------")
	//4、这个方式遍历的是rune 类型 得到的是utf-8编码值
	for _, v := range str {
		fmt.Printf("%c(%T)", v, v)
	}
	fmt.Println("--------------")
	//5、修改字符串
	content := "你好,Golanng"
	// 字符串不可以直接修改 ，是不可变的 只能通过先把字符串改为 字符数组 再修改字符， 最后在转化为字符串

	// 中文类型 UTF-8编码  使用rune类型
	//强制类型转化
	runeStr := []rune(content)
	runeStr[0] = '很'
	fmt.Println(string(runeStr))

	// 字符串没有中文 可以使用byte数组来实现
	content = "hello Golang!"
	byteStr := []byte(content)
	byteStr[0] = 'H'
	fmt.Println(string(byteStr))

}
