package main

import (
	"fmt"
	"reflect"
)

/*
反射
*/
func modify(a interface{}) {
	v := reflect.ValueOf(a)

	//指针类型需要用elem来修改
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}

}

func main() {

	var a int64 = 10

	//反射获取类型名称

	typeName := reflect.TypeOf(a)
	fmt.Printf("%T\n", typeName)
	fmt.Println(typeName.Name() == "int64") //true
	fmt.Println(typeName.Kind())

	// kind ： 友好 种类  这里是种类的意思 获取种类
	//Name 是获取类型名称 字符串
	/*
		kind ： int -> int
				指针是 ptr
				结构体是 struct
				数组是：array
				切片是：slice
		Name

	*/
	var b int = 10
	var flag bool = false
	var str string = "123"
	var p person = person{
		Name: "张三",
	}
	printKind(b)    //int
	printKind(flag) //bool
	printKind(str)  //string
	printKind(p)    //struct
	printKind(&p)   //ptr
	//当传入的是指针的时候 我们可以通过 Elem来操作
	optPoint(&p)
	fmt.Println(p.Name) //私有变量不能反射访问

	// 反射执行方法

	v := reflect.ValueOf(&p)
	var params = make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf("xxx"))
	method := v.MethodByName("SetName") //注意 方法属于结构体 不呢个用elem来获取 ， 也不能获取私有的方法
	if !method.IsValid() {
		fmt.Println("Method not found")
		return
	}
	method.Call(params)
	fmt.Println(p.Name)

}
func printKind(a interface{}) {
	v := reflect.TypeOf(a)
	fmt.Printf("类型是:%v\n", v.Kind())
}
func optPoint(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Printf("类型是:%v\n", v.Elem().Kind())
	v.Elem().FieldByName("Name").SetString("hello!")

}

type person struct {
	Name string
}

func (p *person) SetName(name string) {
	p.Name = name
}
