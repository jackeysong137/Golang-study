package main

import (
	"fmt"
	"sync"
	"time"
)

/*
goroutine ： 就是go语言里面的协程，一个主线程可以对应多个协程，协程可以理解为用户级别的线程，对于操作系统来说不知道协程的存在，
协程我完全业务代码来调度，
在go语言中，开启一个协程的内存消耗，大概是2kb，而在java，c中开启一个线程的内存占用大概是2mb，开销差距大

	协程在主线程结束运行，协程也会结束运行，一般用snyc里面的waitgroup来实现等待协程结束
	开启协程语法：在语法前面加 go就是开启协程
*/
var wp sync.WaitGroup

func main() {
	start := time.Now().Unix()
	/*
		这里和wp.Done()成对出现 如果不是，或造成死锁或者抛出异常
		wp.Wait()表示需要等到内部计数器为0 停止等待， 如果wp.Add(1)和wp.Done不是成对出现就会出现问题

	*/

	wp.Add(1) //这里和wp.Done()成对出现 如果不是，或造成死锁或者抛出异常

	go test()
	wp.Add(1)
	go test()
	wp.Add(1)
	go test()
	wp.Add(1)
	go test()

	wp.Wait()

	end := time.Now().Unix()
	fmt.Println("耗时：", end-start)

}

func test() {

	for i := 0; i < 1000000; i++ {
		fmt.Println("打印--", i)
	}
	wp.Done()
}
