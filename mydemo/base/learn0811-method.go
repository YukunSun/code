package main

import "fmt"

type user struct {
	email string
	name  string
}

func main() {

}

//1. 使用值接收者声明一个方法
//关键字func 和函数名之间的 参数被称作接收者，将函数与接收者的类型绑在一起。
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s> \n", u.name, u.email)
}

//2.使用指针接收者来声明方法；这个方法使用指针接收者声明。这个接收者的类型是指向 user 类型值的指针，而不是 user 类型的值。当调用使用指针接收者声明的方法时，这个方法会共享调用方法时接收者所指向的值
func (u *user) changeEmail(email string) {
	u.email = email
}

func init() {
	//1. user 类型的值可以用来直接调用方法
	bill := user{"bill@gmail.com", "Bill"}
	bill.notify() //这个语法与调用一个包里的函数看起来很类似。但在这个例子里，bill 不是包名，而是变量名。

	//2. 指向user 类型值的指针也可以用来调用这个方法
	lisa := &user{"lisa@gmail.com", "Lisa"}
	lisa.notify()
}

func init() {
	lisa := &user{"lisa@gmail.com", "Lisa"}
	lisa.changeEmail("lisa-new@gmail.com")
	lisa.notify()
}