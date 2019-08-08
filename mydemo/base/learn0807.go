package main

import (
	"fmt"
	"strings"
)

func main() {
	var result = strings.Compare("hello", "hello2")
	fmt.Println(result)

	fmt.Printf("hello %s", 12345)
}

/**
init() 先于 main() 被执行，在执行时会优先扫描所有的 init() 执行，然后才执行 main()
*/
func init() {
	fmt.Println("init start")
}

func init() {
	fmt.Println("init start 2", 2)
}
