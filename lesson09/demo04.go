package main

import "fmt"

// panic使用
func main() {
	defer fmt.Println("main-----1")
	defer fmt.Println("main-----2")
	fmt.Println("main-----3")
	test(1)
	defer fmt.Println("main-----4")
}

func test(num int) {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println("msg:", msg)
		}
	}()
	defer fmt.Println("test-----1")
	defer fmt.Println("test-----2")
	fmt.Println("test-----3")
	if num == 1 {
		panic("num不能为1")
	}
	defer fmt.Println("test-----4")
}
