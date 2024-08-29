package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、数值类型转化 操一般要低类型转为搞类型  避免出现超过位数

	//注意 go是强类型语言 只有相同的类型才可以参与计算

	//2、 这里主要讲的是字符串类型转化
	//其他类型转化为字符串 有2种方法 :fmt.Sprintf 和strconv包

	a := 10
	//使用fmt 的 Sprintf进行转化
	//注意的是 浮点数是%f ，bool 是%t 字符是%c
	str := fmt.Sprintf("%d", a)
	fmt.Printf("值%v 类型:%T \n", str, str)
	//浮点数
	b := 3.12
	str = fmt.Sprintf("%f", b)
	fmt.Printf("值%v 类型:%T \n", str, str)
	//字符
	c := 'a'
	str = fmt.Sprintf("%c", c)
	fmt.Printf("值%v 类型:%T \n", str, str)

	//3 使用strconv包进行字符串和其他类型的转化
	var i = 20
	//这里的类型比心为int64才可以  第一个参数表示10进制
	str = strconv.FormatInt(int64(i), 10)

	fmt.Printf("值%v 类型:%T \n", str, str)

	num, ex := strconv.ParseInt("10", 10, 63)
	fmt.Println(num)
	fmt.Println(ex)
}
