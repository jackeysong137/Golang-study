package main

import "fmt"

/*
defer 定义的语句会延迟执行， 在返回的时候处理数据的时候是倒序执行的
***
注册顺序

	fn1
	fn2

执行顺序

	fn2
	fn1

注意 在defer 注册的时候 所有的参数都是已经确定好了的 后续倒叙执行的时候 参数的值和一开始的时候是一样的

在defer在时机运用中一般运用在释放资源的时候 关闭文件流等，常用余异常处理结合 recover
*/
func main() {
	fmt.Println("开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("结束")

	/*
		执行结果：
			开始
			结束
			3
			2
			1

	*/

	//一般的 defer语句都是定义一个匿名函数
	fmt.Println(fn1()) //打印0
	//defer语句在匿名返回值和命名返回值的表现不一致
	/*
		在go语言中 函数的返回不是一个原子性的操作，分2步执行 ：
			1:给返回值赋值
			2:执行RET命令
		defer的执行时机是在给返回值赋值之后
			匿名返回函数 操作的时候函数的内部变量 不会改变返回值，
			命名返回值是直接操作返回值
	*/
	fmt.Println("-------------------")
	fmt.Println(fn1()) //打印0
	fmt.Println(fn2()) //打印1
	fmt.Println("-------------------")
}

func fn1() int {
	a := 0
	defer func() {
		a++
	}()
	return a
}

func fn2() (b int) {
	defer func() {
		b++
	}()
	return b
}
