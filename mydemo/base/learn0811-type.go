package main

import "fmt"

//go 语言的类型系统
func main() {

}

type user struct {
	name  string
	email string
	age   int
}

func init() {
	//简单声明一个 user 变量，初始化为零值
	var bill user
	fmt.Println(bill)//{  0}

	//声明一个 user 变量，并赋值
	lisa := user{
		name:  "lisa",
		email: "lisa@gmail.com",
		age:   28,
	}
	fmt.Println(lisa)//{lisa lisa@gmail.com 28}
}