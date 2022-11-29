package main

import (
	"fmt"
	"os"
)

func main() {

	fileInfo, er := os.Stat("./a.txt")
	fmt.Println(fileInfo.Mode())
	fmt.Println(er)

	//1、打开文件，建立连接
	file, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("文件名：", file.Name())

	//2、关闭连接
	defer file.Close()

	//读取a.txt
	s := make([]byte, 1, 1024)
	n, er1 := file.Read(s)
	fmt.Println(n)
	fmt.Println(string(s))
	fmt.Println(er1)

	n, er1 = file.Read(s)
	fmt.Println(n)
	fmt.Println(string(s))
	fmt.Println(er1)

	n, er1 = file.Read(s)
	fmt.Println(n)
	fmt.Println(string(s))
	fmt.Println(er1)

	n, er1 = file.Read(s)
	fmt.Println(n)
	fmt.Println(string(s))
	fmt.Println(er1)

	//for循环读取文件内容

	//file1, err1 := os.OpenFile("./a.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	//if err1 != nil {
	//	fmt.Println("err1:", err1)
	//	return
	//}
	//fmt.Println("文件名：", file1.Name())
	//
	//defer file1.Close()

}
