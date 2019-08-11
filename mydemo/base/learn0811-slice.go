package main

import "fmt"

func main() {

}

//强制在 append 时创建一个新的切片
func init() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	//限制新的切片大小为3，以免进行 append 操作时，修改了原来的切片内容
	newSlice := slice[1:3:3]
	newSlice = append(newSlice, 7, 9, 10, 12, 14)
	fmt.Println(newSlice)
	fmt.Println(slice)
}

//将两个切片追加在一起
func init() {
	slice := []int{1, 2, 3}
	slice2 := []int{4, 5, 2}

	fmt.Printf("%v", append(slice, slice2...)) //[1 2 3 4 5 2]
}

//使用 range 或普通的for循环迭代切片
func init() {
	slice := []int{1, 2, 3}

	fmt.Println("使用 for 迭代：")
	for i := 1; i < len(slice); i++ {
		fmt.Printf("index:%d,value:%d \n", i, slice[i])
	}

	//该 range 返回了两个值
	fmt.Println("使用 range 迭代：")
	for index, value := range slice {
		fmt.Printf("index:%d,value:%d \n", index, value)
	}
	//使用空白字符省略掉索引
	for _, value := range slice {
		fmt.Printf("value:%d \n", value)
	}

	//但是 range 创建了每个元素的副本，而不是直接返回对该元素的引用，可以输出地址验证一下
	//因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以value 的地址总是相同的。要想获取每个元素的地址，可以使用切片变量和索引值，而不是使用该值变量的地址作为指向每个元素的指针
	for index, value := range slice {
		fmt.Printf("idx:%d,value address:%X, element Address:%X \n", index, &value, &slice[index])
	}
}

//在函数间传递切片
func init() {
	slice := make([]int, 1e6) //分配1000000个整型值的切片

	//传递时复制的是：指针字段、长度、容量，而不是整个底层数组
	slice2 := foo(slice)
	fmt.Println(slice[0])
	fmt.Println(slice2[0])
}

func foo(values []int) []int {
	values[0] = 1
	return values
}
