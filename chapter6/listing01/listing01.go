package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main 启动一个逻辑核心，协程会并发(即先后执行)
func main() {
	//启动一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	fmt.Println("启动两个goroutine打印3次英文字母")

	//开启一个计数器，设置为2，等待两个协程执行完,从而做到流程控制
	var wg sync.WaitGroup
	wg.Add(2)

	//启动一个协程打印a-z两次
	go func() {
		//协程结束之后，将计数器减1
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for i := 'a'; i < 'a'+26; i++ {
				fmt.Printf("%c ", i)
			}
			//换行便于查看
			fmt.Println("\n")
		}
	}()

	//启动一个协程打印A-Z两次
	go func() {
		//协程结束之后，将计数器减1
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for i := 'A'; i < 'A'+26; i++ {
				fmt.Printf("%c ", i)
			}
			//换行便于查看
			fmt.Println("\n")
		}
	}()

	fmt.Println("等待两个goroutine执行完")
	//等待两个goroutine执行完
	wg.Wait()

	fmt.Println("两个goroutine执行完成")
}
