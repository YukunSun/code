package main

import "fmt"

func main() {
	//无缓冲通道
	/*	unbuffered := make(chan string)
		unbuffered <- "hello world 1"
		value := <-unbuffered
		fmt.Println(value)
	*/
	//有缓冲通道，chan:关键字,string:表示通道需要交换数据的类型，10：表示缓冲区大小
	buffered := make(chan string, 10)
	//向通道发送值或者指针需要用到<-操作符
	buffered <- "hello world 2"
	//从通道里读取数据
	value2 := <-buffered
	fmt.Println(value2)
}
