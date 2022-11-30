package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

// 通过反射修改结构体的属性值
func main() {

	user := User{"mayu", 12, "男"}
	fmt.Println(user)

	userPtr := reflect.ValueOf(&user)
	if userPtr.Kind() == reflect.Ptr {
		value := userPtr.Elem()
		if value.CanSet() {
			value.FieldByName("Name").SetString("马雨于")
			value.FieldByName("Age").SetInt(18)
			value.FieldByName("Sex").SetString("男")
		}
	}
	fmt.Println(user)

}
