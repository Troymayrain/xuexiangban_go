package main

import (
	"errors"
	"fmt"
)

func main() {
	//自定义错误信息
	errorInfo := errors.New("我是一个错误信息")
	fmt.Println(errorInfo)
	fmt.Printf("%T\n", errorInfo)

	errorInfo1 := setAge(-1)
	if errorInfo1 != nil {
		fmt.Println(errorInfo1)
	}

	errorInfo2 := fmt.Errorf("我是一个错误%d", 500)
	fmt.Println(errorInfo2)
	fmt.Printf("%T\n", errorInfo2)

}

func setAge(age int) error {
	if age < 0 {
		err := errors.New("年龄输入不合法")
		return err
	}
	fmt.Println(age)
	return nil
}
