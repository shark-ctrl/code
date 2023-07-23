package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 基于无缓冲通道模拟接力赛跑
func main() {

	wg.Add(1)

	baton := make(chan int)

	//启动协程1模拟第一个起跑运动员
	go Runner(baton)

	fmt.Println("接力赛开始.......")
	//向通道发送数据，模拟开枪通知第1棒起跑
	baton <- 1

	//等待第4个协程wg.Done()
	wg.Wait()
	//关闭通道
	close(baton)
	fmt.Println("接力赛结束")

}

func Runner(baton chan int) {

	var newRunner int
	//等待接力棒
	runner := <-baton

	fmt.Println("选手", runner, "接到第", runner, "棒")

	if runner != 4 {
		newRunner = runner + 1
		fmt.Println("选手", runner, "准备将接力棒交给选手", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Println("第", runner, "选手到达终点")
		wg.Done()
		return
	}

	fmt.Println("选手", runner, "将接力棒交给选手", newRunner)
	baton <- newRunner
}
