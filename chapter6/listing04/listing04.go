package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// 单核情况下长时间执行的协程会被调度器轮流分配
func main() {
	//给调度器分配一个处理器
	runtime.GOMAXPROCS(1)

	//初始化计数器
	wg.Add(2)

	fmt.Println("两个协程开始工作")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("等待两个协程执行完成")
	wg.Wait()

	fmt.Println("两个协程执行结束")

}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 0; outer < 5000; outer++ {
		//除了2和本身以外还能被整数的就不是质数
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		//打印协程号和质数
		fmt.Printf("%s:%d\n", prefix, outer)
	}

}
