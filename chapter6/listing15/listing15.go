package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

// main 使用原子工具包实现安全读和安全写
func main() {
	//计时器设置为2
	wg.Add(2)

	fmt.Println("两个协程开始工作.....")
	go doWork("A")
	go doWork("B")

	//休眠1s让两个协程多工作一会
	time.Sleep(1 * time.Second)

	fmt.Println("main准备停止两个协程")

	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Println("协程", name, "开始工作")

		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("收到停止信号，协程", name, "停止工作")
			break
		}
	}

}
