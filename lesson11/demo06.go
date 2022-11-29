package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//seek用法
	file, _ := os.OpenFile("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\a.txt", os.O_RDWR, os.ModePerm)
	defer file.Close()

	file.Seek(2, io.SeekStart)
	buff := []byte{0}
	n, _ := file.Read(buff)
	fmt.Println(len(buff))
	fmt.Println(n)
	fmt.Println(string(buff))

	file.Seek(2, io.SeekCurrent)
	_, _ = file.Read(buff)
	fmt.Println(string(buff))

	file.Seek(0, io.SeekEnd)
	file.WriteString("hhhhsfdhdsfhsdf")

}
