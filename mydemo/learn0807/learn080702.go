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
	fmt.Println("map test :")
}
