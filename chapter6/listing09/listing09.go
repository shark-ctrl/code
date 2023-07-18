package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

// 参数 CGO_ENABLED=1
//编译参数 -race
/**
==================
WARNING: DATA RACE
Write at 0x00000063c508 by goroutine 7:
  main.incCounter()
      F:/github/code/chapter6/listing09/listing09.go:49 +0xa4
  main.main.func1()
      F:/github/code/chapter6/listing09/listing09.go:25 +0x30

Previous read at 0x00000063c508 by goroutine 8:
  main.incCounter()
      F:/github/code/chapter6/listing09/listing09.go:40 +0x84
  main.main.func2()
      F:/github/code/chapter6/listing09/listing09.go:26 +0x30

Goroutine 7 (running) created at:
  main.main()
      F:/github/code/chapter6/listing09/listing09.go:25 +0x44

Goroutine 8 (finished) created at:
  main.main()
      F:/github/code/chapter6/listing09/listing09.go:26 +0x50
==================
Final Counter: 2
Found 1 data race(s)

Process finished with the exit code 66
*/
// main 模拟多线程之间协程操作counter的线程安全问题
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

	//fmt.Println("协程", id, "开始工作")

	for i := 0; i < 2; i++ {
		value := counter

		//fmt.Println("协程", id, "读取到的counter:", counter)
		//将协程对应的线程执行权归还，并回到队列中
		runtime.Gosched()

		value++

		//fmt.Println("协程", id, "覆盖counter值,value:", value)
		counter = value

		//fmt.Println("协程", id, "完成value自增，自增后的值为", value)
	}

}
