package main

import (
	"fmt"
)

func main() {

	//Golang 语言中 浮点数默认是float64的
	var a float32 = 3.12 // 4个字节
	fmt.Printf("a值: %v 类型:%T \n", a, a)
	fmt.Printf("a值: %.4f 类型:%T \n", a, a) //%2.f 保留4个小数

	f := 3.1415926
	//float64 和操作系统有关  32位的就是32  64 就是64
	fmt.Printf("类型:%T", f)

	//Golang中使用科学计数法来表示浮点数
	var f1 float32 = 3.12e2                //表示3.12 * 10的2次方
	fmt.Printf("f1值: %v 类型:%T \n", f1, f1) //312

	var f2 float32 = 3.12e-2               //表示3.12 除以10的2次方
	fmt.Printf("f2值: %v 类型:%T \n", f2, f2) //0.0312

	//Golang 中 浮点数 精度丢失问题 几乎所有的语言都会有这个问题 在定长条件下， 二进制小数和10进制小数互相转化或出现精度丢失
	d := 1129.6
	fmt.Println(d * 100) //112959.99999999999 预期的值是11296
	//怎噩梦解决 在java中使用的是BigDecimal Golang中使用/github.com/shopspring/decimal包来解决
	//使用decimal来决解

}
