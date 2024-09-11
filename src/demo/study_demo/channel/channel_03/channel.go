package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <= 1200000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if i%num == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num)
		}
	}
	end := time.Now().Unix()

	fmt.Println("执行完成：", end-start, "秒")
}
