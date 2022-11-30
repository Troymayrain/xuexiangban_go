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

func (user User) PrintInfo() {
	fmt.Printf("姓名：%s,年龄：%d,性别：%s\n", user.Name, user.Age, user.Sex)
}

func (user User) Say(msg string) {
	fmt.Println("User say:", msg)
}

// 通过反射调用结构体的方法
func main() {
	user := User{"李四", 23, "男"}
	value := reflect.ValueOf(user)
	value.MethodByName("PrintInfo").Call(nil)
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf("反射来调用的")
	value.MethodByName("Say").Call(args)
}
