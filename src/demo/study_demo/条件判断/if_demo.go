package main

import "fmt"

func main() {

	//if 语法1:
	/*

		    if 条件 {

			 执行逻辑
			}
	*/

	
	//语法2：
	a := 10
	if a > 6 {
		fmt.Println("a>6")
	} else {
		fmt.Println("a < 6")
	}

	/*

		if a := 10; a>2{

		执行逻辑
		}
	*/

	if b := 10; b == 5 {
		fmt.Println("b ==5匹配成功")
	} else if b == 10 {
		fmt.Println("b==10匹配成功")
	} else {
		fmt.Println("没有匹配成功")
	}

	//switch 可以写break 可以不写 不会穿透，如果需要穿透 使用 fallthrough 就可以穿透到下一层 实际上使用很少

	type1 := "1"
	switch type1 {
	case "2":
		fmt.Println("2")
	case "1", "3":
		fmt.Println("1,3")
	default:
		fmt.Println("---")
	}


}
