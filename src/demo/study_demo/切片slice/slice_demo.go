package main

import (
	"fmt"
)

func main() {

	/*
		 切片的定义方式 ：切片本地就是对数组的二次封装 包含：一个指针， 长度len 和容量cap
		长度：切片的元素个数
		容量：切片的起始索引到‘底层数组’的最后一个元素
		指针：指向底层数组
		 和数组的定义方式就是没有指定容量
		 ***数组类型是数字[n]string
		 ***切片类型是不带的[]string  是引用类型
	*/
	//方式1:
	var slice_1 = []string{"java", "python", "golang"}
	fmt.Printf("值:%v,类型:%T,长度：%v, 容量:%v \n", slice_1, slice_1, len(slice_1), cap(slice_1))
	fmt.Println("=========================================")
	//方式2: 这里进行了初始化 默认值是空数组
	slice_2 := []string{}
	fmt.Println(slice_2)
	fmt.Println(slice_2 == nil) //false
	fmt.Println("---------------------------")
	//方式3：这里没有进行初始化 默认值是nil 类似于null
	var slice_3 []string
	fmt.Println(slice_3) //长度和容量都是0
	fmt.Printf("值:%v,类型:%T,长度：%v, 容量:%v \n", slice_3, slice_3, len(slice_3), cap(slice_3))
	fmt.Println(slice_3 == nil) //true
	fmt.Println("---------------------------")
	//方式4:基于数组定义切片 ，因为切片是对数组的封装 可以通过切割数组
	var arr = [4]string{"java", "golang", "python", "c++"}
	//这里的含义是封装一个切片,指针指向0 长度是2 容量是 指针到底层数组的最后一个
	//这里的就是0-到原数组的最后一个 等于原数组的容量4
	var slice_4 = arr[:2] // len =2 cap：5 {"java", "golang"}
	fmt.Printf("值:%v,类型:%T,长度：%v, 容量:%v \n", slice_4, slice_4, len(slice_4), cap(slice_4))
	//上面得出结论 切片的这种语法 arr[起始索引:长度] //长度不填写默认是最后 理解为左闭右开也是可以[)
	//再次验证
	slice_5 := arr[1:] //索引为1 到数组结尾{"golang", "python", "c++"} 长度是 3 容量3
	fmt.Printf("值:%v,类型:%T,长度：%v, 容量:%v\n", slice_5, slice_5, len(slice_5), cap(slice_5))
	fmt.Println("---------------------------")
	//方式5 make() 函数创建切片
	slice_6 := make([]string, 5) //创建一个cap为5的切片 //string的默认值是空 不是nil 所以这里len也是5
	fmt.Printf("%v--%d-%d\n", slice_6, len(slice_6), cap(slice_6))
	slice_6[0] = "java"
	fmt.Printf("%v--%d-%d\n", slice_6, len(slice_6), cap(slice_6))
	//切片追加元素 append方法
	slice_6 = append(slice_6, "hello", "world")
	fmt.Printf("%v--%d-%d\n", slice_6, len(slice_6), cap(slice_6))

	/*
		   切片的扩容：
		   	1:判断当前的长度是不是小于容量 如果小于不扩容
			2:如果长度len不小于容量 需要扩容，判断原来的cap是不是小于1024  如果满足，在原来的cap上直接乘以2
			3:如果容量已经大于等于1024 每次增加1/4

		var oldCap int
		var newCap int

		if len > oldCap {
			oldCap = oldCap
		} else if oldCap < 1024 {
			newCap = oldCap << 2
		} else {
			newCap = oldCap + oldCap/4 //扩容1/4
			//指导 newCap > oldCap
		}

	*/
	//切片的复制
	slice_7 := make([]string, cap(slice_6))
	copy(slice_7, slice_6) // 复制 slice_6的数据到slice_7 这里底层是新创建一个数组， 不同的指针  改变值不影响原来的切片
	fmt.Println(slice_7)

	//切片的删除
	//在golang中 没有直接删除语法 利用append
	//合并2个切片
	fmt.Println("-----------------------")
	slice_8 := []string{"java", "php"}
	slice_9 := []string{"python", "golang"}
	slice_8 = append(slice_8, slice_9...) //追加slice_9 到slice_8
	fmt.Println(slice_8)
	//利用切割 和append的合并来实现元素的删除
	//删除 php
	fmt.Println("-----------------------")
	sliceA := []string{"java", "python", "php", "golang"}
	fmt.Println(sliceA)
	sliceA = append(sliceA[:2], sliceA[3:]...)
	fmt.Println(sliceA) //[java python golang]

}
