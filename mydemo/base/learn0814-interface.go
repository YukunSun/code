package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

}

//例：将 bytes.Buffer 用于 io.Copy
func init() {
	var b bytes.Buffer
	//将字符串写入 buffer
	b.Write([]byte("hello"))
	//使用Fprintf 将字符串拼接到Buffer
	fmt.Fprintf(&b, "world")
	//将Buffer 的内容写到Stdout
	io.Copy(os.Stdout, &b) //helloworld
}

//例2
type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

//定义一个具有通知行为的接口
type notifier interface {
	notify()
}

//定义接口的实现
func (u *user) notify() {
	fmt.Printf("sending email %s<%s> \n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("sending admin email %s<%s> \n", a.name, a.email)
}

//sendNotification 接受一个实现了notifier 接口的值并发送通知
//任何一个实现了 notifier 接口的值都可以传入sendNotification 函数
func sendNotification(n notifier) {
	n.notify()
}

//例
func init() {
	u := user{"Bill", "bill@gmail.com"}
	sendNotification(&u) //sending email Bill<bill@gmail.com>

	//多态
	admin := admin{"Lisa", "lisa@gmail.com"}
	sendNotification(&admin) //sending admin email Lisa<lisa@gmail.com>
}
