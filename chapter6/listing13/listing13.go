package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	wg      sync.WaitGroup
	counter int64
)

func main() {

	//设置计数器为协程数2
	wg.Add(2)

	//启动两个协程对counter进行自增
	go incCounter(1)
	go incCounter(2)

	//等待两个协程执行完
	wg.Wait()

	fmt.Println("执行完毕,counter:", counter)

}

func incCounter(id int) {
	//函数执行完成后 计数器减1
	defer wg.Done()

	for i := 0; i < 2; i++ {

		//将自增操作原子化解决协程安全问题
		atomic.AddInt64(&counter, 1)

		//将协程对应的线程执行权归还，并回到队列中
		runtime.Gosched()

	}

}
