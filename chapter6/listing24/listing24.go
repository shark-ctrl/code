package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //协程数设置为4
	taskLoad         = 10 //任务数为10
)

var wg sync.WaitGroup

func main() {

	wg.Add(4)

	//创建一个有缓冲通道
	taskChannel := make(chan string, taskLoad)

	for i := 0; i < numberGoroutines; i++ {
		go Worker(i, taskChannel)
	}

	//往通道里提交任务
	for i := 0; i < taskLoad; i++ {
		taskChannel <- fmt.Sprintf("task %d", i)
	}

	//关闭通道
	close(taskChannel)

	//等待4个协程执行完成
	wg.Wait()

}

func Worker(workNo int, taskChannel chan string) {
	//函数退出时扣减计数器
	defer wg.Done()

	for {
		task, ok := <-taskChannel

		if !ok {
			fmt.Println("任务通道已关闭,worker", workNo, "退出")
			return
		}

		fmt.Println("worker", workNo, "执行任务", task)

		time.Sleep(2000 * time.Millisecond)
	}

}
