package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	//声明一个共享资源，供多个 groutine 同时读写
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go increaseCounter(1)
	go increaseCounter(2)

	wg.Wait()
	fmt.Println("result :", counter) //result : 2
}

func increaseCounter(i int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := i
		//退出当前 goroutine，给其他 goroutine 执行的机会。类比 java 的 Thread.join()
		runtime.Gosched()
		value++
		counter = value
	}
}
