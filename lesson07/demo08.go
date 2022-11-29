package main

import "fmt"

type Address struct {
	city, state string
}

type Person struct {
	name string
	age  int
	//结构体嵌套
	address Address
}

func main() {

	p := Person{}
	p.name = "mayu"
	p.age = 18
	p.address = Address{city: "武汉", state: "中国"}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.address)

}
