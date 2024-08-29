package main

import (
	"fmt"
)

func main() {

	// 运算符 + - * / % 和java一样
	//注意 在golang中 ++ 和 -- 是单独的预语法 只能单独使用
	a := 10
	a--
	//9
	fmt.Println(a)
	// 不可以书写 a -- + a 这种情况  没有了java中那种繁琐的计算 分析
	// 且没有  --a  这种写法 只有  a-- 这种
	a++
	//10
	println(a)

	/*
		逻辑运运算符 和java的一样的 < > <= >= && || != 这些
	*/

	//取余数
	num := 10
	c := 5
	res := c % num
	//5
	fmt.Println(res)

	//位运算
	x := 1
	y := 2
	//异或 实现交换
	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Printf("x=%d,y=%d\n", x, y)
}
