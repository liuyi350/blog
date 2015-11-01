package main

import (
	_ "blog/helloapp/models"
	"blog/helloapp/models/class"
	_ "blog/helloapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	class.TestORM()
	beego.Run()
}
