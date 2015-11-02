package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "115.28.202.156"
	c.Data["Email"] = "350117156@qq.com"
	c.TplNames = "index.tpl"
}
