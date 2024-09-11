package main

import (
	"fmt"
)

/*
空接口: 就是 interface 里面没有任何的方法， 所以所有的类型都是他的实现（子类），有点想java中的object的意思

	在项目中可以通过空接口来实现一个方法的参数多类型话， map的值可以存储不同类型的值
	类型断言：
		1:判断变量的类型 a.(T)  返回变量的值和是否断言成功  v,ok =a.(string)  v 是a的值，ok表示是否是字符串
		2: a.(type) 获取变量的类型  这个语法只能在switch语法中使用
	结构体可以同时实现多个接口：就是吧多个接口的方法都实现一遍
	接口可以嵌套（继承）： 和结构体的匿名变量类似

	*** 实现接口方法的时候 接收者方法传入的时候指针类型，不是结构体类型，
	*** 那么把结构体赋值给接口变量的时候，结构体也是需要时是指针类型,否则就不可以实现，而 接收者方法是结构体类型，则没有限制  2种都可以
	例子：
		//实现接口的接收者方法，接收者是指针类型
	d := Dog{
		Name: "大黄",
	}
	var an animal = d //编译失败 这里需要是指针
	an.setName("小黄")
	type animal interface {
	setName(name string)
    }
	type Dog struct {
		Name string
	}
	// 指针类型
	func (d *Dog) setName(name string) {
	d.Name = name
}


*/

/*
接口interfaceAer 嵌套了interfaceBer和interfaceCer 想要实现inteafaceAer 就需要实现b和c的方法
*/
type interfaceAer interface {
	interfaceBer
	interfaceCer
}
type interfaceBer interface {
}
type interfaceCer interface {
}

func main() {
	var a interface{} //空接口
	a = "123"         //a可以是任何类型的值
	fmt.Println(a)

	a = true
	fmt.Println(a)
	//断言 判断类型并返回值
	if v, ok := a.(bool); ok {
		fmt.Println("a 是bool类型：", v)
	}
	fmt.Println("-----------")
	// mak初始化map
	//map 保存不同类型的值
	userInfo := make(map[string]interface{})
	userInfo["username"] = "张三"
	userInfo["age"] = 18
	userInfo["married"] = true
	fmt.Println(userInfo)

	for k, value := range userInfo {
		fmt.Println(k, value)
	}

	//
	Predicate("1")
	Predicate(2)
	Predicate(true)

	fmt.Println("-----------")
	//实现接口的接收者方法，接收者是指针类型
	d := Dog{
		Name: "大黄",
	}
	var an animal = &d //需要指针
	an.setName("小黄")
	fmt.Println(d.Name)

	//golang是强类型语言 所以空接口是没办法直接使用子类的方法和熟悉 需要类型转化  （类型断言）

	userInfo["pet"] = Dog{
		Name: "阿奇",
	}

	//fmt.Println(userInfo["pet"].Name)//userInfo["pet"].Name undefined (type interface{} has no field or method Name)
	//这里想要访问dog的name熟悉 需要转化为Dog类型
	pet, ok := userInfo["pet"].(Dog) //类型断言其实也是类型转化
	if ok {
		fmt.Println(pet.Name)
	}
}

type animal interface {
	setName(name string)
}

type Dog struct {
	Name string
}

func (d *Dog) setName(name string) {
	d.Name = name
}

func Predicate(a interface{}) {
	//a.(type)
	switch a.(type) {
	case string:
		fmt.Println("字符串类型")
	case int:
		fmt.Println("int 类型")
	default:
		fmt.Println("其他类型")
	}
}
