package main

import "fmt"

func main() {
	//值类型：基本数据类型，array，struct 都是深拷贝；引用类型：slice,map都是浅拷贝
	//切片深拷贝实现方式一
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 0, 0)

	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1, s2)
	s2[0] = 100
	fmt.Println(s1, s2)

	//切片深拷贝实现方式二
	s3 := []int{6, 7, 8}
	copy(s3, s1)
	fmt.Println(s1, s3)

	copy(s1, s3)
	fmt.Println(s1, s3)

}
