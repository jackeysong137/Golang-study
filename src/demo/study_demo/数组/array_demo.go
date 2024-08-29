package main

import "fmt"

func main() {

	/*
		 数组的定义 var 变量名 [数组大小]类型
		 var arr [5]string 这里是申明了 容量大小为5的数组 数组里面是string类型的元素
		 数组的大小一旦定义了 就不可以更改
			******
			在golang中 数组是值类型， 当我们使用数据作为入参的时候， 传递的时候传递的是数组的副本，
			会造成额外的内存开销， 可以使用切片，切片传递的是数组的引用
			******
	*/
	//在golang中  数组的长度也是类型的一部分
	var arr [5]string //类型为[5]string
	arr[0] = "123"
	fmt.Println(arr)

	var arr1 [7]int //类型为[7]int
	arr1[0] = 1
	fmt.Println(arr1)
	//修改值后  没有变化  因为数组是值类型传递  传递的是数组值的副本 没办法改变值的
	updatedArr(arr1)
	fmt.Println("====", arr1)
	//数组的初始化
	// 第一种
	var arr2 [5]string //类型为[5]string
	arr2[0] = "123"

	fmt.Println(arr2)

	//第二种  指定长度的同时  初始化元素
	var strArr = [3]string{"1", "2", "3"}
	fmt.Println(strArr)

	//第三种 长度不用指定， 使用[...]代替  编译器自动推段
	var strArr2 = [...]string{"php", "java", "golang", "python"}
	fmt.Printf("长度:%v 值:%v\n", len(strArr2), strArr2)

	//第四种：可以指定下标的值 不常用
	var strArr3 = [...]string{1: "11", 2: "22", 4: "33", 7: "55"}
	fmt.Printf("长度:%v 值:%v\n", len(strArr3), strArr3)

	//遍历
	for i := 0; i < len(strArr2); i++ {
		fmt.Println(strArr2[i])
	}

	//遍历2
	for idx, val := range strArr2 {
		fmt.Printf("索引:%v值：%v ,", idx, val)
	}

}
func updatedArr(arr [7]int) {
	arr[0] = 100
}
