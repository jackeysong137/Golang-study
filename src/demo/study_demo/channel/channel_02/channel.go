package main

import (
	"fmt"
	"sync"
	"time"
)

/*
统计 1-120000之间的素数 prime number
1: 素数 就是质数

	开启一个goroutine 往 intChannel 里面放入 1-120000数字
	开启一个goroutine 计算 intChannel里面的素数 计算出的素数放入 prime channel
	开启一个goroutine 大于prime channel的数字
*/
var wg sync.WaitGroup

func putInt(ch chan int) {
	defer func() {
		close(ch)
		wg.Done()
	}()
	for i := 1; i <= 1200000; i++ {
		ch <- i
	}
	//关闭
}

func calc(putInt chan int, primeChannel chan int, exitChannel chan bool) {
	// 释放资源 defer 在返回之前执行  延迟执行的 exitChannel每个goroutine的素数计算完成
	//这里不能直接关闭primeChannel 因为这里需要开启多个gorotine来计算 如果关闭 其他的goroutine就不能再往primeChannel里面放入数据
	defer func() {
		exitChannel <- true
		wg.Done()
	}()
	for v := range putInt {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChannel <- v
		}
	}
}

func PrintlnPrimeNumber(primeChannel chan int) {
	defer wg.Done()
	for val := range primeChannel {
		fmt.Println(val)
	}
}

func main() {

	intChannel := make(chan int, 10)
	primeChannel := make(chan int, 10)
	exitChannel := make(chan bool, 16)

	start := time.Now().Unix()
	wg.Add(1)
	//添加数据
	go putInt(intChannel)
	//计算
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go calc(intChannel, primeChannel, exitChannel) //开启16 个协程来计算
	}
	wg.Add(1)

	//打印
	wg.Add(1)
	go PrintlnPrimeNumber(primeChannel)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChannel //阻塞 直到16个协程的数据都计算完成 close primeChannel，防止PrintlnPrimeNumber 死锁
		}
		close(primeChannel) //关闭
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("执行完成", end-start, "秒") //7秒完成 如果不开启goroutine协程  耗时很长 具体可见channel_03/
}
