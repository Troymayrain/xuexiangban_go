package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlx连接mysql
func main() {
	//获取mysql连接
	mysqlDb := connectMYSQL()
	//关闭连接
	defer mysqlDb.Close()

	//testExec(mysqlDb)

	//testQuery(mysqlDb)

	type user struct {
		Id         int    `db:"id"`
		Userid     int    `db:"userid"`
		Username   string `db:"username"`
		Password   string `db:"password"`
		Avatar     string `db:"avatar"`
		CreateTime string `db:"create_time"`
		UpdateTime string `db:"update_time"`
	}

	//Get查询
	userData := new(user)
	err := mysqlDb.Get(userData, "select * from user where id=9")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("userData:", userData)
	fmt.Println("userData.username:", userData.Username)

	//Select查询
	var userSlice []user
	err = mysqlDb.Select(&userSlice, "select * from user")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range userSlice {
		fmt.Println(userSlice[i])
	}
}

// 测试Query的查询数据
func testQuery(mysqlDb *sqlx.DB) {
	//查询数据
	selectSQL := "select * from user"
	rows, err := mysqlDb.Query(selectSQL)
	if err != nil {
		fmt.Println("查询数据失败：", err)
		return
	}
	fmt.Println("rows:", rows)
	for rows.Next() {
		var id, userid int
		var username, password, avatar, create_time, update_time string
		rows.Scan(&id, &userid, &username, &password, &avatar, &create_time, &update_time)
		fmt.Println(id, userid, userid, username, password, avatar, create_time, update_time)
	}
	rows.Close()
}

// 测试Exec的增删改
func testExec(mysqlDb *sqlx.DB) {
	//插入数据
	insertSQL := "insert into user (userid, username, password, avatar, create_time, update_time) values (?,?,?,?,?,?)"
	result, err := mysqlDb.Exec(insertSQL, 10000, "mayu1", "123456", "ks.png", "2022-12-03 21:36:30", "2022-12-03 21:37:30")
	if err != nil {
		fmt.Println("数据插入失败：", err)
	}
	id, _ := result.LastInsertId()
	fmt.Println("数据插入成功，last id:", id)

	//修改
	updateSQL := "update user set username=? where id=?"
	result1, err := mysqlDb.Exec(updateSQL, "张三", 7)
	rowNum, _ := result1.RowsAffected()
	fmt.Println("更新成功，rows affected:", rowNum)

	//删除
	deleteSQL := "delete from user where id=?"
	result2, err := mysqlDb.Exec(deleteSQL, 7)
	rowsAffected, _ := result2.RowsAffected()
	fmt.Println("删除成功，rows affected:", rowsAffected)
}

// 准备数据库连接信息
var (
	username  string = "root"
	password  string = "005071"
	ipAddress string = "127.0.0.1"
	port      string = "3306"
	schema    string = "go_test"
	charset   string = "utf8mb4"
)

// 连接mysql数据库
func connectMYSQL() *sqlx.DB {
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", username, password, ipAddress, port, schema, charset)
	db, err := sqlx.Open("mysql", dbStr)
	fmt.Println(err)

	ping(db)

	return db
}

// 测试是否连接成功
func ping(db *sqlx.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
		return
	} else {
		fmt.Println("ping success")
	}
}
