package main

import (
	"fmt"
	"reflect"
)

// 通过反射调用函数
func main() {
	v1 := reflect.ValueOf(fun1)
	v1.Call(nil)

	v2 := reflect.ValueOf(fun2)
	args2 := make([]reflect.Value, 2)
	args2[0] = reflect.ValueOf(1)
	args2[1] = reflect.ValueOf("dfsfsdfsd")
	v2.Call(args2)

	v3 := reflect.ValueOf(fun3)
	args3 := make([]reflect.Value, 2)
	args3[0] = reflect.ValueOf(2)
	args3[1] = reflect.ValueOf("hello")
	resultVal := v3.Call(args3)
	fmt.Println(resultVal[0].Interface())
}

func fun1() {
	fmt.Println("func1")
}

func fun2(i int, s string) {
	fmt.Println("i:", i, "s:", s)
}

func fun3(i int, s string) string {
	fmt.Println("i:", i, "s:", s)
	return s
}
