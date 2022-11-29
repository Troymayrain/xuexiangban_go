package main

import "fmt"

func main() {
	//定义数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	//定义切片
	var s1 []int
	fmt.Println(s1)

	if s1 == nil {
		fmt.Println("切片是空的")
	}

	s2 := []int{5, 6, 7, 8}
	fmt.Printf("%T,%T\n", arr, s2)
	fmt.Println(s2[0])

}
