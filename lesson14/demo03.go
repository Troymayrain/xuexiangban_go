package main

import "fmt"

type MySlice[T int | float64] []T

// 给泛型添加方法
func (m MySlice[T]) Sum() T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

//func Add(a int, b int) int {
//	return a+b
//}

// 添加泛型函数
func Add[T int | string | float64](a T, b T) T {
	return a + b
}

// 泛型方法的使用
func main() {

	//泛型方法
	var mys1 MySlice[int] = []int{1, 2, 3, 4}
	fmt.Println(mys1.Sum())
	var mys2 MySlice[float64] = []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(mys2.Sum())

	//泛型函数
	fmt.Println(Add[int](1, 2))
	fmt.Println(Add(1, 2))
	fmt.Println(Add[string]("1", "2"))
	fmt.Println(Add("1", "2"))
	fmt.Println(Add[float64](1.1, 2.2))
	fmt.Println(Add(1.1, 2.2))

}
