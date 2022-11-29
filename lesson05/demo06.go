package main

import "fmt"

func main() {

	//定义二维数组
	arr := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 23},
	}

	fmt.Println(arr[0][0])
	fmt.Println(arr[1][2])

	//fori遍历二维数组
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(arr[i][j])
		}
	}

	//arr.for遍历二维数组
	for i, v := range arr {
		fmt.Println(i, v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

}
