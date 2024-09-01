package main

import (
	"errors"
	"fmt"
)

/*
在golang中没有异常处理机制，需要手动去处理 通过panic和recover来解决

	1:panic在任何地方都可以发生，就是有异常就会抛出panic  需要通过recover来处理
	2:recover 只能在defer中使用
*/

func main() {

	/* 不处理异常：
	    panic: runtime error: integer divide by zero
		goroutine 1 [running]:
		main.fn1(...)
		        /Users/songzhixiong/gocode/project01/src/demo/study_demo/异常处理panic和recover/exception.go:21
		main.main()
		        /Users/songzhixiong/gocode/project01/src/demo/study_demo/异常处理panic和recover/exception.go:13 +0x20
		exit status 2
	*/
	//fmt.Println(fn1(1, 0))

	fmt.Println(fn2(1, 0))
	fmt.Println("运行结束")

	err := fn3(1)
	if err != nil {
		fmt.Println("err:", err)
	}
	ex := fn3(0)
	if ex != nil {
		fmt.Println("err:", ex)
	}
}

/*
无异常处理 如果出现异常 就会停止运行
*/
func fn1(a, b int) int {
	return a / b
}

/*
通过 自执行函数 来处理异常  内置的函数recover 可以获取到异常
*/
func fn2(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err", err)
		}
	}()
	return a / b
}

/*
通过主动抛出异常来完成 也可以 手动panic
*/
func fn3(a int) error {
	if a > 0 {
		fmt.Println("参数合法")
		return nil
	}
	return errors.New("参数不合法")

}
