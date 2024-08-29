package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 1 定义 int类型  默认值为0
	var num1 int = 3
	fmt.Printf("num1的值%v 类型%T \n ", num1, num1)

	// 2  int8的范围 （-127 ～ 128 ）等同于 java的byte类型 占用1个字节
	var num2 int8 = 2
	fmt.Printf("num2的值%v 类型%T \n", num2, num2)
	fmt.Println(unsafe.Sizeof(num2))

	//3 int16 等同java的short 占用2个字节
	var num3 int16 = 130
	fmt.Printf("num3的值%v 类型%T  \n", num3, num3)

	//4 使用 unsafe.SizeOf 来查询不同长度整型 在内存的存储空间
	fmt.Println(unsafe.Sizeof(num3))

	//5 int32 等同java的int 4个字节（一个字节8个bit位）
	var num4 int32 = 2999
	fmt.Println(unsafe.Sizeof(num4))

	//6 int64 等同java的long 8个字节（一个字节8个bit位）
	var num5 int64 = 2222
	fmt.Println(unsafe.Sizeof(num5))

	/*
		上面的是有符号整型 ，接下来下面介绍无符号整型

		uint8 取值范围（0 -255）占一个字节 因为无符号 所以比有符号的多了一个符号位的bit 0-2^ 8
		uint16 0~65536 2个字节
		uint32 0 ~2^32 4个字节
		uint64 0~2^64 8个字节
	*/

	var n1 uint8 = 130
	fmt.Printf("n1值%v 类型%T\n", n1, n1)
	fmt.Println(unsafe.Sizeof(n1))

	// 7 int 的类型转换

	var a1 int8 = 10
	var a2 int32 = 20
	//类型转化才可以相加
	fmt.Println(int32(a1) + a2)

	// 使用fmt的%d的是10进制输出 %o八进制 %x 16进制
	//%v 是原样输出

	fmt.Printf("10进制:%d 2进制:%b 八进制:%o 十六进制:%x \n", a1, a1, a1, a1)

}
