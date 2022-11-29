package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.OpenFile("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11\\a.txt", os.O_RDWR, os.ModePerm)
	defer file.Close()

	//bufio 读
	reader := bufio.NewReader(file)
	buf := make([]byte, 1024, 1024)
	n, _ := reader.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf[:n]))

	//键盘输入
	inputReader := bufio.NewReader(os.Stdin)
	readString, _ := inputReader.ReadString('\n')
	fmt.Println("键盘输入的内容：", readString)

	//bufio 写
	writer := bufio.NewWriter(file)
	writeNum, _ := writer.WriteString("hello")
	fmt.Println(writeNum)

	//当写入文件的内容长度小于缓冲区的长度时，写完后的内容还是停留在缓冲区，故需要调用flush()方法，刷到文件中
	writer.Flush()
}
