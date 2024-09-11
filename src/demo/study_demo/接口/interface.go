package main

import "fmt"

/*
	 1:接口定义 tpye 接口名 interface ：定义接口的时候一班名字后面+er
		接口就是为了规范， 定义了接口 一般就是使用结构体或者自定义的类型取实现接口，
		实现接口必须实现接口的所有方法
		实现的语法：
				有接口： Aer
				结构体 B
				实现语法：
					var b = B{}
					var a Aer =b 这样就是B实现了Aer，B必须要实现Aer的所有方法
*/

type Usber interface {
	start()
}

type Phone struct {
	Name string
}

type Computer struct {
	cpu uint
}

func (c Computer) start() {
	fmt.Println(c.cpu)
}

// 实现方法
func (p Phone) start() {
	fmt.Println(p.Name)
}

func printInfo(u Usber) {
	u.start()
}

func main() {
	p := Phone{
		Name: "苹果手机",
	}
	var usb Usber = p //这里就是Phone 实现了Usber  如果Phone没有实现start方法 就会报错
	usb.start()
	printInfo(p) //这里入参是Usber 可以直接传Phone 结构体 ，这里会默认实现的  只要实现了方法就可以

	c := Computer{
		cpu: 10,
	}
	var u1 Usber = c //实现
	u1.start()
}
