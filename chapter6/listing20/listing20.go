package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//计数器设置为协程数2
	wg.Add(2)

	//创建一个无缓冲通到
	court := make(chan int)

	go player("运动员A", court)
	go player("运动员B", court)

	fmt.Println("开球......")
	//往无缓冲通到存一个值
	court <- 0

	wg.Wait()
	fmt.Printf("比赛结束")

}

func player(name string, court chan int) {

	defer wg.Done()
	for {
		//等待缓冲通道的值
		ball, ok := <-court
		//没收到则说明某个协程将通道关闭了(即对方没有击中球)
		if !ok {
			fmt.Println("对手击球失败", name, "获胜")
			break
		}

		//模拟运动员击球，如果能被13整除则说明该运动员击球失败，关闭通道
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println(name, "未能击中球")
			close(court)
			break
		}

		//增加击球数并写入通道中，模拟击球给对方
		ball++

		fmt.Println(name, "击中球:", ball, "次")
		court <- ball
	}
}
