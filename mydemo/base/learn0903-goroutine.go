package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//只允许调度一个逻辑 CPU
	runtime.GOMAXPROCS(1);
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start gorountines...")

	//创建一个 goroutine
	go func() {
		//在函数退出时调用 Done 来通知 main() 函数工作已经完成
		defer wg.Done()
		for count := 0; count < 3; count++ {
			//这里的 char，只是变量名而已，而非像 java 里的关键字
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			//这里的 char，只是变量名而已，而非像 java 里的关键字
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish...")
	wg.Wait()
	fmt.Println("end...")
}
