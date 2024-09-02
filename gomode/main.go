package main

import (
	"fmt"
	"gomode/calc" //这里是文件夹的名称 这个也可以定义别名  c "gomode/calc"  c就是别名
)

func main() {
	s := calc.Sub(7, 5)
	//使用自定义的包方法：包名.方法名
	fmt.Printf("calc.Add(0, 1): %v\n", calc.Add(0, 1))

	fmt.Println(s)
}

func init() {

	fmt.Println("main init ...")
}
