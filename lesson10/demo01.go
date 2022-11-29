package main

import (
	"fmt"
	_ "xuexiangban_go/lesson10/test"
)

// init函数没有入参，没有返回值，也无法被引用
func init() {
	fmt.Println("main--init--a1")
}

func main() {
	fmt.Println("sds")
}
