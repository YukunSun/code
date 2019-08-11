package main

import "fmt"

func main() {

}

//map 定义
func init() {
	//1. 创建一个 map：key 为 string 类型，值为 int 类型
	dict := make(map[string]int)
	dict["k1"] = 1
	dict["k2"] = 2
	fmt.Println(dict)
	value, exists := dict["k1"];
	if exists {
		fmt.Println("value:", value) //value: 1
	}
	value3, exists3 := dict["k3"];                          //会初始化为0值
	fmt.Printf("value: %d,isExists:%t \n", value3, exists3) //value: 0,isExists:false

	//2.
	dict2 := map[string]string{"stu1": "sun", "stu2": "yu"}
	fmt.Println(dict2)

	//3. 定义一个nil map，该map无法赋值
	var dict3 map[string]string
	//dict3["k1"] = "hello"
	fmt.Println(dict3)
}

//迭代 map
func init() {
	dict := map[string]string{
		"stu1": "sun",
		"stu2": "yu",
		"stu3": "kun", //为什么允许最后一个逗号
	}
	//删除指定的 key
	delete(dict, "stu2")
	for key, value := range dict {
		fmt.Printf("key:%s,value:%s \n", key, value)
	}
}

func init() {
	dict := map[string]string{
		"stu1": "sun",
		"stu2": "yu",
		"stu3": "kun", //为什么允许最后一个逗号
	}
	removeKey(dict, "stu2")
	fmt.Println(dict)
}

func removeKey(dict map[string]string, key string) {
	delete(dict, key)
}
