package main

import "fmt"

func main() {

}

/**
数组
*/
func init() {
	fmt.Println("array test :")
	array := [5]int{1, 2, 4, 5, 7}
	//可以修改其对应的值
	array[2] = 80
	fmt.Println(array)

	//数组大小可以不指定
	array2 := [...]int{1, 2, 4, 5, 7}
	fmt.Println(array2)

	//声明一个包含5个元素的指向 int 的数组,用整型指针初始化索引为0 和1 的数组元素
	array3 := [5]*int{0: new(int), 1: new(int)}
	//赋值
	*array3[0] = 10
	*array3[1] = 20

	fmt.Println(array3) //输出示例：[0xc00008a010 0xc00008a018 <nil> <nil> <nil>]
	fmt.Println(*array3[1])
}

func init() {
	//赋值操作
	var array1 [5]string
	array2 := [5]string{"red", "blue", "yellow", "green", "pink"}
	array1 = array2

	fmt.Println(array1)
	fmt.Println(array2)

	//赋值操作2
	var array3 [3]*string
	array4 := [3]*string{new(string), new(string), new(string)}

	*array4[0] = "a"
	*array4[1] = "b"
	*array4[2] = "c"

	array3 = array4

	fmt.Println(*array3[0])
	fmt.Println(*array4[1])
}