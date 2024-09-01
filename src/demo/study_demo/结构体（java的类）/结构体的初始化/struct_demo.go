package main

import "fmt"

/*
		 结构体的定义语法 ：
			type 结构体名称 struct{
				变量名1 类型
				变量名2 类型
			}
	        **** 结构体是值类型！！！！***

			在golang中变量名大写表示public 小写表示private

			p2 := new(person) //这里的初始化类型是指针类型 new关键字初始化的都是指针  在golang中， 结构体的指针可以直接访问变量
			p2.name = "王五"
			p2.age = 20
			p2.sex = "男"
			底层其实是（*p2).name ="xxx"
*/
type person struct {
	name string
	age  int
	sex  string
}

func main() {

	//结构体初始化
	//方式1:
	p := person{}
	p.name = "张三"
	p.age = 19
	p.sex = "男"
	fmt.Printf("%#v ,类型：%T \n", p, p) //main.person{name:"张三", age:19, sex:"男"} ,类型：main.person
	//方式2:
	p2 := new(person) //这里的初始化类型是指针类型 new关键字初始化的都是指针  在golang中， 结构体的指针可以直接访问变量
	p2.name = "王五"
	p2.age = 20
	p2.sex = "男"
	fmt.Printf("%#v ,类型：%T \n", p2, p2) //&main.person{name:"王五", age:20, sex:"男"} ,类型：*main.person

	//方式3:
	p3 := &person{}
	p3.name = "李四"
	p3.age = 20
	p3.sex = "男"
	fmt.Printf("%#v ,类型：%T \n", p3, p3) //&main.person{name:"李四", age:20, sex:"男"} ,类型：*main.person

	//方式4:
	p4 := person{
		name: "小华",
		age:  20,
		sex:  "男",
	}
	fmt.Printf("%#v ,类型：%T \n", p4, p4) //main.person{name:"小华", age:20, sex:"男"} ,类型：main.person

	//方式5:
	p5 := &person{
		name: "小华1",
		age:  20,
		sex:  "男",
	}
	//指针类型
	fmt.Printf("%#v ,类型：%T \n", p5, p5) //&main.person{name:"小华1", age:20, sex:"男"} ,类型：*main.person

	//方式6 注意顺序
	p6 := person{
		"刘亦菲",
		20,
		"女",
	}
	fmt.Printf("%#v ,类型：%T \n", p6, p6) //main.person{name:"刘亦菲", age:20, sex:"女"} ,类型：main.person
}
