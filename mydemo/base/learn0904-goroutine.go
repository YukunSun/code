package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("创建 gogroutine:")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("End...")
}

//耗时函数，比如寻找素数
func printPrime(s string) {
	defer wg.Done()
next:
	for outer := 2; outer < 500000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d \n", s, outer)//会观察到，在 A、B 之间来回切换
	}
	fmt.Println("task completed：", s)
}
