package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/http"
)

func main() {

	app := iris.New()

	app.Use(iris.Compression)

	//页面渲染配置
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))
	//静态资源配置
	app.HandleDir("static", http.Dir("./static"))

	//404
	app.OnErrorCode(iris.StatusNotFound, not404)
	//500
	app.OnErrorCode(iris.StatusInternalServerError, func(c *context.Context) {

	})

	//302重定向
	app.Get("/index", func(context *context.Context) {
		login := true
		if login {
			context.ViewData("msg", "hello,world")
			context.View("index.html")
		} else {
			context.StatusCode(302)
			context.View("login.html")
		}
	})
	app.Get("/login", func(context *context.Context) {
		context.View("login.html")
	})

	app.Run(iris.Addr(":8080"))
}

func not404(c *context.Context) {
	c.View("404.html")
}
