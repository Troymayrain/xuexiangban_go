package main

import "fmt"

/**
1、继承：is-a关系，通过定义匿名变量来实现 子类中可以直接.父类中的属性
2、聚合：has-a关系，通过定义带名字的变量来实现 子类中必须.变量名.父类中的属性
*/

// 定义一个父"类" 结构体
type Person struct {
	name string
	age  int
}

// 定义一个子"类" 结构体
type Student struct {
	Person //匿名变量 继承
	school string
}

func main() {

	p := Person{"mayu", 18}
	fmt.Println(p.name)

	s := Student{Person{"小明", 23}, "北大"}
	fmt.Println(s.Person.name)
	fmt.Println(s.school)

	//结构体中定义了匿名变量，匿名变量中的所有字段都是提升字段，子类中可以直接使用
	var s1 Student
	s1.name = "小红"
	s1.age = 34
	s1.school = "人大"
	fmt.Println(s1)

	s2 := Student{
		Person: Person{"小张", 35},
		school: "北邮",
	}
	fmt.Println(s2.name)
	fmt.Println(s2.age)

}
