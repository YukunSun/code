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

func init() {
	fmt.Println("map test :")
}
