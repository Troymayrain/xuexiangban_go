package main

import "fmt"

type A interface {
	test1()
}
type B interface {
	test2()
}

// 通过接口继承实现了接口嵌套
type C interface {
	A
	B
	test3()
}

type Dogg struct {
}

func (dog Dogg) test3() {
	fmt.Println("test3")
}
func (dog Dogg) test2() {
	fmt.Println("test2")
}
func (dog Dogg) test1() {
	fmt.Println("test1")
}

func main() {

	dog := Dogg{}
	dog.test1()
	dog.test2()
	dog.test3()

	var ia A = dog
	ia.test1()

	var ib B = dog
	ib.test2()

	var ic C = dog
	ic.test1()
	ic.test2()
	ic.test3()

}
