package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//有缓冲通道

var wg sync.WaitGroup

const (
	numGoroutines = 4
	taskLoad      = 10 //缓冲区大小
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numGoroutines)

	//启动4个goroutine
	for gr := 0; gr < numGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//往通道添加任务
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task:%d", post)
	}

	//待所有任务都退出时，关闭通道
	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d shutting down!\n", worker)
			return
		}
		//开始工作
		fmt.Printf("Worker %d started %s \n", worker, task)

		//模拟工作时间
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Printf("Worker %d completed %s \n", worker, task)
	}
}

// 输出示例：
//Worker 3 started Task:3
//Worker 2 started Task:2
//Worker 1 started Task:4
//Worker 3 completed Task:3
//Worker 3 started Task:5
//Worker 0 started Task:1
//Worker 2 completed Task:2
//Worker 2 started Task:6
//Worker 2 completed Task:6
//Worker 2 started Task:7
//Worker 3 completed Task:5
//Worker 3 started Task:8
//Worker 1 completed Task:4
//Worker 1 started Task:9
//Worker 0 completed Task:1
//Worker 0 started Task:10
//Worker 3 completed Task:8
//Worker 3 shutting down!
//Worker 1 completed Task:9
//Worker 1 shutting down!
//Worker 2 completed Task:7
//Worker 2 shutting down!
//Worker 0 completed Task:10
//Worker 0 shutting down!
