package main

import "fmt"

func main() {
	//数组遍历
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1[0])
	fmt.Println(arr1[1])
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	fmt.Println(arr1[4])

	fmt.Println("===================")

	//fori
	for i := 0; i < 5; i++ {
		fmt.Println(arr1[i])
	}

	//arr1.for
	for index, value := range arr1 {
		fmt.Println(index, value)
	}

}
