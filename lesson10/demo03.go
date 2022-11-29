package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "true"
	//将字符串转为bool类型
	b1, _ := strconv.ParseBool(str)
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)

	str1 := strconv.FormatBool(b1)
	fmt.Println(str1)
	fmt.Printf("%T\n", str1)

	//字符串转为int类型
	s := "100"
	s1, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	formatInt := strconv.FormatInt(s1, 10)
	fmt.Println(formatInt)
	fmt.Printf("%T\n", formatInt)

	//十进制数与字符串之间的快速转换
	atoi, _ := strconv.Atoi("23")
	fmt.Println(atoi)
	fmt.Printf("%T\n", atoi)

	itoa := strconv.Itoa(123)
	fmt.Println(itoa)
	fmt.Printf("%T\n", itoa)

}
