package main

import "fmt"

func main() {
	//切片扩容是成倍增长的
	s1 := []int{1, 2, 3}
	fmt.Printf("len:%d,cap:%d,s1内容：%v, 地址：%p\n", len(s1), cap(s1), s1, s1)

	s1 = append(s1, 4, 5)
	fmt.Printf("len:%d,cap:%d,s1内容：%v, 地址：%p\n", len(s1), cap(s1), s1, s1)

	s1 = append(s1, 6, 7)
	fmt.Printf("len:%d,cap:%d,s1内容：%v, 地址：%p\n", len(s1), cap(s1), s1, s1)

	s1 = append(s1, 8, 9, 10)
	fmt.Printf("len:%d,cap:%d,s1内容：%v, 地址：%p\n", len(s1), cap(s1), s1, s1)

	s1 = append(s1, 11, 12, 13)
	fmt.Printf("len:%d,cap:%d,s1内容：%v, 地址：%p\n", len(s1), cap(s1), s1, s1)
}
