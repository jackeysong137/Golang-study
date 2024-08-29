package main

import "fmt"

func main() {

	// goto 语法结合标签使用 可以跳转到指定的标签   下面的案例会无限的打印 :
	/*
			0 1 2 3 4
		    第n行
	*/
lable1:
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println()
			goto lable2
		}
		fmt.Printf("%d ", i)
	}
	//这里会被跳过
	fmt.Println("第一行")
	fmt.Println("第二行")
	fmt.Println("第三行")
lable2:
	fmt.Println("第n行")
	goto lable1 //回到lable1
}
