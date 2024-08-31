package main

import "fmt"

func main() {
	//1:方法的基本使用
	sum1 := sumFn1(1, 2)
	fmt.Println(sum1)
	fmt.Println("-------------------")
	//第二种
	sum2 := sumFn2(10, 2)
	fmt.Println(sum2)
	fmt.Println("-------------------")
	//第三种 返回多个值
	fmt.Println("-------------------")
	sum3, sub3 := sumAndSub(10, 5)
	fmt.Println(sum3, sub3)
	fmt.Println("-------------------")
	//第四种 返回值命名
	sum4, sub4 := sumAndSub2(10, 5)
	fmt.Println(sum4, sub4)
	fmt.Println("-------------------")
	//第五种可变参数
	sum5 := sumFn5(0, 1, 2, 3, 4, 5)
	fmt.Println(sum5)
	numSlice := []int{1, 2, 3, 4, 5}
	sum6 := sumFn5(0, numSlice...)
	fmt.Println(sum6)
}

/*
相加 返回int类型的和
*/
func sumFn1(a int, b int) int {
	sum := a + b
	return sum
}

/*
相加 返回int类型的和: 参数都是相同类型的时候 可以简写
*/
func sumFn2(a, b int) int {
	sum := a + b
	return sum
}

/*

返回多个值 多个返回值 需要用（）括起来
*/

func sumAndSub(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

/*
返回值命名 可以给返回值命名 在方法的内部可以直接使用 ， 然后直接retrun就可以了
*/

func sumAndSub2(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func sumFn5(a int, b ...int) int {
	for _, v := range b {
		a += v
	}
	return a
}
