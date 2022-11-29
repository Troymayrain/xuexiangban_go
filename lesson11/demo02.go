package main

import (
	"fmt"
	"os"
)

func main() {

	//目录存在，就报错，反之可以创建
	err := os.Mkdir("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\test", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		//return
	}
	fmt.Println("文件夹创建成功")

	//创建层级目录
	err1 := os.MkdirAll("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\test2\\a\\b\\c", os.ModePerm)
	if err1 != nil {
		fmt.Println("err1：", err1)
		//return
	}

	//删除层级目录
	//err2 := os.Remove("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\test")
	err2 := os.RemoveAll("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\test2")
	if err2 != nil {
		fmt.Println("err1：", err2)
		return
	}

	//创建文件
	//file, err3 := os.Create("./b.txt")
	//fmt.Println(file)
	//fmt.Println(file.Name())
	//if err3 != nil {
	//	fmt.Println("err3：", err3)
	//	return
	//}

	//删除文件
	err4 := os.Remove("./b.txt")
	fmt.Println(err4)
}
