package main

import (
	_ "github.com/QQorp/QQorpBackend/docs"
	_ "github.com/QQorp/QQorpBackend/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
