package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件，建立连接
	file, err := os.OpenFile("./a.txt", os.O_RDONLY|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭文件
	defer file.Close()

	//写入数据
	bs := []byte{65, 66, 67, 68, 69}
	n, err := file.Write(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	writeString, err := file.WriteString("hahhahah,你好")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(writeString)

}
