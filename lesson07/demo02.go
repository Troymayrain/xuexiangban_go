package main

import "fmt"

func main() {
	//定义一个变量
	var a int = 10
	fmt.Println("a变量的值：", a)
	fmt.Println("a变量的内存地址：", &a)

	//定义一个指针变量
	var p *int
	p = &a
	fmt.Println("p变量存储的内存地址：", p)
	fmt.Println("p变量的内存地址：", &p)
	fmt.Println("p变量存储的内存地址的数据：", *p)

	//指针嵌套
	var ptr **int
	ptr = &p
	fmt.Println("ptr变量存储的内存地址：", ptr)
	fmt.Println("ptr变量的内存地址：", &ptr)
	fmt.Println("ptr变量存储的内存地址的内存地址的数据：", **ptr)

	*p = 20
	fmt.Println("a变量的值：", a)

	**ptr = 30
	fmt.Println("a变量的值：", a)

}
