package main

import "fmt"

type Dog struct {
	name string
	age  int
}

// 定义一个方法,函数加上调用类型就成为了方法
func (dog Dog) eat() {
	fmt.Println("DOG, ...eat")
}

func (dog Dog) sleep() {
	fmt.Println("DOG, ...sleep")
}

type Cat struct {
	name string
	age  int
}

// 定义一个方法
func (cat Cat) eat() {
	fmt.Println("CAT, ...eat")
}

func main() {
	dog := Dog{
		name: "旺财",
		age:  1,
	}
	dog.eat()
	dog.sleep()

	cat := Cat{
		name: "喵喵",
		age:  1,
	}
	cat.eat()
}
