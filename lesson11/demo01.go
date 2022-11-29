package main

import (
	"fmt"
	"os"
)

func main() {
	//fileInfo, err := os.Stat("./a.txt")
	fileInfo, err := os.Stat("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\test")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	//通过反射获取文件的更加详细信息
	fmt.Println(fileInfo.Sys())
}
