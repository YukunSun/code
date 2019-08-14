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
type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("sending email %s<%s>", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

//例
func init() {
	u := user{"Bill", "bill@gmail.com"}
	sendNotification(u)
}
