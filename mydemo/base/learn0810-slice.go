package main

import "fmt"

func main() {

}

func init() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	//限制新的切片大小为3，以免进行 append 操作时，修改了原来的切片内容
	newSlice := slice[1:3:3]
	newSlice = append(newSlice, 7, 9, 10, 12, 14)
	fmt.Println(newSlice)
	fmt.Println(slice)
}
