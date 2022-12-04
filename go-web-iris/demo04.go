package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	//全局中间件处理
	app.Use(before)
	app.Done(after)

	//app.Get("/index", before, index, after)
	app.Get("/index", index)
	app.Get("/", index)
	app.Listen(":8080")
}

func before(c *context.Context) {
	//鉴权、用户权限登记
	c.Values().Set("msg", "hello")
	fmt.Println("before")
	//交给一下个处理器处理
	c.Next()
}

func index(context *context.Context) {
	msg := context.Values().Get("msg")
	fmt.Println("index===>", msg)
	//交给一下个处理器处理
	context.Next()
}

func after(c *context.Context) {
	fmt.Println("after")
	//交给一下个处理器处理
	c.Next()
}
