package main
//演示 AddInt64 等原子函数
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	//声明一个共享资源，供多个 goroutine 同时读写
	counter2 int64
	wg2      sync.WaitGroup
)

func main() {
	wg2.Add(2)
	go increaseCounterSafe(1)
	go increaseCounterSafe(2)
	wg2.Wait()
	fmt.Println("result :", counter2) //result : 4
}

func increaseCounterSafe(i int) {
	defer wg2.Done()
	for i := 0; i < 2; i++ {
		//使用原子安全函数
		atomic.AddInt64(&counter2, 1)
		runtime.Gosched()
	}
}
