package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//发送GET请求
	response, _ := http.Get("http://127.0.0.1/hello")
	//关闭流
	defer response.Body.Close()

	buf := make([]byte, 1024)
	str := ""
	for {
		//循环读取
		n, err := response.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			str += string(buf[:n])
			if err == io.EOF {
				fmt.Println("数据读取完毕")
				fmt.Println("data:", str)
				break
			}
		}
	}
}
