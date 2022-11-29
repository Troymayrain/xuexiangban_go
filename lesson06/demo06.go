package main

import "fmt"

func main() {
	//数组 值类型 拷贝
	arr := [4]int{1, 2, 3, 4}
	arr1 := arr
	fmt.Println(arr, arr1)
	arr1[0] = 100
	fmt.Println(arr, arr1)

	//切片 引用类型
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	fmt.Println(s1, s2)
	s2[1] = 12
	fmt.Println(s1, s2)

	fmt.Printf("%p,%p\n", s1, s2)
}
