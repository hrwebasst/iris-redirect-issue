package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/kataras/iris"
)

var w sync.WaitGroup

func staticMedia(ctx *iris.Context) {

	filepath := fmt.Sprintf("assets/static%s", ctx.Param("path"))

	data, err := Asset(filepath)
	if err != nil {
		ctx.Render("text/markdown", "## No asset found")
	} else {
		if strings.HasSuffix(filepath, ".exe") {
			ctx.Render("application/javascript", string(data), iris.RenderOptions{"gzip": true})
		} else {
			ctx.Data(iris.StatusOK, data)
		}
	}
}

func index(ctx *iris.Context) {
	name := ctx.Session().GetString("name")
	login := ctx.Session().GetString("login")

	ctx.Render("index.html", map[string]interface{}{
		"username": name,
		"is_admin": login,
		"color":    "green",
	}, iris.RenderOptions{"gzip": true})
}

func kill(ctx *iris.Context) {
	ctx.Render("dude.html", map[string]interface{}{
		"username": "dude",
	}, iris.RenderOptions{"gzip": true})

}

func birth(ctx *iris.Context) {
	ctx.Render("start.html", map[string]interface{}{
		"username": "dude",
	}, iris.RenderOptions{"gzip": true})

}
