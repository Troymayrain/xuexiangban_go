package main

import (
	"./pojo"
	"fmt"
)

func main() {

	var user1 pojo.User
	user1.Name = "王五"
	user1.Money = "123"
	user1.Age = 23
	fmt.Println(user1)

	//var cat pojo.Cat
	//cat.age

}
