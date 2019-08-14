package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}
}

//使用 io.Reader 和 io.Writer 接口，写一个简单的 curl 程序
func main() {
	//url := "http://coderdaily.net/domain/getDomain"
	//r, err := http.Get(url)
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	//从 body 复制到 Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
