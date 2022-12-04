package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/index", index1)
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Get("/destory", func(c *context.Context) {
		session.Destroy(c)
	})

	app.Listen(":8080")
}

var (
	cookie  = "sessionid"
	session = sessions.New(sessions.Config{Cookie: cookie})
)

func login(c *context.Context) {
	//开启会话
	session := session.Start(c)

	//用户身份已验证
	session.Set("authenticated", true)
}

func logout(c *context.Context) {
	//开启会话
	session := session.Start(c)

	//用户身份已验证
	session.Set("authenticated", false)
}

func index1(context *context.Context) {
	//开启会话
	session := session.Start(context)
	auth, _ := session.GetBoolean("authenticated")
	if auth == false {
		context.StatusCode(iris.StatusForbidden)
		return
	}
	fmt.Println("用户已经登录了")
}
