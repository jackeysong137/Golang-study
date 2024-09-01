package main

import "fmt"

/*
  在golang中  给每一个变量 分配内存的同时， 都会有一个指针指向这个内存地址 ，指针也是一个变量 取变量的内存地址 &变量
  //指针是一个引用类型
	var a =10
	p := &a //这个就是指针  &a就是获取 a的地址值 也是指针的值
	*p 可以得到指针锁指向内存的值
	*p =100 修改值 这个时候 a也变为100 了
*/

func main() {
	a := 10
	fmt.Println(&a) //内存地址 是一个指针类型 0x14000104020

	//定义指针
	p := &a
	fmt.Printf("%v,%T \n", p, p) //0x1400000e0a0,*int int类型的指针
	//获取指针对应内存地址的存储的值
	fmt.Println("value=", *p)
	fmt.Println("-----------------")
	//通过操作指针来改变值
	*p = 100
	fmt.Println("原变量的值：", a) //100 原变量的值也改变了

	b := 5
	fn1(&b)
	fmt.Println("b==", b) //100 这样可以实现 值类型的值在方法中也可以改变
	fmt.Println("-----------------")
	arr := [3]uint{1, 2, 3}
	fmt.Println("arr修改前:", arr)
	fn2(&arr)
	fmt.Println("arr修改后:", arr) //实现修改数组的值

}

/*
入参是int类型的指针 给这个指针的指向的内存赋值为100
*/
func fn1(p *int) {
	*p = 100
}

func fn2(p *[3]uint) {
	(*p)[0] = 100
}
