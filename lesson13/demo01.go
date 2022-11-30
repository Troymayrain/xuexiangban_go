package main

import (
	"fmt"
	"reflect"
)

// 反射的使用
func main() {

	var a float64 = 3.14

	//通过反射可以获取变量的类型和值
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

	value := reflect.ValueOf(a)
	if value.Kind() == reflect.Float64 {
		fmt.Println("value.kind() is float64")
	}
	fmt.Println(value.Type())
	fmt.Println(value.Float())

}
