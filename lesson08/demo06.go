package main

import "fmt"

// 空接口的使用，空接口相当于java中的Object
type A interface {
}

type Dog struct {
	name string
}

type Cat struct {
	name string
	age  int
}

func test2(a A) {
	fmt.Println("test2")
	fmt.Println(a)
}

func test3(a interface{}) {
	fmt.Println("test2")
	fmt.Println(a)
}

func main() {
	var a1 A = Dog{name: "小黄"}
	var a2 A = Cat{name: "小喵", age: 1}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	test2(a1)
	test2(a2)
	test2(a3)
	test2(a4)

	test3(a1)
	test3(a2)
	test3(a3)
	test3(a4)

	//创建一个key string value interface{}的map
	map1 := make(map[string]interface{})
	map1["name"] = "haha"
	map1["age"] = 12
	map1["dog"] = Dog{name: "田园犬"}
	fmt.Println(map1)

	//创建一个interface{}的切片
	s := make([]interface{}, 0, 10)
	s = append(s, "dsf")
	s = append(s, 213)
	s = append(s, Cat{
		name: "橘猫",
		age:  12,
	})
	fmt.Println(s)
}
