package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	//iris.Compression是iris内部对io数据进行压缩的模块，可以提高数据的传输速度
	app.Use(iris.Compression)

	app.Get("/hello", hello)

	//http://localhost:8080/user?userid=1&username=mayu
	//http://localhost:8080/user/{userid:int64}/{username:string}
	app.Get("/user/{userid:int64}/{username:string}", func(context *context.Context) {
		//解析query参数
		//userid := context.URLParam("userid")
		//username := context.URLParam("username")
		//解析path参数
		userid := context.Params().Get("userid")
		username := context.Params().Get("username")
		app.Logger().Info(userid + "--" + username)
		context.WriteString(userid + "--" + username)
	})

	//处理表单
	app.Post("/user/login", func(context *context.Context) {
		password := context.PostValue("password")
		username := context.PostValue("username")
		app.Logger().Info(password + "--" + username)
		context.WriteString(password + "--" + username)
	})

	//处理json
	app.Post("/user/json", func(context *context.Context) {
		//获取json参数
		var user = struct {
			Username string `json:"username`
			Password string `json:"password`
		}{}
		err := context.ReadJSON(&user)
		if err != nil {
			panic(err)
		}
		app.Logger().Info(user)
		context.JSON(iris.Map{"msg": user.Username, "code": 200})
	})

	app.Listen(":8080")
}

func hello(ctx iris.Context) {
	ctx.HTML("Hello <strong>%s</strong>!", "World")
}
