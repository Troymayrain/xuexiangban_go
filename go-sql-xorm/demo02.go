package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {
	// 准备数据库连接信息
	var (
		username  string = "root"
		password  string = "005071"
		ipAddress string = "127.0.0.1"
		port      string = "3306"
		schema    string = "go_test"
		charset   string = "utf8mb4"
	)
	//构建数据库连接
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, ipAddress, port, schema, charset)
	fmt.Println(dbStr)

	engine, err := xorm.NewEngine("mysql", dbStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")

	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	//插入1条数据
	user := User{Id: 1, Name: "mayu", Age: 12, Passwd: "123456"}
	n, _ := engine.Insert(&user)
	fmt.Println(n)
	//插入2条数据
	user1 := User{Id: 2, Name: "mayu2", Age: 12, Passwd: "123456"}
	user2 := User{Id: 3, Name: "mayu3", Age: 12, Passwd: "123456"}
	n, _ = engine.Insert(&user1, &user2)
	fmt.Println(n)
	//插入多条数据
	user3 := User{Id: 4, Name: "mayu4", Age: 12, Passwd: "123456"}
	user4 := User{Id: 5, Name: "mayu5", Age: 12, Passwd: "123456"}
	var users []User
	users = append(users, user3, user4)
	n, _ = engine.Insert(&users)
	fmt.Println(n)
}
