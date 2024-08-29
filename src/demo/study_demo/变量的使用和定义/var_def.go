package main

import (
	"fmt"
)

var global_cons = "全局变量"

func main() {

	/*
		在go中变量的命名使用的是驼峰命名的方式， 小拖峰表示私有 privte 大拖峰表示public
		遇到特有名称的时候需要更具是否私有或者公有 去全部大写或者全部小写 如：DNS
	*/

	//第一种  var 变量名 类型  = 值 变量名推介
	// go 语言中  定义变量和导入的包必须使用
	// var a int = 10
	// var b int = 20
	// var c int = 30

	// fmt.Printf("a=%v,b=%v,c=%v", a, b, c)

	//短变量的方式申明  这个只能声明局部量的时候时候， 全局变量必须使用 var来定义
	a := 10
	b := 20
	c := 30
	fmt.Printf("a=%v, b=%v,c=%v\n", a, b, c)

	fmt.Printf("全局变量%v\n", global_cons)

	//变量没有初始化的话 值为空
	var username string
	fmt.Printf("username=%v", username)
	username = "张三" //初始化
	fmt.Printf("username=%v \n", username)

	//第二种
	var username1 string = "zhangsan"
	fmt.Println(username1)
	//第三种 使用类型推导
	var username2 = "李四"
	fmt.Println(username2)
	// 多个变量一起申明

	var a1, a2 string
	a1 = "zhang"
	a2 = "ces "

	fmt.Println(a1, a2)

	// 一次申明多个变量和类型

	var (
		name string
		age  int
		sex  string
	)

	name = "escanor"
	age = 18
	sex = "man"
	fmt.Println(name, age, sex)
	//定义多个类型
	cos1, cos2, cos3 := 10, 20, "c"

	fmt.Printf("cos1的类型%T cos2的类型%T cos3的类型%T", cos1, cos2, cos3)

	// 使用匿名变量：使用场景 如果我们调用多返回值的方法 需要忽略某个变量的时候 可以使用匿名变量：_
	//匿名变量不会占用命名空间 不会分配内存，所以匿名变量之间不存在变量名冲突
	//只需要名字的时候
	var sname, _ = getUserInfo()
	fmt.Println(sname)
	//只想要age
	_, sage := getUserInfo()
	fmt.Printf("age=%v", sage)

}

/*
函数可以支持多返回值
*/
func getUserInfo() (string, int) {

	return "张三", 18
}
