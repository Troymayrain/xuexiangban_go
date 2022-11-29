package main

import "fmt"

// 自定义错误
type MyDiyError struct {
	msg  string
	code int
}

func (e MyDiyError) Error() string {
	return fmt.Sprint("错误信息：", e.msg, "状态码：", e.code)
}

func test1(i int) (int, error) {
	if i != 0 {
		return i, &MyDiyError{"非零数据", 500}
	}
	return i, nil
}

func main() {

	i, err := test1(0)
	if err != nil {
		fmt.Println(err)
		myDiyError, ok := err.(*MyDiyError)
		if ok {
			fmt.Println(myDiyError.code)
			fmt.Println(myDiyError.msg)
		}
	}
	fmt.Println(i)

}
