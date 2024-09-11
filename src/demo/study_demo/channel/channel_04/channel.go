package main

import "fmt"

/*

  1:单向管道
	channel 默认是双向的 可读可写 ch := make(chan int,10)
	 在某些情况下， 我们希望我们的管道只能读或者只能写 这时候就需要使用到单向管道
	用法：我们在定义方法的时候可以限定管道客户或者可写
	定义： 在定义管道channel的时候 加上箭头就可以
	  只读：ch:=make(<-chan int )
	  只写：ch:= make(chan<- int)

  2: 多路复用
	 某些情况我们需要同时在多个管道读取数据的时候， 除了使用多个goroutine之外 还可以使用select关键字来实现
	 语法 和swith差不多
*/

func main() {
	/*
		单向管道
	*/

	ch1 := make(chan<- int, 2)
	ch1 <- 2
	//fmt.Print(<-ch1)//不可以 send-only channel ch1 (variable of type chan<- int)
	//ch2 := make(<-chan int, 2)
	//ch2 <- 1 //receive-only channel ch2 (variable of type <-chan int)

	ch3 := make(chan int, 3)
	put(ch3)
	read(ch3)

	/*
		  select 多路复用:同时从channel_A和channel_B读取数据
		  	并发执行 从多个channel获取数据 ，等待所有的数据处理完成 就执行default的部分
			*** 在用select 操作的时候 不要关闭channel  会导致select一直在读取到零值 ，而for range 在读取到关闭标识就会退出
			seelct 可以处理零值 也可以读取到关闭标识

	*/

	channel_A := make(chan int, 3)
	channel_B := make(chan int, 3)
	// put(channel_A)
	// put(channel_B) 这里会关闭channel 导致死循环 即使处理了关闭标识 也执行不了default的逻辑
	for i := 1; i <= 3; i++ {
		channel_A <- i
	}
	for i := 1; i <= 3; i++ {
		channel_B <- i
	}

	for {
		select {
		case v, ok := <-channel_A:
			if !ok {
				fmt.Println("channel_A 被关闭了")
				return
			}
			fmt.Println("从channel_A读取数据:", v)
		case v := <-channel_B:
			fmt.Println("从channel_B读取数据:", v)
		default:
			fmt.Println("所有的channel的数据读取完成")
			return //不要用break break只是退出了select循环
		}

	}

}

// 单向管道  只读
func read(ch <-chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

// 只写
func put(ch chan<- int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("插入数据失败", err)
		}
		close(ch)
	}()
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	//异常
	panic("测试抛出异常")
}
