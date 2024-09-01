package main

import "fmt"

/*
在golang中  结构体是可以嵌套的

 ***在结构体中没有指定变量名的时候  变量名默认是类型的名称
*/

type Person struct {
	Name    string
	Age     int
	Address Address //这里也可以是指针
	House           //没有指定变量名 默认是类型的名称 切父类（结构体） Person可以直接访问House的变量 p.price
}

type Address struct {
	Location string //这里大写 就是public
}
type House struct {
	price uint
}

func main() {
	var p Person
	p.Name = "张三"
	p.Age = 20
	p.Address.Location = "三门路"

	fmt.Printf("%#v \n", p) //main.Person{Name:"张三", Age:20, Address:main.Address{location:"三门路"}}
	p.price = 12            //先去本结构体查找price变量 再去匿名结构体查找

	/*
		结构体的继承: 利用匿名变量的方式 父结构体可以直接访问嵌套的匿名结构体的属性 实现继承的效果
	*/
	d := Dog{
		age: 2,
		Animal: Animal{
			name: "阿奇",
		},
		Cat: Cat{
			eat: "吃屎",
		},
	}
	d.PrintName() //阿奇
	d.PrintEat()  //吃屎
	d.setName("傻狗")
	d.PrintName() //傻狗

}

type Animal struct {
	name string
}

func (a Animal) PrintName() {
	fmt.Println("name= \n", a.name)
}

/*
这里Dog 继承了 Animal 和Cat 的属性和方法
*/
type Dog struct {
	age int
	Animal
	Cat
}

func (d *Dog) setName(name string) {
	d.name = name //继承Animal
}

type Cat struct {
	eat string
}

// 接收者方法 接收者是cat结构体
func (c Cat) PrintEat() {
	fmt.Println("eat= \n", c.eat)
}
