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

	//Query
	resultsSlice, _ := engine.Query("select * from user")
	fmt.Println(resultsSlice)
	//QueryString
	resultsSlice2, _ := engine.QueryString("select * from user")
	fmt.Println(resultsSlice2)
	//QueryInterface
	resultsSlice3, _ := engine.QueryInterface("select * from user")
	fmt.Println(resultsSlice3)

	//Get
	user := User{}
	engine.Get(&user)
	fmt.Println(user)

	//Get 根据条件
	user1 := User{}
	engine.Where("name =?", "mayu3").Get(&user1)
	fmt.Println(user1)
	//Get 指定字段
	name := ""
	engine.Table(&user).Where("id = 3").Cols("name").Get(&name)
	fmt.Println(name)

	//Find
	var users1 []User
	engine.Where("age=12").And("passwd=123456").Limit(10, 0).Find(&users1)
	fmt.Println(users1)

	//Iterate
	engine.Iterate(&User{Age: 12}, func(idx int, bean interface{}) error {
		u := bean.(*User)
		fmt.Println(u)
		return nil
	})

	//Rows
	rows, _ := engine.Rows(&User{Age: 12})
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user)
		fmt.Println(user)
	}

}
