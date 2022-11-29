package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func main() {
	//结构体值传递类型
	user1 := User{"张三", 23, "男"}
	fmt.Println(user1)
	fmt.Printf("%T\n", user1)
	user2 := user1
	user2.name = "李四"
	fmt.Println(user1)

	//结构体指针
	var p *User
	p = &user1
	fmt.Printf("p的类型：%T\n", p)
	fmt.Println(p)
	fmt.Println(*p)
	p.name = "王五"
	fmt.Println(user1)

	fmt.Println("==========")

	p1 := new(User)
	p1 = &user1
	fmt.Printf("p1的类型：%T\n", p1)
	fmt.Println(p1)
	fmt.Println(*p1)
	p1.age = 233
	fmt.Println(user1)

}
