package main

import (
	"fmt"
)

func main() {
	//常量的定义使用 const 定义后不可以再次修改常量的值 类似 java 的final
	const a = "x"
	fmt.Println(a)

	//定义常量的时候必须给他赋值 切可以不用使用  变量是必须要使用
	const pi = 3.1415826
	//一次性也可以申明多个常量
	const (
		A = "a"
		B = "b"
	)

	fmt.Printf("A=%v ,B=%v \n", A, B)
	//r如果定义多个常量的时候省略了后面的值，则值和前面的值相同
	// N1= N2=N3 =USA
	const (
		N1 = "USA"
		N2
		N3
	)
	// N1= N2=N3 =USA
	fmt.Println(N1, N2, N3)
	fmt.Println("---------------")
	//常量 const 结合 iota 的使用
	/*
		iota 是golang 语言的常量计数器 只能在常量表达式中使用
		iota 在 const 关键字出现的时候会重置为0（const内部第一行之前），const 中没新增一个常量 计数+1

	*/

	const cons1 = iota
	fmt.Println("cons1=", cons1)

	const (
		CN1 = iota
		CN2
		CN3
		CN4
	)
	//打印 0 1 ，2 3
	fmt.Println(CN1, CN2, CN3, CN4)

	// 可以使用匿名变量_跳过某些值
	const (
		CN5 = iota
		CN6
		_
		CN7
	)
	//打印 0 1，3
	fmt.Println(CN5, CN6, CN7)

	//多个iota定义在一行
	const (
		C1, C2 = iota + 1, iota + 2
		C3, C4
		C5, C6
	)
	//打印出 1 2 2 3 3 4 -> 下面比上面一行多1
	fmt.Println(C1, C2, C3, C4, C5, C6)
	
}

