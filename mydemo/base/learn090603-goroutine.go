package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//演示 LoadInt64/StoreInt64 原子函数
var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	//表示要等待2个goroutine
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	//准备停止程序
	fmt.Println("prepare to terminal...")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Println("Doing ", name)
		time.Sleep(500 * time.Microsecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("shutdown ", name)
			break
		}
	}
}
