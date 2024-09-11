package main

import (
	"fmt"
	"sync"
	"time"
)

/*
channel  管道  类似于队列，先入先入 FIFO的

	可用户线程之间的通讯，是引用数据类型  ，当channel 没有初始化 就忘里面添加元素，fatal error: all goroutines are asleep - deadlock!
	如果里面没数据 就 <- ch 也会报错 fatal error: all goroutines are asleep - deadlock!

	1: channel的定义 var ch channel
	使用 make初始化 ch := make(chan int, 3)
	 往channel里面添加数据 ch <- 10
	 取数据   <- ch
	简单记忆： 箭头指向管道就是往冠道添加数据， 箭头不指向管道，就是取数据
*/
var wg sync.WaitGroup

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 22
	ch <- 24
	fmt.Println(<-ch) //取数据

	//使用for range循环 遍历 channel  不关闭管道就会报错
	//for range 会一直循环读取channel的数据 知道channel 被关闭， 如果channel没有被关闭就会一直等待数据， 就会报错 死锁

	// for val := range ch {
	// 	fmt.Println(val)
	// }
	// 可以使用普通for来读取

	fmt.Println(len(ch))
	fmt.Println("-------------")
	//**正确的写法
	len := len(ch)
	for i := 0; i < len; i++ { //len(ch) 会变化 不应该这个写 需要 	len := len(ch)来记录
		a := <-ch
		fmt.Println(a)
	}
	fmt.Println("-------------")
	//或者
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("-------------")
	//在一个goroutine里面添加数据   一个goroutine取数据

	/*
		 ***抛出死锁的条件： 所以的goroutine都处于阻塞状态
		 	这就很好解释 fo range 遍历会抛出死锁
			1: for range 会一只阻塞等待数据 直到channel close

	*/
	ch1 := make(chan int, 5) //初始化容量只是表示channel管道最多同时容纳5个 但是取走后是可以继续存放的
	wg.Add(1)
	go add(ch1) //开启goroutine 协程
	wg.Add(1)
	go sub(ch1) //开启goroutine 协程

	wg.Wait() //等待协程完成
	fmt.Println("-----完成--------")

}

func add(ch chan int) {
	defer func() { //释放资源
		//需要关闭 如果不关闭就会造成下游的遍历channel造成阻塞 抛出死锁
		close(ch)
		wg.Done()
	}()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("添加数据", i)
		ch <- i
	}
}

func sub(ch chan int) {
	defer wg.Done()
	for val := range ch { //尽管这里在等待获获取数据 ，但是其他的gorouting没有全部处于阻塞状态，就不会报deadLock！
		fmt.Println("读取数据", val)
	}
}
