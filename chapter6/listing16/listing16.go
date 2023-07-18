package main

import (
	"fmt"
	"sync"
)

var (
	counter int64
	//计数器，可以理解为Java的countdownLactch
	wg sync.WaitGroup
	//互斥锁
	mutex sync.Mutex
)

func main() {
	fmt.Println("main方法开始工作")

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	fmt.Println("等待两个协程结束........")
	wg.Wait()
	fmt.Println("协程运行结束，counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	fmt.Println("协程", id, "开始工作")
	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			value := counter
			value++
			counter = value
			fmt.Println("协程", id, "上锁成功并修改值成功，counter:", value)
		}
		mutex.Unlock()
	}

}
