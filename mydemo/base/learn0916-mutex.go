package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Printf("result counter:%d \n", counter)
}

func incCounter() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()

		value := counter
		//退出当前 goroutine，给其他 goroutine 执行的机会。类比 java 的 Thread.join()
		runtime.Gosched()
		value++
		counter = value

		mutex.Unlock()
	}
}
