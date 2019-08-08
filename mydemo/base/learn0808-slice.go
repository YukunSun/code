package main

import "fmt"

func main() {

}

//创建切片
func init() {
	//使用 make()，并初始化长度
	slice := make([]string, 5)
	fmt.Println(slice)

	//创建一个长度为 len=3，容量为 cap=5 的切片：表示该切片只能访问3个元素，而其底层数组是拥有5个元素的
	slice2 := make([]string, 3, 5)
	fmt.Println(slice2)
	//len>5 是不行的,
	//slice3 := make([]string, 5, 3)

	//使用字面量创建切片，如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice3)
}

//空切片、nil切片
func init() {
	//创建一个 nil 切片
	var slice []int;
	fmt.Println(slice) //[]

	//创建一个空切片
	slice2 := []int{}
	fmt.Println(slice2) //[]
	slice3 := make([]int, 0)
	fmt.Println(slice3) //[]
}

func init() {
	slice := []int{10, 20, 30, 40, 50}
	//len=2,cap=4，因为 newSlice 与 slice 共享一个底层数组
	//另，对于 slice[i,j]，len = j-i,cap = k-i
	newSlice := slice[1:3]
	fmt.Println(newSlice) //[20 30]

	//因为是共享底层数组，所以一旦修改，则会影响到别的值
	newSlice[1] = 2
	fmt.Println(slice) //[10 20 2 40 50]
}
