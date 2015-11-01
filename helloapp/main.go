package main

import (
	_ "blog/helloapp/routers"
	_ "blog/models"
	"blog/models/class"
	"github.com/astaxie/beego"
)

func main() {
	class.TestROM
	beego.Run()
}
