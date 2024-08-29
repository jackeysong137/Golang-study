package main

import "fmt"

func main() {

	//在golang 中 没有while循环 都是for循环代替  只有for循环

	//1、第一种写法

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println("--------------")

	//第二种写法
	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println("--------------")

	//第三种写法 这里就和while 循环是一样的
	count := 0
	for count < 10 {

		if count >= 10 {
			break
		}
		fmt.Printf("%d ", count)
		count++
	}

	fmt.Println()
	fmt.Println("--------------")

	//第四种写法
	var c = 0
	for {
		if c >= 10 {
			fmt.Println()
			break
		}
		fmt.Printf("%d ", c)
		c++
	}

}
