package main

import "fmt"

func main() {

}

//二维数组
func init() {
	var array [4][2]int
	array = [4][2]int{{1, 1}, {2, 1}, {3, 1}, {4, 1}}
	fmt.Println(array)

	array2 := [4][2]int{1: {12, 11}, 3: {13, 12}}
	array2[0][0] = 1
	array2[0][1] = 2
	fmt.Println(array2) //[[1 2] [12 11] [0 0] [13 12]]
}
