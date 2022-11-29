package main

import "fmt"

func main() {

	var a int = 10
	fmt.Println("变量a的值：", a)
	fmt.Println("变量a的内存地址：", &a)
	//指针变量
	b := &a
	fmt.Println("变量b的值为：", b)
	fmt.Println("变量b指向的内存地址的数据为：", *b)

	*b = 30
	fmt.Println("变量a的值：", a)

}
