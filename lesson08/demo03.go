package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println(animal.name + "正在吃...")
}

func (animal Animal) sleep() {
	fmt.Println(animal.name + "在睡觉...")
}

type Dog struct {
	Animal
}

// 重写父类方法
func (dog Dog) eat() {
	fmt.Println(dog.name + "正在吃-dog")
}

type Cat struct {
	Animal
	Color string
}

// 重写父类方法
func (cat Cat) eat() {
	fmt.Println(cat.name + "正在吃-cat")
}

func (cat Cat) miao() {
	fmt.Println("喵喵喵...")
}

func main() {
	animal := Animal{
		name: "动物",
		age:  1,
	}
	fmt.Println(animal.name)
	animal.eat()

	dog := Dog{
		Animal: Animal{name: "旺财", age: 2},
	}
	fmt.Println(dog.name)
	dog.eat()

	cat := Cat{
		Animal: Animal{name: "喵喵", age: 3},
		Color:  "red",
	}
	fmt.Println(cat.Color)
	cat.eat()
	cat.miao()
}
