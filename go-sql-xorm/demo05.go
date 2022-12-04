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

	session := engine.NewSession()
	defer session.Close()

	//开启事务
	session.Begin()

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	user1 := User{Id: 8, Name: "张三", Age: 23, Passwd: "123456"}
	if _, err := session.Insert(user1); err != nil {
		panic(err)
	}

	if _, err := session.Exec("delete from user where id1='2'"); err != nil {
		panic(err)
	}

}
