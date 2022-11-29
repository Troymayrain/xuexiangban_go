package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	//源文件
	srcFile := "E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\bg5.png"
	//目标文件
	destFile := "E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\seek\\11.png"
	//临时文件
	tempFile := "E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\temp.txt"

	//建立连接
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	//关闭连接
	defer file1.Close()
	defer file2.Close()

	//1、读取临时文件记录的数据
	file3.Seek(0, io.SeekStart)
	buf := make([]byte, 1024, 1024)
	n, _ := file3.Read(buf)
	countStr := string(buf[:n])
	fmt.Println("countStr:", countStr)

	//字符串转int64
	count, _ := strconv.ParseInt(countStr, 10, 64)

	//设置源文件、目标文件的读写位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)

	//int64转int
	total := int(count)

	readData := make([]byte, 1024, 1024)
	for {

		//读源文件
		readNUm, err := file1.Read(readData)
		if err == io.EOF {
			fmt.Println("文件读取完了")
			file3.Close()
			os.Remove(tempFile)
			break
		}

		//写目标文件
		file2.Write(readData[:readNUm])
		total += readNUm

		//将传输进度写入临时文件
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		//模拟断电
		//if total > 4000 {
		//	panic("断电了")
		//}

	}

}
