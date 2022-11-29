package main

import "fmt"

func main() {
	//指针函数
	p := f1()
	fmt.Printf("p->的内存地址：%p\n", p)
	fmt.Println("p的内存地址：", &p)
	fmt.Println("p->的内存地址中的数据：", *p)

	(*p)[0] = 100
	fmt.Println("p->的内存地址中的数据：", *p)

	p[0] = 200
	fmt.Println("p->的内存地址中的数据：", *p)

	//指针参数
	a := 10
	fmt.Println(a)
	f2(&a)
	fmt.Println(a)
}

// 指针函数
func f1() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}

// 指针作为参数
func f2(ptr *int) {
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = 20
}
