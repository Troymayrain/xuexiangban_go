package main

import "fmt"

func main() {
	//切片扩容
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2)
	fmt.Println(s1)

	s1 = append(s1, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s1)

	s2 := []int{10, 11, 12}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	//切片遍历
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}

	fmt.Println("================")

	for _, v := range s1 {
		fmt.Println(v)
	}

}
