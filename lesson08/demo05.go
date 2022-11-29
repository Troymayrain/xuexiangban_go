package main

import "fmt"

//接口实现多态
type Animal interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println(dog.name, "在吃")
}
func (dog Dog) sleep() {
	fmt.Println(dog.name, "在睡")
}

type Cat struct {
	name string
}

func (cat Cat) eat() {
	fmt.Println(cat.name, "在吃")
}
func (cat Cat) sleep() {
	fmt.Println(cat.name, "在睡")
}

func test1(a Animal) {
	fmt.Println("test1")
	a.eat()
}

func main() {

	dog := Dog{
		name: "大黄",
	}
	fmt.Println(dog.name)
	dog.eat()
	dog.sleep()

	test1(dog)

	var animal Animal
	animal = dog
	animal.eat()

	cat := Cat{name: "大喵"}
	animal = cat
	animal.eat()

}
