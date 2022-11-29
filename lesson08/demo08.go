package main

import "fmt"

// 接口断言
func main() {
	assertString(23)

	assertInt(12)
	assertInt("12")

	test("dsf")
	test(123)
	var i I
	test(i)
	var s string
	test(s)
}

func assertString(i interface{}) {
	s := i.(string)
	fmt.Println(s)
}

func assertInt(i interface{}) {
	i, ok := i.(int)
	if ok {
		fmt.Println("变量i是int类型的")
	} else {
		fmt.Println("变量i不是int类型的")
	}
}

type I interface {
}

func test(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case I:
		fmt.Println("I")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("未知类型")
	}
}
