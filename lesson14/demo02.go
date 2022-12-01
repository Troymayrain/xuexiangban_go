package main

import "fmt"

// 定义泛型类型
func main() {

	//自定义切片类型
	type Slice[T int | float32 | float64] []T

	var a Slice[int] = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Type name:%T\n", a)

	var b Slice[float32] = []float32{1, 2, 3}
	fmt.Println(b)
	fmt.Printf("Type name:%T\n", b)

	var c Slice[float64] = []float64{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("Type name:%T\n", c)

	//var d Slice[string] = []string{"1", "2", "3"}

	//自定义map类型
	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	var map1 MyMap[string, float32] = map[string]float32{
		"go":   9.9,
		"java": 9.0,
	}
	fmt.Println(map1)

}
