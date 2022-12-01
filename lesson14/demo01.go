package main

import "fmt"

// 泛型的使用，泛型就是类型形参
func main() {

	strs := []string{"张三三", "李四", "王五", "赵六"}
	is := []int{1, 2, 3, 4}
	fs := []float64{1.23, 2.31, 3.13}

	printSlice(strs)
	printSlice(is)
	printSlice(fs)

}

// 定义形参为泛型类型，并规定了泛型的类型
//
//	func printSlice[T string | int | float64](s []T) {
//		for _, v := range s {
//			fmt.Println(v)
//		}
//	}

// any:可以代表go里面所有内置的基本数据类型
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

//comparable:代表go里面内置的所有可比较的数据类型
//func printSlice[T comparable](s []T) {
//	for _, v := range s {
//		fmt.Println(v)
//	}
//}

//func printSlice[T interface{}](s []T) {
//	for _, v := range s {
//		fmt.Println(v)
//	}
//}
