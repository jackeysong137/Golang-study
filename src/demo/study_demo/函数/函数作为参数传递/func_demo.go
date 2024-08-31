package main

import "fmt"

func main() {

	/*
		函数可以作为参数传递，函数的名称也是一个变量
	*/
	//第一种
	sum1 := calc1(10, 5, sum)
	fmt.Println(sum1)

	//第二种 type关键字自定义类型 定义函数类型作为参数
	fmt.Println("------------")
	sum2 := calc2(10, 5, sum)
	fmt.Println(sum2)
	//函数可以作为返回值
	fn1 := fnChoose("-")
	if fn1 != nil {
		fmt.Printf(" %v\n", fn1(10, 5))
	}
	fn2 := fnChoose("+")
	if fn2 != nil {
		fmt.Printf(" %v\n", fn2(10, 5))
	}
	//匿名函数
	fn3 := func(a, b int) int {
		return a * b
	}
	fmt.Printf("fn3(4, 5): %v\n", fn3(4, 5))

	//匿名可执行函数 只要加载 就会执行
	func(a, b int) {
		fmt.Println(a * b)
	}(10, 10)
}

/*
函数作为参数传递：第一种定义方式 1
*/
func calc1(a, b int, fn func(x, b int) int) int {
	return fn(a, b)
}

/*函数作为参数传递：第一种写法 函数的参数太长 还可以使用type定义一个函数类型
 */
//type 可以用来自定义类型
type calc = func(a, b int) int

func calc2(a, b int, fn calc) int {
	return fn(a, b)
}
func fnChoose(opt string) calc {
	switch opt {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return nil
	}

}
func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
