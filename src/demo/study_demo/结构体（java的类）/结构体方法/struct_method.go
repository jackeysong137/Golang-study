package main

import "fmt"

/*
在golang中 可以通过接收者方法给结构体定义方法

格式：	func (接收者变量 接收者类型) 方法名（变量 变量类型）（返回值）{

	}
	这里的接受者变量就等同于java的this python的self
*/

type Person struct {
	Name string
	Age  int
	Sex  string
}

/*
(p Person)表明只有Person对象可以使用这个方法 直接调用不了
*/
func (p Person) PrintInfo() {
	fmt.Printf("name=%v,age=%v \n", p.Name, p.Age)
}

/*
通过指针来修改属性的值
*/
func (p *Person) setName(name string) {
	p.Name = name
}

func main() {
	p1 := Person{
		Name: "张三",
		Age:  20,
	}
	//PrintInfo() 使用不了 编译就会报错
	p1.PrintInfo() //name=张三,age=20
	p2 := p1
	p2.Name = "李四" //修改p2 不影响p1 因为结构体是值类型

	p1.PrintInfo()   //name=张三,age=20  没有发生改变 说明结构体是值类型 p2改变的只是自己的副本
	p1.setName("王五") //通过指针来修改
	p1.PrintInfo()   //name=王五,age=20
}
