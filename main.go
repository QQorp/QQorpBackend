package main

import (
	_ "github.com/QQorp/QQorpBackend/app/docs"
	_ "github.com/QQorp/QQorpBackend/app/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
