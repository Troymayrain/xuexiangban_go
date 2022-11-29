package main

import "fmt"

func main() {
	//数组常规初始化
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr1[2])
	//数组快速初始化
	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr2)
	//不确定长度的数组初始化
	arr3 := [...]string{"a", "b", "c", "dsdsf"}
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	fmt.Println(cap(arr3))

	//未指定下标的用默认值
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println(arr4)

}
