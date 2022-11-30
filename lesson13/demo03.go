package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改变量的值
func main() {
	var num float64 = 1.23
	fmt.Println(num)
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println(newValue.Type())
	fmt.Println(newValue.Kind())
	fmt.Println(newValue.CanSet())

	if newValue.CanSet() {
		newValue.SetFloat(3.14)
	}

	fmt.Println(num)
}
