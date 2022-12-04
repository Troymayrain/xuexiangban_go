package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {

	app := iris.New()

	app.Use(iris.Compression)

	//路由组
	userParty := app.Party("/user", func(c *context.Context) {
		app.Logger().Warn("hahahahahha")
		c.Next()
	})
	userParty.Get("/login", func(context *context.Context) {})
	userParty.Get("/logout", func(context *context.Context) {})
	userParty.Get("/register", func(context *context.Context) {})
	userParty.Get("/get", func(context *context.Context) {})

	orderParty := app.Party("/order", func(c *context.Context) {
		c.Next()
	})
	orderParty.Get("/...", func(context *context.Context) {})
	orderParty.Get("/...", func(context *context.Context) {})
	orderParty.Get("/...", func(context *context.Context) {})
	orderParty.Get("/...", func(context *context.Context) {})

	app.Listen(":8081")

}
