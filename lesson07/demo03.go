package main

import "fmt"

func main() {
	//数组指针
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Printf("arr的内存地址：%p\n", &arr)
	var p1 *[4]int
	p1 = &arr
	fmt.Printf("p1变量指向的内存地址：%p\n", p1)
	fmt.Printf("p1变量的内存地址：%p\n", &p1)
	fmt.Println("p1变量指向的内存地址的数据：", *p1)

	(*p1)[0] = 100
	fmt.Println(arr)
	fmt.Println("p1变量指向的内存地址的数据：", *p1)

	//简写
	p1[0] = 200
	fmt.Println(arr)
	fmt.Println("p1变量指向的内存地址的数据：", *p1)

	//指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr1 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr1)

	a = 10
	fmt.Println(arr1)
	fmt.Println(*arr1[0])

}
