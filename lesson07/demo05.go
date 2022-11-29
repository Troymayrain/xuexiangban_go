package main

import "fmt"

// 定义一个结构体
// 大写 公开，小写 私有
type User struct {
	name string
	age  int
	sex  string
}

func main() {

	//创建对象
	var user1 User
	fmt.Println(user1)
	user1.name = "mmm"
	user1.age = 12
	user1.sex = "男"
	fmt.Println(user1)

	user2 := User{}
	user2.name = "uzi"
	user2.age = 23
	user2.sex = "男"
	fmt.Println(user2)

	user3 := User{
		name: "小明",
		age:  10,
		sex:  "男",
	}
	fmt.Println(user3)

	user4 := User{"张三", 43, "男"}
	fmt.Println(user4)
	fmt.Println(user4.name)
	fmt.Println(user4.age)

}
