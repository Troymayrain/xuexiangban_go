package main

import "fmt"

type Student struct {
	name string
	age  int
}

// 匿名字段
type Teacher struct {
	string
	int
}

func main() {
	s1 := Student{name: "may", age: 23}
	fmt.Println(s1)

	//匿名结构体
	s2 := struct {
		name string
		age  int
	}{name: "李四", age: 12}
	fmt.Println(s2)

	//匿名字段
	t1 := Teacher{"王老师", 45}
	fmt.Println(t1)
	fmt.Println(t1.int)
	fmt.Println(t1.string)

}
