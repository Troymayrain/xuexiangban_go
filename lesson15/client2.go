package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	urlStr := "http://127.0.0.1:80/login"

	//拼接参数
	data := url.Values{}
	data.Set("username", "mayu")
	data.Set("password", "123456")

	rurl, _ := url.ParseRequestURI(urlStr)
	rurl.RawQuery = data.Encode()
	fmt.Println("rurl:", rurl.String())

	//发送GET请求
	response, _ := http.Get(rurl.String())
	readAll, _ := io.ReadAll(response.Body)
	fmt.Println(string(readAll))

}
