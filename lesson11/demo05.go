package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	source := "E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\img\\bg5.png"
	destination := "E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\bg5-1.png"
	//copy(source, destination, 1024)\
	//copy2(source, destination)
	copy3(source, destination)
}

// 利用os包实现
// 不建议使用此方式，文件过大，容易内存溢出
func copy3(source, destinaton string) {
	file, _ := os.ReadFile(source)
	os.WriteFile(destinaton, file, os.ModePerm)
}

// 利用io.copy()实现文件复制
func copy2(source, destination string) {
	//输入文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceFile.Close()

	//输出文件
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destinationFile.Close()

	written, _ := io.Copy(destinationFile, sourceFile)
	fmt.Println("复制的文件大小为：", written)
}

// 手动实现文件复制
func copy(source, destination string, buffSize int) {

	//输入文件
	sourceFile, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sourceFile.Close()

	//输出文件
	destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destinationFile.Close()

	//读
	buff := make([]byte, buffSize)
	for {
		read, err := sourceFile.Read(buff)
		if read == 0 || err == io.EOF {
			fmt.Println("文件读完了。。。")
			break
		}

		//写
		_, err = destinationFile.Write(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("文件复制成功！")

}
