package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//无缓冲通道，模拟网球比赛 🎾

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建球场，模拟无缓冲通道
	court := make(chan int)
	//表示等待两个 goroutine
	wg.Add(2)
	//两名选手
	go player("p1", court)
	go player("p2", court)

	//传输数据，开始发球了
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won!\n", name)
			return
		}

		//随机输球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed!\n", name)

			//关闭通道，表示输了
			close(court)
			return
		}

		fmt.Printf("player %s Hit %d \n", name, ball)
		ball++

		//击球
		court <- ball
	}
}

//输出示例：
// player p2 Hit 1
//player p1 Hit 2
//player p2 missed!
//player p1 won!
