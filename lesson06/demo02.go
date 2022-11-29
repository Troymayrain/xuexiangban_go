package main

import "fmt"

func main() {
	//make函数定义切片
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	fmt.Println("长度：", len(s1))
	fmt.Println("容量：", cap(s1))

	s1[1] = 100
	fmt.Println(s1)
	//s1[6] = 120
	//fmt.Println(s1)

}
