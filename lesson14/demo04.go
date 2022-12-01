package main

import "fmt"

type int8AAA int8

// 自定义泛型约束，“~”代表可以兼容延伸类型，比如别名
type MyInt interface {
	int | ~int8 | int16 | int32 | int64
}

func GetMaxNum[T MyInt](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(GetMaxNum(a, b))
	fmt.Println(GetMaxNum[int](a, b))

	var c int8AAA = 10
	var d int8AAA = 20
	fmt.Println(GetMaxNum(c, d))
	fmt.Println(GetMaxNum[int8AAA](c, d))
}
