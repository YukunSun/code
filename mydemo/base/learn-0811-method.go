package main

import "fmt"

type user struct {
	email string
	name  string
}

//接收参数 user;关键字func 和函数名之间的 参数被称作接收者，将函数与接收者的类型绑在一起。
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s> \n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

/**
例子
*/
func main() {
	//1. user 类型的值可以用来调用使用 值接收者声明的方法
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	//2. 指向user 类型值的指针也可以用来调用 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

}
