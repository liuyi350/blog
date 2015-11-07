package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.krsen.cc"
	c.Data["Email"] = "ping.zeng@krsen.cc"
	c.TplNames = "index.tpl"
}
