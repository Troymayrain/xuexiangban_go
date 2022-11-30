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

func (user User) Say(msg string) {
	fmt.Println(msg)
}

func (user User) PrintInfo() {
	fmt.Printf("姓名：%s,年龄：%d,性别：%s", user.Name, user.Age, user.Sex)
}

// 利用反射获取结构体的属性和方法
func main() {

	user := User{
		Name: "张三",
		Age:  10,
		Sex:  "男",
	}

	reflectGetInfo(user)

}

func reflectGetInfo(inter interface{}) {
	getType := reflect.TypeOf(inter)
	fmt.Println("getType name is ", getType.Name())
	fmt.Println("getType kind is ", getType.Kind())

	getValue := reflect.ValueOf(inter)
	fmt.Println("getValue is", getValue)

	fmt.Println("属性数量：", getType.NumField())
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名：%s,字段类型：%s,字段值：%v\n", field.Name, field.Type, value)
	}

	fmt.Println("方法数量", getType.NumMethod())
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s,方法类型：%v\n", method.Name, method.Type)
	}
}
