package main

import (
	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {
	iris.Set(iris.OptionDisableBanner(true))

	iris.UseTemplate(django.New()).Directory("./assets/templates", ".html").Binary(Asset, AssetNames)

	iris.Get("/static/*path", staticMedia)

	iris.Get("/login/:user", func(ctx *iris.Context) {

		ctx.Session().Set("name", ctx.Param("user"))
		ctx.Session().Set("login", "yup")

		ctx.Write("All ok session set to: %s", ctx.Session().GetString("name"))
	})

	iris.Get("/get", func(ctx *iris.Context) {
		name := ctx.Session().GetString("name")
		login := ctx.Session().GetString("login")

		ctx.Write("The name on the /set was: %s, logged in: %s", name, login)
	})

	iris.Get("/yousuck", func(ctx *iris.Context) {
		ctx.Write("Dun do it")
	})

	iris.Get("/test/test2.css", func(ctx *iris.Context) {
		ctx.Render("text/css", "body {color:blue;}", iris.RenderOptions{"charset": "UTF-8"})
	})

	iris.Get("/test/test.css", func(ctx *iris.Context) {
		ctx.Render("text/css", "body {color:blue;}", iris.RenderOptions{"charset": "UTF-8"})
	})

	iris.Get("/", index)

	iris.Get("/kill", kill)

	iris.Get("/birth", birth)

	iris.Listen(":8080")

}
