package test

import "fmt"

import _ "xuexiangban_go/lesson10/test2"

// 匿名导包时会先于main函数执行
func init() {
	fmt.Println("test--init--a1")
}
func init() {
	fmt.Println("test--init--a2")
}
func init() {
	fmt.Println("test--init--a3")
}
