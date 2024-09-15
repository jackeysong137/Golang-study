package main

import (
	"fmt"
	"sync"
	"time"
)

/*
多线程下操作共享变量 会有线程安全问题（channel是线程安全的）
*/

var count = 1
var count1 = 1
var wg sync.WaitGroup
var mutex sync.Mutex
var rwMutex sync.RWMutex

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()

	fmt.Println(count) // 100  98  101（没加锁之前）

	/*
		 运行结果来看是错误的 ：原因 多个线程读取的时候可能读取到count的值是相同的 count++ 得到同样的值
			解决方法 互斥锁
			var mutex sync.Mutex 的lock和unlock
	*/
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go syncAdd()
	}
	wg.Wait()
	fmt.Println(count1)

	//还有读写锁  获取写锁的时候， 其他线程不能读，也不能写 获取读锁的时候 ，其他线程可以读， 但是不可以写
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writ()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println("--complete--")
}

/*
加锁  互斥锁
*/
func add() {
	count++
	wg.Done()
}

func syncAdd() {
	mutex.Lock()
	count1++
	mutex.Unlock()
	wg.Done()
}

func sub() {

	count--
	wg.Done()
}

func writ() {
	rwMutex.Lock()
	fmt.Println("----写操作")
	time.Sleep(time.Second * 3)
	rwMutex.Unlock()
	wg.Done()
}

func read() {
	rwMutex.RLock()
	fmt.Println("读操作")
	rwMutex.RUnlock()
	wg.Done()

}
