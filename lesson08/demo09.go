package main

import "fmt"

// 定义别名
type DiyInt int

func main() {

	//var a int =10
	//var b DiyInt = 20
	//c:=(int)b+a
	//fmt.Println(c)
	//
	type myint = int
	var e int = 30
	var f myint = 40
	fmt.Println(e + f)

}
