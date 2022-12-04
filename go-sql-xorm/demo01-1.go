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
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("结构体同步成功")
}
