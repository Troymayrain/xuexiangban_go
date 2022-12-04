package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	//请求处理
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	//模板渲染
	http.HandleFunc("/temp", temp)
	http.HandleFunc("/temp2", temp2)
	//端口
	http.ListenAndServe("127.0.0.1:80", nil)
}

type User struct {
	Name string
	Age  int
}

func temp2(writer http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("./client4.html")

	userMap := make(map[int]User)
	userMap[1] = User{"张三", 12}
	userMap[2] = User{"李四", 13}
	userMap[3] = User{"王五", 14}
	userMap[4] = User{"赵六", 15}

	data := make(map[string](map[int]User))
	data["data"] = userMap

	temp.Execute(writer, data)
}

func temp(writer http.ResponseWriter, request *http.Request) {
	//temp, _ := template.ParseFiles("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson15\\client4.html")
	temp, _ := template.ParseFiles("./client4.html")

	data := make(map[string]string)
	data["info"] = "hello,mayu"

	temp.Execute(writer, data)
}

func register(writer http.ResponseWriter, request *http.Request) {
	//解析post参数
	request.ParseForm()
	postForm := request.PostForm
	fmt.Println(postForm)
	username := postForm.Get("username")
	password := postForm.Get("password")
	fmt.Println("username:", username, "password:", password)
	writer.Write([]byte(`{"注册状态":"ok"}`))
}

func login(writer http.ResponseWriter, request *http.Request) {
	//解析query参数
	query := request.URL.Query()
	username := query.Get("username")
	password := query.Get("password")
	fmt.Println("username:", username, "password:", password)

	//业务逻辑
	if username == "admin" {
		writer.Write([]byte(`{"登录状态":"登录成功"}`))
	} else {
		writer.Write([]byte(`{"登录状态":"登录失败"}`))
	}

}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)
	fmt.Println(request.URL)
	fmt.Println(request.RemoteAddr)
	fmt.Println(request.Header)
	fmt.Println(request.Body)
	fmt.Println("程序被访问到了")
	writer.Write([]byte("<h1>hello world</h1>"))
}
