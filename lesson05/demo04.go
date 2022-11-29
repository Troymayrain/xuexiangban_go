package main

import "fmt"

func main() {
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [3]int{1, 2, 3}
	arr2 := [2]string{"mayu", "hello"}
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)

	fmt.Println("=============")

	//基本数据类型值传递
	num2 := num
	fmt.Println(num, num2)
	num2 = 21
	fmt.Println(num, num2)

	//数组也是值传递
	arr3 := arr1
	fmt.Println(arr1, arr3)
	arr3[2] = 233
	fmt.Println(arr1, arr3)

}
